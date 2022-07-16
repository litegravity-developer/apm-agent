package apm

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func createError(message string) CustomError {
	error_ := CustomError{Message: message}
	return error_
}

type CustomError struct {
	Message string `json:"message"`
}

func (c CustomError) Error() string {
	return c.Message
}

func ValidationError(message string) CustomError {

	error_ := createError(message)
	return error_
}

//dtos

type Transaction struct {
	Id   string `json:"id"`
	Span Span   `json:"span"`
}

type Span struct {
	Parent Record `json:"parent"`
}

type Record struct {
	Name      string     `json:"name"`
	StartTime *time.Time `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	Child     []*Record  `json:"child"`
	IsParent  bool       `json:"-"`
}

//dtos methods

func GetSpan(spanName string) *Record {
	record := Record{
		Name:     spanName,
		IsParent: true,
	}

	return &record
}

func (parentSpan *Record) GetChildSpan(childSpanName string) (*Record, error) {
	if parentSpan == nil {
		err := ValidationError("got nil parent span")
		return nil, err
	}

	subRecord := Record{
		Name:     childSpanName,
		IsParent: false,
	}

	parentSpan.Child = append(parentSpan.Child, &subRecord)
	return &subRecord, nil
}

func (record *Record) StartSpan() error {
	if record == nil {
		err := ValidationError("got nil span")
		return err
	}
	currentTime := time.Now()
	record.StartTime = &currentTime
	return nil
}

func (record *Record) EndSpan() error {
	if record == nil {
		err := ValidationError("got nil span")
		return err
	}

	currentTime := time.Now()
	record.EndTime = &currentTime

	if record.IsParent == true {
		id := uuid.New()
		transaction := Transaction{
			Id: id.String(),
			Span: Span{
				Parent: *record,
			},
		}
		err := ProduceToKafka("apmAgentLocal", transaction)
		if err != nil {
			println("err in producing data %v", err.Error())
		} else {
			fmt.Printf("pushed to kafka %+v", transaction.Span.Parent.Child)
		}
	}
	return nil
}

//kafka

func GetKafkaWriterInstance() *kafka.Writer {

	writerConfig := kafka.WriterConfig{
		Brokers:      []string{"localhost:9092"},
		BatchSize:    1,
		BatchTimeout: 100 * time.Millisecond,
		RequiredAcks: 1,
	}
	Writer := kafka.NewWriter(writerConfig)
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
