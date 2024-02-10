package llm_check

import (
	"context"
	"encoding/json"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/platform/logger"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	logger logger.Logger
	writer *kafka.Writer
}

func NewProducer(
	logger logger.Logger,
	writer *kafka.Writer,
) *Producer {
	return &Producer{
		logger: logger,
		writer: writer,
	}
}

func (producer *Producer) ProduceEvents(ctx context.Context, events []CheckEvent) error {
	kafkaMessages := make([]kafka.Message, len(events))
	for i := range events {
		kafkaMessages[i] = convertEventModelToMessage(events[i])
	}

	err := producer.writer.WriteMessages(ctx, kafkaMessages...)
	if err != nil {
		return error2.Wrap(err, "producer.writer.WriteMessages")
	}

	return nil
}

func convertEventModelToMessage(event CheckEvent) kafka.Message {
	messageKey := []byte(event.ModelCheckId)
	messageValue, err := json.Marshal(&event)
	if err != nil {
		panic(err)
	}

	return kafka.Message{
		Key:   messageKey,
		Value: messageValue,
		Headers: []kafka.Header{
			{
				Key:   CheckEventMessageVersionHeader,
				Value: []byte(CheckEventVersion),
			},
		},
	}
}
