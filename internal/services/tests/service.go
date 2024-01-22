package tests

import (
	"context"
	"fmt"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"net/http"
)

type Service struct {
	testsStorage testsStorage
}

func NewTestsService(
	testsStorage testsStorage,
) *Service {
	return &Service{
		testsStorage: testsStorage,
	}
}

func (service *Service) CreateTest(
	ctx context.Context,
	authorID users.UserId,
	data CreateTestData,
) (*Test, error) {
	createdTest, err := service.testsStorage.CreateTest(ctx, authorID, data)
	if err != nil {
		return nil, err
	}

	return createdTest, nil
}

func (service *Service) DeleteTest(
	ctx context.Context,
	authorId users.UserId,
	testId TestId,
) error {
	targetTest, err := service.testsStorage.GetTestById(ctx, testId)
	if err != nil {
		return err
	}

	if targetTest.AuthorId != authorId {
		err := error2.NewBackendError(
			error2.NotOwnerError,
			"Not your test",
			http.StatusForbidden,
		)
		return err
	}

	err = service.testsStorage.DeleteTestById(ctx, testId)
	if err != nil {
		return err
	}

	return nil
}

func (service *Service) GetTestsByAuthorId(
	ctx context.Context,
	authorId users.UserId,
	pageNumber, pageSize int32,
) ([]Test, error) {
	offset := pageNumber * pageSize
	tests, err := service.testsStorage.GetTestsByAuthorId(ctx, authorId, offset, pageSize)
	if err != nil {
		return nil, err
	}

	return tests, nil
}

func (service *Service) GetTestById(
	ctx context.Context,
	userId users.UserId,
	testId TestId,
) (*Test, error) {
	test, err := service.testsStorage.GetTestById(ctx, testId)
	if err != nil {
		return nil, err
	}
	if test == nil {
		return nil, error2.NewBackendError(
			error2.TestNotFound,
			fmt.Sprintf("Test with id %s not found", testId.String()),
			http.StatusBadRequest,
		)
	}

	if test.AuthorId != userId {
		err := error2.NewBackendError(
			error2.NotOwnerError,
			"Not your test",
			http.StatusForbidden,
		)
		return nil, err
	}

	return test, nil
}
