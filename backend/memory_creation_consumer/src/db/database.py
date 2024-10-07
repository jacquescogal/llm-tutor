from sqlalchemy import create_engine
from sqlalchemy.ext.asyncio import AsyncSession, create_async_engine
from sqlalchemy.orm import sessionmaker
from src.utils.configuration import Configuration


class Database:
    def __init__(self):
        self.config = Configuration.get_instance()
        self.engine = create_async_engine(self.config.database_url, echo=True)
        self.SessionLocal = sessionmaker(
            bind=self.engine,
            class_=AsyncSession,
            expire_on_commit=False
        )

    async def get_db(self):
        async with self.SessionLocal() as session:
            yield session

