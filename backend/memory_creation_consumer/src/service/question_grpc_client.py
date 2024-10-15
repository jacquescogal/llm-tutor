import grpc
from grpc import aio
import src.protos.question_pb2 as question_pb2
import src.protos.question_pb2_grpc as question_pb2_grpc
from src.utils.configuration import Configuration

# in charge of making gRPC calls to the question service
class QuestionGRPCClientService:
    def __init__(self):
        self.memory_service_address = Configuration.get_instance().memory_service_address

    async def create_question(self, message: question_pb2.CreateQuestionRequest) -> question_pb2.CreateQuestionResponse:
        try:
            async with aio.insecure_channel(self.memory_service_address) as channel:
                stub = question_pb2_grpc.QuestionServiceStub(channel)
                response: question_pb2.CreateQuestionResponse = await stub.CreateQuestion(message)
                return response
        except grpc.RpcError as e:
            print(f"gRPC Error: {e.code()}, {e.details()}")
            return question_pb2.CreateQuestionResponse()