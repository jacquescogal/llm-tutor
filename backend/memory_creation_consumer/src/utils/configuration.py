import os
from dotenv import load_dotenv
from datetime import datetime

class Configuration:
    _instance = None

    def __init__(self):
        if Configuration._instance is not None:
            raise Exception("This class is a singleton!")
        else:
            load_dotenv()
            # Define environment variables here
            self.s3_bucket_name = os.getenv("S3_BUCKET_NAME")
            epoch_date = os.getenv("EPOCH_DATE_YYYYMMDD")
            self.epoch_timestamp = int(datetime.strptime(epoch_date, "%Y%m%d").timestamp()) 
            self.openai_api_key = os.getenv("OPENAI_API_KEY")
            self.openai_api_model = os.getenv("OPENAI_API_MODEL")
            self.memory_service_address = os.getenv("MEMORY_SERVICE_ADDRESS")
            self.kafka_address = os.getenv("KAFKA_ADDRESS")
            self.kafka_topic = os.getenv("KAFKA_TOPIC")
            self.s3_bucket_name = os.getenv("S3_BUCKET_NAME")
            Configuration._instance = self

    @classmethod
    def get_instance(cls):
        if cls._instance is None:
            cls._instance = Configuration()
        return cls._instance