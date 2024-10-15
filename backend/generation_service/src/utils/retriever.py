from typing import List

from langchain_core.callbacks import CallbackManagerForRetrieverRun
from langchain_core.documents import Document
from langchain_core.retrievers import BaseRetriever
from src.services.generation_grpc_client import GenerationGRPCClientService
from src.protos.vector_pb2 import SearchMemoryVectorResponse, SearchMemoryVectorRequest
from src.protos.common_pb2 import IDType

class QuestionRetriever(BaseRetriever):
    k: int
    vector_id: int
    id_type: IDType
    def _get_relevant_documents(
        self, query: str, *, run_manager: CallbackManagerForRetrieverRun
    ) -> List[Document]:
        """Sync implementations for retriever."""
        client = GenerationGRPCClientService()
        search_vector_request:SearchMemoryVectorRequest = SearchMemoryVectorRequest(search_query=query, limit=self.k, id=self.vector_id, id_type=self.id_type)
        search_memory_vector_response:SearchMemoryVectorResponse = client.search_memory_vector_request(search_vector_request)
        # json_response is []bytes
        # unmarshall it now
        print("working", search_memory_vector_response)
        json_string = search_memory_vector_response.json_response.decode('utf-8')
        
        aDocument = Document(page_content=json_string)  
        return [aDocument]

    # Optional: Provide a more efficient native implementation by overriding
    # _aget_relevant_documents
    # async def _aget_relevant_documents(
    #     self, query: str, *, run_manager: AsyncCallbackManagerForRetrieverRun
    # ) -> List[Document]:
    #     """Asynchronously get documents relevant to a query.
# 
    #     Args:
    #         query: String to find relevant documents for
    #         run_manager: The callbacks handler to use

    #     Returns:
    #         List of relevant documents
    #     """

# from typing import List

# from langchain_core.callbacks import CallbackManagerForRetrieverRun
# from langchain_core.documents import Document
# from langchain_core.retrievers import BaseRetriever


# class ToyRetriever(BaseRetriever):
#     """A toy retriever that contains the top k documents that contain the user query.

#     This retriever only implements the sync method _get_relevant_documents.

#     If the retriever were to involve file access or network access, it could benefit
#     from a native async implementation of `_aget_relevant_documents`.

#     As usual, with Runnables, there's a default async implementation that's provided
#     that delegates to the sync implementation running on another thread.
#     """

#     documents: List[Document]
#     """List of documents to retrieve from."""
#     k: int
#     """Number of top results to return"""

#     def _get_relevant_documents(
#         self, query: str, *, run_manager: CallbackManagerForRetrieverRun
#     ) -> List[Document]:
#         """Sync implementations for retriever."""
#         matching_documents = []
#         for document in documents:
#             if len(matching_documents) > self.k:
#                 return matching_documents

#             if query.lower() in document.page_content.lower():
#                 matching_documents.append(document)
#         return matching_documents

#     # Optional: Provide a more efficient native implementation by overriding
#     # _aget_relevant_documents
#     # async def _aget_relevant_documents(
#     #     self, query: str, *, run_manager: AsyncCallbackManagerForRetrieverRun
#     # ) -> List[Document]:
#     #     """Asynchronously get documents relevant to a query.

#     #     Args:
#     #         query: String to find relevant documents for
#     #         run_manager: The callbacks handler to use

#     #     Returns:
#     #         List of relevant documents
#     #     """