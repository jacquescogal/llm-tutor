from sqlalchemy import Column, BigInteger, String, Text
from src.models import Base

class Memory(Base):
    __tablename__ = 'memory_tab'
    
    memory_id = Column(BigInteger, primary_key=True, autoincrement=True)
    topic_id = Column(BigInteger, nullable=False)
    job_id = Column(BigInteger, nullable=False)
    memory_title = Column(String(255), nullable=False)
    memory_body = Column(Text, nullable=False)
    created_time = Column(BigInteger, nullable=False)
    last_updated_time = Column(BigInteger, nullable=False)
