import grpc
from grpc import aio
import src.protos.memory_pb2 as memory_pb2
import src.protos.memory_pb2_grpc as memory_pb2_grpc
from src.utils.configuration import Configuration

# in charge of making gRPC calls to the memory service
class MemoryGRPCClientService:
    def __init__(self):
        self.memory_service_address = Configuration.get_instance().memory_service_address

    async def create_memory(self, message: memory_pb2.CreateMemoryRequest) -> memory_pb2.CreateMemoryResponse:
        try:
            async with aio.insecure_channel(self.memory_service_address) as channel:
                stub = memory_pb2_grpc.MemoryServiceStub(channel)
                response: memory_pb2.CreateMemoryResponse = await stub.CreateMemory(message)
                return response
        except grpc.RpcError as e:
            print(f"gRPC Error: {e.code()}, {e.details()}")
            return memory_pb2.CreateMemoryResponse() 