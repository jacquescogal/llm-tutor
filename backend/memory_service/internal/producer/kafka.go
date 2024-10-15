package producer

import (
	"context"
	"log"
	"memory_core/internal/proto/document_job"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"google.golang.org/protobuf/proto"
)

// KafkaProducer wraps the Kafka producer instance
type KafkaProducer struct {
	Producer *kafka.Producer
	topic   string
}

// NewKafkaProducer creates a new Kafka producer instance
func NewKafkaProducer() *KafkaProducer {
	bootstrapServers := os.Getenv("KAFKA_SERVER")
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		log.Fatalf("Failed to create producer: %v\n", err)
	}
	topic := os.Getenv("KAFKA_TOPIC")
	return &KafkaProducer{Producer: p, topic: topic}
}

// ProduceMessage produces a message to a Kafka topic by sending a serialized protobuf message
func (kp *KafkaProducer) ProduceMessage(ctx context.Context, job *document_job.DocumentProcessingJob) error {
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	// Serialize the protobuf message to binary format
	jobBytes, err := proto.Marshal(job)
	if err != nil {
		log.Printf("Failed to marshal protobuf message: %v\n", err)
		return err
	}

	// Produce the serialized message to the Kafka topic
	err = kp.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kp.topic, Partition: kafka.PartitionAny},
		Value:         jobBytes, // Sending the serialized protobuf message
	}, deliveryChan)
	if err != nil {
		log.Printf("Failed to produce message: %v\n", err)
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		return m.TopicPartition.Error
	}

	log.Printf("Delivered message to topic %s [%d] at offset %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	return nil
}

