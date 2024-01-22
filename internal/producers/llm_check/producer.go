package llm_check

import (
	"context"
	"github.com/sirupsen/logrus"
)

type Producer struct {
	logger *logrus.Logger
}

func NewProducer(
	logger *logrus.Logger,
) *Producer {
	return &Producer{
		logger: logger,
	}
}

func (producer *Producer) ProduceEvents(ctx context.Context, events []CheckEvent) error {
	return nil
}
