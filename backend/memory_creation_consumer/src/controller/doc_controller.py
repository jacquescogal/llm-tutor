from src.utils.time_util import TimeUtil
from src.utils.chunk_util import ChunkUtil
from src.utils.splitter_util import SplitterUtil
from src.llm.summary_llm import SummaryLLM
from src.service.doc_grpc_client import DocGRPCClientService
from src.service.memory_grpc_client import MemoryGRPCClientService
from src.service.question_grpc_client import QuestionGRPCClientService
from src.protos.document_job_pb2 import DocumentProcessingJob
from src.protos.document_pb2 import UpdateSummaryRequest, UpdateUploadStatusRequest, GetDocByIdRequest
from src.protos.memory_pb2 import CreateMemoryRequest
from src.protos.question_pb2 import CreateQuestionRequest, MCQChoice, MCQQuestion, TextInputQuestion
from src.protos.common_pb2 import QuestionType
from src.protos.common_pb2 import UploadStatus
from src.db.object_store import ObjectStore
import json
import asyncio

class DocumentController:
    def __init__(self, summary_llm: SummaryLLM, docClient: DocGRPCClientService, memoryClient: MemoryGRPCClientService, questionClient: QuestionGRPCClientService, object_store: ObjectStore):
        self.docClient = docClient
        self.memoryClient = memoryClient
        self.questionClient = questionClient
        self.object_store = object_store
        self.summary_llm = summary_llm

    async def process_doc(self, document_processing_job: DocumentProcessingJob) -> None:
        # Update job status to 'processing'
        updateUploadStatusRequest = UpdateUploadStatusRequest(
            user_id=document_processing_job.user_id,
            module_id=document_processing_job.module_id,
            doc_id=document_processing_job.doc_id,
            upload_status=UploadStatus.UPLOAD_STATUS_PROCESSING
        )
        await self.docClient.update_upload_status(updateUploadStatusRequest)
        print(f"Processing document: {document_processing_job.doc_id}")


        # get from the doc_tab
        db_document = await self.docClient.get_doc_by_id(GetDocByIdRequest(
            user_id=document_processing_job.user_id,
            module_id=document_processing_job.module_id,
            doc_id=document_processing_job.doc_id
        ))
        print(f"Document: {db_document}")


        # Get the temp file from S3
        # TODO: add conditional between text only and text+image
        md_text = self.object_store.get_text_object(db_document.doc.s3_object_key)

        # count chunks in the whole document
        total_token_count = ChunkUtil.ChunkTokenCount(md_text)

        summary_docs = ""
        # summary chunk count should be 100_000 max so document + instructions fit into llm context window
        if total_token_count > 100_000:
            chunks = SplitterUtil.split_by_markdown(md_text)
            summary_docs = await self._recursive_summarise_memories(db_document.doc.doc_name, chunks, total_token_count)
        else:
            summary_docs = await self.summary_llm.summarise_document(db_document.doc.doc_name, md_text)
        print(f"Summary docs: {summary_docs}")
        # Create the memories
        memory_task = self.summary_llm.create_memory_documents(db_document.doc.doc_name, summary_docs)
        mcq_task = self.summary_llm.create_mcq_question(db_document.doc.doc_name, summary_docs, 2)
        multi_answer_task = self.summary_llm.create_multi_answer_question(db_document.doc.doc_name, summary_docs, 2)
        open_ended_task = self.summary_llm.create_open_answer_question(db_document.doc.doc_name, summary_docs, 2)

        # Gather the tasks with individual error handling
        results = await asyncio.gather(
            memory_task,
            mcq_task,
            multi_answer_task,
            open_ended_task,
            return_exceptions=True  # Allows other tasks to complete even if one fails
        )

        # Process the results
        memories, mcq_questions, multi_answer_questions, open_ended_questions = results

        if isinstance(memories, Exception):
            print(f"Error creating memories: {memories}")
        else:
            print(f"Memories created: {memories}")
        if isinstance(mcq_questions, Exception):
            print(f"Error creating MCQ questions: {mcq_questions}")
        else:
            print(f"MCQ questions created: {mcq_questions}")

        if isinstance(multi_answer_questions, Exception):
            print(f"Error creating multi-answer questions: {multi_answer_questions}")
        else:
            print(f"Multi-answer questions created: {multi_answer_questions}")

        if isinstance(open_ended_questions, Exception):
            print(f"Error creating open-ended questions: {open_ended_questions}")
        else:
            print(f"Open-ended questions created: {open_ended_questions}")
        
        # Update the job status to 'completed'
        await self.docClient.update_summary(UpdateSummaryRequest(
            user_id=document_processing_job.user_id,
            doc_id=db_document.doc.doc_id,
            module_id=db_document.doc.module_id,
            doc_summary=summary_docs,
            upload_status=UploadStatus.UPLOAD_STATUS_SUCCESS
        ))

        await self.upload_data(memories, mcq_questions, multi_answer_questions, open_ended_questions, document_processing_job)

        # # Upload the memories
        # for each_memory in memories:
        #     await self.memoryClient.create_memory(
        #         CreateMemoryRequest(
        #             user_id=document_processing_job.user_id,
        #             doc_id=document_processing_job.doc_id,
        #             module_id=document_processing_job.module_id,
        #             memory_title=each_memory['memory_title'],
        #             memory_content=each_memory['memory_content']
        #         )
        #     )
        
        # # Upload the questions
        # for each_mcq_question in mcq_questions:
        #     mcqQuestionList = []
        #     for choice in each_mcq_question["choices"]:
        #         mcqQuestionList.append(MCQChoice(
        #             choice=choice["choice"],
        #             is_correct=choice["is_correct"]
        #         ))
        #     questionBlob = MCQQuestion(
        #         choices=mcqQuestionList,
        #     )
        #     serialized_question_blob = questionBlob.SerializeToString()  # Capture serialized bytes

        #     await self.questionClient.create_question(
        #         CreateQuestionRequest(
        #             user_id=document_processing_job.user_id,
        #             doc_id=document_processing_job.doc_id,
        #             module_id=document_processing_job.module_id,
        #             question_title=each_mcq_question["question"],
        #             question_blob=serialized_question_blob,
        #             question_type=QuestionType.QUESTION_TYPE_MCQ
        #         )
        #     )
        # for each_multi_answer_question in multi_answer_questions:
        #     mcqQuestionList = []
        #     for choice in each_multi_answer_question["choices"]:
        #         mcqQuestionList.append(MCQChoice(
        #             choice=choice["choice"],
        #             is_correct=choice["is_correct"]
        #         ))
        #     questionBlob = MCQQuestion(
        #         choices=mcqQuestionList,
        #     )
        #     serialized_question_blob = questionBlob.SerializeToString()  # Capture serialized bytes

        #     await self.questionClient.create_question(
        #         CreateQuestionRequest(
        #             user_id=document_processing_job.user_id,
        #             doc_id=document_processing_job.doc_id,
        #             module_id=document_processing_job.module_id,
        #             question_title=each_multi_answer_question["question"],
        #             question_blob=serialized_question_blob,
        #             question_type=QuestionType.QUESTION_TYPE_MULTI_ANSWER_MCQ
        #         )
        #     )

        # for open_question in open_ended_questions:
        #     mcqQuestionList = []
        #     questionBlob = TextInputQuestion(
        #         answer=open_question["answer"]
        #     )
        #     serialized_question_blob = questionBlob.SerializeToString()  # Capture serialized bytes

        #     await self.questionClient.create_question(
        #         CreateQuestionRequest(
        #             user_id=document_processing_job.user_id,
        #             doc_id=document_processing_job.doc_id,
        #             module_id=document_processing_job.module_id,
        #             question_title=open_question["question"],
        #             question_blob=serialized_question_blob,
        #             question_type=QuestionType.QUESTION_TYPE_OPEN_ANSWER
        #         )
        #     )
        
        # print(f"Document processed: {document_processing_job.doc_id}")

        # turn the questions into question
        
        

    async def _recursive_summarise_memories(self, doc_name, chunks, chunks_token_count)->str:
        if len(chunks) == 0:
            return ""
        elif chunks_token_count < 100_000:
            # summarise down to 16_000 tokens max (max for llm)
            return await self.summary_llm.summarise_document(doc_name, '\n'.join(chunks))
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
    
    async def upload_data(self, memories, mcq_questions, multi_answer_questions, open_ended_questions, document_processing_job: DocumentProcessingJob):
        # Create the tasks list for memories, MCQs, multi-answer MCQs, and open-ended questions
        tasks = []

        # Upload the memories
        if memories is not None:
            for each_memory in memories:
                tasks.append(self.memoryClient.create_memory(
                    CreateMemoryRequest(
                        user_id=document_processing_job.user_id,
                        doc_id=document_processing_job.doc_id,
                        module_id=document_processing_job.module_id,
                        memory_title=each_memory['memory_title'],
                        memory_content=each_memory['memory_content']
                    )
                ))

        # Upload the MCQ questions
        if mcq_questions is not None:
            for each_mcq_question in mcq_questions:
                mcqQuestionList = [
                    MCQChoice(
                        choice=choice["choice"],
                        is_correct=choice["is_correct"]
                    ) for choice in each_mcq_question["choices"]
                ]

                questionBlob = MCQQuestion(
                    choices=mcqQuestionList,
                )
                serialized_question_blob = questionBlob.SerializeToString()

                tasks.append(self.questionClient.create_question(
                    CreateQuestionRequest(
                        user_id=document_processing_job.user_id,
                        doc_id=document_processing_job.doc_id,
                        module_id=document_processing_job.module_id,
                        question_title=each_mcq_question["question"],
                        question_blob=serialized_question_blob,
                        question_type=QuestionType.QUESTION_TYPE_MCQ
                    )
                ))

        # Upload the multi-answer MCQ questions
        if multi_answer_questions is not None:
            for each_multi_answer_question in multi_answer_questions:
                mcqQuestionList = [
                    MCQChoice(
                        choice=choice["choice"],
                        is_correct=choice["is_correct"]
                    ) for choice in each_multi_answer_question["choices"]
                ]

                questionBlob = MCQQuestion(
                    choices=mcqQuestionList,
                )
                serialized_question_blob = questionBlob.SerializeToString()

                tasks.append(self.questionClient.create_question(
                    CreateQuestionRequest(
                        user_id=document_processing_job.user_id,
                        doc_id=document_processing_job.doc_id,
                        module_id=document_processing_job.module_id,
                        question_title=each_multi_answer_question["question"],
                        question_blob=serialized_question_blob,
                        question_type=QuestionType.QUESTION_TYPE_MULTI_ANSWER_MCQ
                    )
                ))

        # Upload the open-ended questions
        if open_ended_questions is not None:
            for open_question in open_ended_questions:
                questionBlob = TextInputQuestion(
                    answer=open_question["answer"]
                )
                serialized_question_blob = questionBlob.SerializeToString()

                tasks.append(self.questionClient.create_question(
                    CreateQuestionRequest(
                        user_id=document_processing_job.user_id,
                        doc_id=document_processing_job.doc_id,
                        module_id=document_processing_job.module_id,
                        question_title=open_question["question"],
                        question_blob=serialized_question_blob,
                        question_type=QuestionType.QUESTION_TYPE_OPEN_ANSWER
                    )
                ))
        print("tasks length:", len(tasks))
        # Run all tasks concurrently
        results = await asyncio.gather(*tasks, return_exceptions=True)

        # Process the results to check for exceptions
        for result in results:
            if isinstance(result, Exception):
                print(f"Task failed: {result}")
            else:
                print("Task completed successfully")