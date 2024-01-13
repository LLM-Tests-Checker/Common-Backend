package tests

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type testsStorage interface {
	GetTestById(
		ctx context.Context,
		testId TestId,
	) (*Test, error)

	GetTestsByAuthorId(
		ctx context.Context,
		authorId users.UserId,
		offset int32,
		size int32,
	) ([]Test, error)

	CreateTest(
		ctx context.Context,
		authorId users.UserId,
		createData CreateTestData,
	) (*Test, error)

	DeleteTestById(
		ctx context.Context,
		testId TestId,
	) error
}
