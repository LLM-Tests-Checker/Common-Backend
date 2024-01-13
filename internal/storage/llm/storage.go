package llm

import (
	"context"
	"errors"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	logger   *logrus.Logger
	database *mongo.Database
}

func NewLLMStorage(
	logger *logrus.Logger,
	database *mongo.Database,
) *Storage {
	return &Storage{
		logger:   logger,
		database: database,
	}
}

func (service *Storage) GetLLMChecksByTestId(
	ctx context.Context,
	testId tests.TestId,
) ([]llm.ModelCheck, error) {
	return nil, errors.New("not implemented yet")
}

func (service *Storage) InsertNotStartedLLMCheck(
	ctx context.Context,
	modelSlug llm.ModelSlug,
	testId tests.TestId,
	authorId users.UserId,
) (*llm.ModelCheck, error) {
	return nil, errors.New("not implemented yet")
}
