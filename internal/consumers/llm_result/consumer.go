package llm_result

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/platform/logger"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	logger     logger.Logger
	reader     *kafka.Reader
	llmStorage llmStorage
}

func NewConsumer(
	logger logger.Logger,
	reader *kafka.Reader,
	llmStorage llmStorage,
) *Consumer {
	return &Consumer{
		logger:     logger,
		reader:     reader,
		llmStorage: llmStorage,
	}
}

func (consumer *Consumer) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			err := consumer.consume(ctx)
			if err != nil {
				consumer.logger.Errorf("Consumer fetch finished with error: %s", err)
			}
		}
	}
}

// Method is not transactional. Updating database and consuming event execute not consistently.
func (consumer *Consumer) consume(
	ctx context.Context,
) error {
	message, err := consumer.reader.FetchMessage(ctx)
	if err != nil {
		return error2.Wrap(err, "consumer.reader.FetchMessage")
	}

	resultEvent, err := parseLLMResultMessage(message)
	if err != nil {
		commitErr := consumer.reader.CommitMessages(ctx, message)
		if commitErr != nil {
			consumer.logger.Errorf("consumer.reader.CommitMessages: %s", commitErr)
		}
		return error2.Wrap(err, "parseLLMResultMessage")
	}

	modelCheckId := uuid.MustParse(resultEvent.ModelCheckId)
	modelAnswers := make([]llm.ModelTestAnswer, len(resultEvent.ModelAnswers))
	for i := range resultEvent.ModelAnswers {
		modelAnswers[i] = llm.ModelTestAnswer{
			QuestionNumber:       resultEvent.ModelAnswers[i].QuestionNumber,
			SelectedAnswerNumber: resultEvent.ModelAnswers[i].AnswerNumber,
		}
	}

	err = consumer.llmStorage.SetLLMCheckCompleted(ctx, modelCheckId, modelAnswers)
	if err != nil {
		return error2.Wrap(err, "consumer.llmStorage.SetLLMCheckCompleted")
	}

	err = consumer.reader.CommitMessages(ctx, message)
	if err != nil {
		return error2.Wrap(err, "consumer.reader.CommitMessages")
	}

	return nil
}

func parseLLMResultMessage(message kafka.Message) (ResultEvent, error) {
	messageVersion := ""
	for _, header := range message.Headers {
		if header.Key == ResultEventMessageVersionHeader {
			messageVersion = string(header.Value)
			break
		}
	}
	if messageVersion != ResultEventVersion {
		return ResultEvent{}, errors.New(
			fmt.Sprintf(
				"Unsupported result event version: %s, supported only %s",
				messageVersion,
				ResultEventVersion,
			),
		)
	}

	resultEvent := ResultEvent{}
	err := json.Unmarshal(message.Value, &resultEvent)
	if err != nil {
		return ResultEvent{}, error2.Wrap(err, "json.Unmarshal")
	}

	return resultEvent, nil
}
