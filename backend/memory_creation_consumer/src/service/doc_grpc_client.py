import grpc
from grpc import aio
import src.protos.document_pb2 as document_pb2
import src.protos.document_pb2_grpc as document_pb2_grpc
from src.utils.configuration import Configuration

# in charge of making gRPC calls to the doc service
class DocGRPCClientService:
    def __init__(self):
        self.memory_service_address = Configuration.get_instance().memory_service_address

    async def update_summary(self, message: document_pb2.UpdateSummaryRequest) -> document_pb2.UpdateSummaryResponse:
        try:
            async with aio.insecure_channel(self.memory_service_address) as channel:
                stub = document_pb2_grpc.DocServiceStub(channel)
                response: document_pb2.UpdateSummaryResponse = await stub.UpdateSummary(message)
                return response
        except grpc.RpcError as e:
            print(f"gRPC Error: {e.code()}, {e.details()}")
            return document_pb2.UpdateSummaryResponse() 

    async def update_upload_status(self, message: document_pb2.UpdateUploadStatusRequest) -> document_pb2.UpdateUploadStatusResponse:
        try:
            async with aio.insecure_channel(self.memory_service_address) as channel:
                stub = document_pb2_grpc.DocServiceStub(channel)
                response: document_pb2.UpdateUploadStatusResponse = await stub.UpdateUploadStatus(message)
                return response
        except grpc.RpcError as e:
            print(f"gRPC Error: {e.code()}, {e.details()}")
            return document_pb2.UpdateUploadStatusResponse()
        
    async def get_doc_by_id(self, message: document_pb2.GetDocByIdRequest) -> document_pb2.GetDocByIdResponse:
        try:
            async with aio.insecure_channel(self.memory_service_address) as channel:
                stub = document_pb2_grpc.DocServiceStub(channel)
                response: document_pb2.GetDocByIdResponse = await stub.GetDocById(message)
                return response
        except grpc.RpcError as e:
            print(f"gRPC Error: {e.code()}, {e.details()}")
            return document_pb2.GetDocByIdResponse()