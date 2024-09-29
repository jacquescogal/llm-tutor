import boto3
from src.utils.configuration import Configuration
import tempfile
import pymupdf4llm

class ObjectStore:
    def __init__(self, bucket_name):
        self.config = Configuration.get_instance()
        self.bucket_name = bucket_name
        self.s3 = boto3.client('s3')

    def get_text_object(self, object_key):
        try:
            # Create a temporary file
            with tempfile.NamedTemporaryFile(delete=True) as temp_file:
                self.s3.download_fileobj(self.bucket_name, object_key, temp_file)
                temp_file_path = temp_file.name
                md_text = pymupdf4llm.to_markdown(temp_file_path)
                print(f"PDF processed successfully.")
                return md_text
        except Exception as e:
            print(f"Error processing the PDF: {e}")


    def delete_object(self, object_key):
        self.s3.delete_object(Bucket=self.bucket_name, Key=object_key)
