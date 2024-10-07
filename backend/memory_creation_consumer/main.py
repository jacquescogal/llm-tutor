from concurrent import futures
import time
import sys
import os
# import relative to cwd/src for generated proto files
cwd = os.getcwd()
src_path = os.path.join(cwd, 'src')
sys.path.append(src_path)

from src.protos import job_pb2, job_pb2_grpc
import grpc


class GreeterServicer(job_pb2_grpc.JobServiceServicer):
    def ProcessJob(self, request, context):
        print("Received job request")
        return job_pb2.ProcessJobResponse(message="hello")
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    job_pb2_grpc.add_JobServiceServicer_to_server(GreeterServicer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Server started at port 50051")
    try:
        while True:
            time.sleep(86400)  # Sleep for a day
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()
