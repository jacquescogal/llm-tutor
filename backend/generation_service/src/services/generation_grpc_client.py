import grpc
from grpc import aio
import src.protos.vector_pb2 as vector_pb2
import src.protos.vector_pb2_grpc as vector_pb2_grpc
from src.utils.configuration import Configuration

# in charge of making gRPC calls to the memory service for the vector documents
class GenerationGRPCClientService:
    def __init__(self):
        self.memory_service_address = Configuration.get_instance().memory_service_address

    async def asearch_memory_vector_request(self, message: vector_pb2.SearchMemoryVectorRequest) -> vector_pb2.SearchMemoryVectorResponse:
        try:
            async with aio.insecure_channel(self.memory_service_address) as channel:
                stub = vector_pb2_grpc.VectorServiceStub(channel)
                response: vector_pb2.SearchMemoryVectorResponse = await stub.SearchMemoryVector(message)
                return response
        except grpc.RpcError as e:
            print(f"gRPC Error: {e.code()}, {e.details()}")
            return vector_pb2.SearchMemoryVectorResponse() 
        
        
    def search_memory_vector_request(self, message: vector_pb2.SearchMemoryVectorRequest) -> vector_pb2.SearchMemoryVectorResponse:
        try:
            with grpc.insecure_channel(self.memory_service_address) as channel:
                stub = vector_pb2_grpc.VectorServiceStub(channel)
                response: vector_pb2.SearchMemoryVectorResponse = stub.SearchMemoryVector(message)
                return response
        except grpc.RpcError as e:
            print(f"gRPC Error: {e.code()}, {e.details()}")
            return vector_pb2.SearchMemoryVectorResponse()