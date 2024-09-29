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
            self.database_url = os.getenv("DATABASE_URL")
            self.s3_bucket_name = os.getenv("S3_BUCKET_NAME")
            epoch_date = os.getenv("EPOCH_DATE_YYYYMMDD")
            self.epoch_timestamp = int(datetime.strptime(epoch_date, "%Y%m%d").timestamp()) 
            self.openai_api_key = os.getenv("OPENAI_API_KEY")
            self.openai_api_model = os.getenv("OPENAI_API_MODEL")
            Configuration._instance = self

    @classmethod
    def get_instance(cls):
        if cls._instance is None:
            cls._instance = Configuration()
        return cls._instance