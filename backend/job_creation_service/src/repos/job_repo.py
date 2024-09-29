from sqlalchemy import update
from sqlalchemy.ext.asyncio import AsyncSession
from src.models.job_model import Job

class JobRepository:
    def __init__(self, db: AsyncSession):
        self.db = db

    async def update_job(self, job_id: int, status: str, last_updated_time: int, job_summary: str):
        stmt = (
            update(Job)
            .where(Job.job_id == job_id)
            .values(job_status=status, last_updated_time=last_updated_time, job_summary=job_summary)
        )
        await self.db.execute(stmt)
        await self.db.commit()
