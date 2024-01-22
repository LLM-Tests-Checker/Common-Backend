package llm_result

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
	logger *logrus.Logger
	reader *kafka.Reader
}

func NewConsumer(
	logger *logrus.Logger,
	reader *kafka.Reader,
) *Consumer {
	return &Consumer{
		logger: logger,
		reader: reader,
	}
}

func (consumer *Consumer) Start(ctx context.Context) error {
	return nil
}
