from sqlalchemy import Column, BigInteger, Enum, String, Text
from src.models import Base

class Job(Base):
    __tablename__ = 'job_tab'
    
    job_id = Column(BigInteger, primary_key=True, autoincrement=True)
    job_status = Column(Enum('queueing', 'processing', 'to approve', 'inserting', 'inserted', 'failed'), nullable=False)
    object_key = Column(String(255), nullable=False)
    created_time = Column(BigInteger, nullable=False)
    last_updated_time = Column(BigInteger, nullable=False)
    job_summary = Column(Text, nullable=False)