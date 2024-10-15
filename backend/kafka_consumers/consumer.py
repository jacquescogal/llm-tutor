from confluent_kafka import Consumer, KafkaError

# Kafka Consumer Configuration
conf = {
    'bootstrap.servers': 'localhost:9092',
    'group.id': 'test-group',
    'auto.offset.reset': 'earliest' 
}

# Create a Consumer instance
consumer = Consumer(conf)

# Subscribe to the topic
consumer.subscribe(['test-topic'])

try:
    print("Consumer is listening for messages...")

    while True:
        msg = consumer.poll(1.0)  # Poll for messages (timeout of 1 second)

        if msg is None:
            continue

        if msg.error():
            if msg.error().code() == KafkaError._PARTITION_EOF:
                print(f"End of partition reached {msg.topic()} [{msg.partition()}] at offset {msg.offset()}")
            elif msg.error():
                print(f"Error: {msg.error()}")
        else:
            # Print the message
            print(f"Received message: {msg.value().decode('utf-8')} from {msg.topic()} [{msg.partition()}]")

except KeyboardInterrupt:
    print("Consumer interrupted")

finally:
    # Close down the consumer when done
    consumer.close()
