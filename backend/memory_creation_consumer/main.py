from src.consumer.kafka_consumer import KafkaListener
import asyncio
import sys
import os

# sys.path.append(os.path.join(os.path.dirname(__file__), 'src'))
sys.path.append(os.path.abspath(os.path.join(os.path.dirname(__file__), '..', 'src')))

async def main():
    listener = KafkaListener(max_workers=5)
    await listener.listen()

if __name__ == "__main__":
    asyncio.run(main())