import grpc
from concurrent import futures
import src.protos.generation_pb2 as generation_pb2
import src.protos.generation_pb2_grpc as generation_pb2_grpc
from src.llm.query_bot import  QueryBot
from src.protos.generation_pb2 import CreateGenerationRequest, CreateGenerationResponse

# Implementing the gRPC generationService
class GenerationServiceServicer(generation_pb2_grpc.generationServiceServicer):
    def CreateGeneration(self, request:CreateGenerationRequest, context):
        qb = QueryBot.get_instance()
        response = qb.query(request.chat_history, request.id, request.id_type)  
        print(response)
        return CreateGenerationResponse(response=response)
# gRPC server setup
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    generation_pb2_grpc.add_generationServiceServicer_to_server(GenerationServiceServicer(), server)
    server.add_insecure_port('[::]:50060')
    server.start()
    print("Server started on port 50060")
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
