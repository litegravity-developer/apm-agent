package dtos

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/litegravity-developer/apm-agent/custom_error"
	kafka_producers "github.com/litegravity-developer/apm-agent/kafka/producer"
)

func GetSpan(spanName string) *Record {
	record := Record{
		Name:     spanName,
		IsParent: true,
	}

	return &record
}

func (parentSpan *Record) GetChildSpan(childSpanName string) (*Record, error) {
	if parentSpan == nil {
		err := custom_error.ValidationError("got nil parent span")
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
		err := custom_error.ValidationError("got nil span")
		return err
	}
	currentTime := time.Now()
	record.StartTime = &currentTime
	return nil
}

func (record *Record) EndSpan() error {
	if record == nil {
		err := custom_error.ValidationError("got nil span")
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
		err := kafka_producers.ProduceToKafka("apmAgentLocal", transaction)
		if err != nil {
			println("err in producing data %v", err.Error())
		} else {
			fmt.Printf("pushed to kafka %+v", transaction.Span.Parent.Child)
		}
	}
	return nil
}
