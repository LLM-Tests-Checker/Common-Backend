package tests

import (
	"context"
	"errors"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	logger   *logrus.Logger
	database *mongo.Database
}

func NewTestsStorage(
	logger *logrus.Logger,
	database *mongo.Database,
) *Storage {
	return &Storage{
		logger:   logger,
		database: database,
	}
}

func (storage *Storage) GetTestById(
	ctx context.Context,
	testId tests.TestId,
) (*tests.Test, error) {
	return nil, errors.New("not implemented yet")
}

func (storage *Storage) GetTestsByAuthorId(
	ctx context.Context,
	authorId users.UserId,
	offset int32,
	size int32,
) ([]tests.Test, error) {
	return nil, errors.New("not implemented yet")
}

func (storage *Storage) CreateTest(
	ctx context.Context,
	authorId users.UserId,
	createData tests.CreateTestData,
) (*tests.Test, error) {
	return nil, errors.New("not implemented yet")
}

func (storage *Storage) DeleteTestById(
	ctx context.Context,
	testId tests.TestId,
) error {
	return errors.New("not implemented yet")
}
