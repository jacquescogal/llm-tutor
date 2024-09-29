from src.repos.job_repo import JobRepository
from src.repos.doc_repo import DocumentRepository
from src.repos.object_repo import ObjectRepository
from src.utils.time_util import TimeUtil
from src.utils.chunk_util import ChunkUtil
from src.utils.splitter_util import SplitterUtil
from src.llm.summary_llm import SummaryLLM
import json

class DocumentController:
    def __init__(self, job_repo: JobRepository, doc_repo: DocumentRepository, object_repo: ObjectRepository, summary_llm: SummaryLLM):
        self.job_repo = job_repo
        self.doc_repo = doc_repo
        self.object_repo = object_repo
        self.summary_llm = summary_llm

    async def create_docs(self, job_id: int, topic_id: int, object_key: str):
        # Update job status to 'processing'
        await self.job_repo.update_job(job_id, 'processing', last_updated_time=TimeUtil.get_current_unix_time_from_epoch())

        # Get the temp file from S3
        # TODO: add conditional between text only and text+image
        md_text = self.object_repo.get_object(object_key)

        # count chunks in the whole document
        total_token_count = ChunkUtil.ChunkTokenCount(md_text)

        summary_docs = ""
        # summary chunk count should be 100_000 max so document + instructions fit into llm context window
        if total_token_count > 100_000:
            chunks = SplitterUtil.split_by_markdown(md_text)
            summary_docs = await self._recursive_create_memories(chunks, total_token_count)
        else:
            summary_docs = await self.summary_llm.create_memory_documents(md_text)
        
        # Create the documents
        # TODO: ensure that this is json and update the job status to failed if not
        try:
            docs = json.loads(summary_docs)
        except Exception as e:
            print(f"Error parsing the summary: {e}")
            await self.job_repo.update_job(job_id, 'failed', last_updated_time=TimeUtil.get_current_unix_time_from_epoch())
            docs = []
            return

        for doc in docs:
            unix_time_now = TimeUtil.get_current_unix_time_from_epoch()
            await self.doc_repo.create_document(topic_id, job_id, doc['memory_card_header'], doc['memory_card_body'], created_time=unix_time_now, last_updated_time=unix_time_now)
            # TODO: put into vector store here (vector)

        # Update job status to 'to approve'
        await self.job_repo.update_job(job_id, 'to approve', last_updated_time=TimeUtil.get_current_unix_time_from_epoch())

        # Delete the object from the S3 bucket
        # self.object_repo.delete_object(object_key)

    async def _recursive_create_memories(self, chunks, chunks_token_count)->str:
        if len(chunks) == 0:
            return ""
        elif chunks_token_count < 100_000:
            # summarise down to 16_000 tokens max (max for llm)
            return await self.summary_llm.summarise_to_structured_document('\n'.join(chunks))
        elif len(chunks) == 1:
            # is more than 100_000 tokens and only one chunk
            # split into chunks by character and recursively summarise
            chunks = SplitterUtil.split_by_characters(chunks[0], token_size=50_000, chunk_overlap_ratio=0)
        mid = len(chunks) // 2
        left_chunks = chunks[:mid]
        right_chunks = chunks[mid:]
        left_chunk_token_count = ChunkUtil.ChunkTokenCount('\n'.join(left_chunks))
        right_chunk_token_count = ChunkUtil.ChunkTokenCount('\n'.join(right_chunks))
        result = (await self._recursive_create_memories(left_chunks, left_chunk_token_count) + '\n' + await self._recursive_create_memories(right_chunks, right_chunk_token_count)).strip()
        result_chunk_token_count = ChunkUtil.ChunkTokenCount(result)
        if result_chunk_token_count> 100_000:
            # split into chunks by character and recursively summarise
            # this will rarely happen unless the document is very very large
            result = await self._recursive_create_memories([result], result_chunk_token_count)
        return result