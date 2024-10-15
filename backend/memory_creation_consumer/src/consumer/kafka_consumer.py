from confluent_kafka import Consumer, KafkaException
from src.protos.document_job_pb2 import DocumentProcessingJob
from src.controller.doc_controller import DocumentController
import asyncio
from concurrent.futures import ThreadPoolExecutor
from src.utils.configuration import Configuration
from src.llm.summary_llm import SummaryLLM
from src.service.doc_grpc_client import DocGRPCClientService
from src.service.memory_grpc_client import MemoryGRPCClientService
from src.service.question_grpc_client import QuestionGRPCClientService
from src.db.object_store import ObjectStore

class KafkaListener:
    def __init__(self, max_workers=5):
        # Kafka consumer configuration
        kafka_server = Configuration.get_instance().kafka_address
        kafka_topic = Configuration.get_instance().kafka_topic
        conf = {
            'bootstrap.servers': kafka_server,
            'group.id': 'create_group',
            'auto.offset.reset': 'earliest',
        }
        self.consumer = Consumer(conf)
        self.topic = kafka_topic
        summaryLLM = SummaryLLM()
        docClient = DocGRPCClientService()
        memoryClient = MemoryGRPCClientService()
        questionClient = QuestionGRPCClientService()
        object_store = ObjectStore()
        self.document_controller = DocumentController(
            summaryLLM, docClient, memoryClient, questionClient, object_store
        )

        # Subscribe to the Kafka topic
        self.consumer.subscribe([self.topic])

        # Thread pool executor to handle Kafka polling
        self.executor = ThreadPoolExecutor(max_workers=max_workers)

    async def process_message(self, msg):
        try:
            # Unmarshal the RPC-encoded Kafka message (deserialize protobuf)
            received_message = DocumentProcessingJob()
            received_message.ParseFromString(msg.value())

            # Extract the content from the deserialized message
            print(f"Received RPC-encoded message from Kafka: {received_message}")

            # Forward the content to the async message controller method
            response = await self.document_controller.process_doc(received_message)
            print(f"Response from gRPC server: {response}")

        except Exception as e:
            # Handle exceptions during deserialization or processing
            print(f"Error processing message: {e}")

    async def listen(self):
        # Get the event loop
        loop = asyncio.get_event_loop()

        try:
            while True:
                # Poll messages from Kafka in the thread pool (non-blocking)
                msg = await loop.run_in_executor(self.executor, self.consumer.poll, 1.0)

                if msg is None:
                    continue
                if msg.error():
                    raise KafkaException(msg.error())

                # Schedule the message processing task asynchronously
                asyncio.create_task(self.process_message(msg))

        except KafkaException as e:
            # Log Kafka errors
            print(f"Kafka error: {e}")

        finally:
            # Close the consumer on exit
            self.consumer.close()

    async def shutdown(self):
        print("Shutting down KafkaListener...")
        self.consumer.close()


# from confluent_kafka import Consumer, KafkaException
# from src.protos.document_job_pb2 import DocumentProcessingJob
# from src.controller.doc_controller import DocumentController
# import asyncio
# from concurrent.futures import ThreadPoolExecutor
# from confluent_kafka import Consumer, KafkaException
# from src.utils.configuration import Configuration

# class KafkaListener:
#     def __init__(self, topic, max_workers=5):
#         # Kafka consumer configuration
#         kafka_server = Configuration.get_instance().kafka_address
#         conf = {
#             'bootstrap.servers': kafka_server,
#             'group.id': 'create_group',
#             'auto.offset.reset': 'earliest',
#         }
#         self.consumer = Consumer(conf)
#         self.topic = topic
#         self.document_controller = DocumentController()

#         # Subscribe to the Kafka topic
#         self.consumer.subscribe([self.topic])

#         # Semaphore to limit the number of concurrent tasks (workers)
#         self.semaphore = asyncio.Semaphore(max_workers)

#     async def process_message(self, msg):
#         async with self.semaphore:
#             try:
#                 # Unmarshal the RPC-encoded Kafka message (deserialize protobuf)
#                 received_message = DocumentProcessingJob()
#                 received_message.ParseFromString(msg.value())

#                 # Extract the content from the deserialized message
#                 received_message
#                 print(f"Received RPC-encoded message from Kafka: {received_message}")

#                 # Forward the content to the async message controller method
#                 response = await self.document_controller.process_doc(received_message)
#                 print(f"Response from gRPC server: {response}")
#             except Exception as e:
#                 print(f"Error processing message: {e}")

#     async def listen(self):
#         # Create a thread pool executor to handle Kafka polling
#         loop = asyncio.get_event_loop()
#         with ThreadPoolExecutor() as executor:
#             try:
#                 while True:
#                     # Poll messages from Kafka in a thread pool (non-blocking)
#                     msg = await loop.run_in_executor(executor, self.consumer.poll, 1.0)

#                     if msg is None:
#                         continue
#                     if msg.error():
#                         raise KafkaException(msg.error())

#                     # Schedule the message processing task
#                     asyncio.create_task(self.process_message(msg))
            
#             except KafkaException as e:
#                 print(f"Kafka error: {e}")
                
#             finally:
#                 # Close the consumer on exit
#                 self.consumer.close()
    
#     async def shutdown(self):
#         # Close the Kafka consumer
#         self.consumer.close()
