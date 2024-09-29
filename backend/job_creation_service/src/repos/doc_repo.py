from sqlalchemy import insert
from sqlalchemy.ext.asyncio import AsyncSession
from src.models.doc_model import Document

class DocumentRepository:
    def __init__(self, db: AsyncSession):
        self.db = db

    async def create_document(self, topic_id: int, job_id: int, doc_title: str, markdown_text: str, created_time: int, last_updated_time: int):
        stmt = (
            insert(Document)
            .values(
                topic_id=topic_id,
                job_id=job_id,
                doc_title=doc_title,
                markdown_text=markdown_text,
                created_time=created_time,
                last_updated_time=last_updated_time
            )
        )
        await self.db.execute(stmt)
        await self.db.commit()
