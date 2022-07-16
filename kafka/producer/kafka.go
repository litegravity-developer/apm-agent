package kafka_producers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

var Writer *kafka.Writer

func GetKafkaWriterInstance() *kafka.Writer {
	if Writer != nil {
		return Writer
	}
	writerConfig := kafka.WriterConfig{
		Brokers:      []string{"localhost:9092"},
		BatchSize:    1,
		BatchTimeout: 100 * time.Millisecond,
		RequiredAcks: 1,
	}
	Writer = kafka.NewWriter(writerConfig)
	return Writer
}

func ProduceToKafka(topicName string, event interface{}) error {
	writer := GetKafkaWriterInstance()

	byteEvent, err := json.Marshal(&event)
	if err != nil {
		return err
	}

	if err := writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Topic: topicName,
			Value: byteEvent,
		}); err != nil {
		return err
	}

	return nil
}
