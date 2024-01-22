package llm

import (
	"context"
	"fmt"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"net/http"
)

type Service struct {
	testsStorage testsStorage
	llmStorage   llmStorage
}

func NewLLMService(
	testsStorage testsStorage,
	llmStorage llmStorage,
) *Service {
	return &Service{
		testsStorage: testsStorage,
		llmStorage:   llmStorage,
	}
}

func (service *Service) GetAllResultsForTest(
	ctx context.Context,
	userId users.UserId,
	testId tests.TestId,
) ([]ModelTestResult, error) {
	test, err := service.testsStorage.GetTestById(ctx, testId)
	if err != nil {
		return nil, err
	}
	if test == nil {
		err := error2.NewBackendError(
			error2.TestNotFound,
			fmt.Sprintf("Test with id %s not found", test.Identifier),
			http.StatusBadRequest,
		)
		return nil, err
	}
	if test.AuthorId != userId {
		err := error2.NewBackendError(
			error2.NotOwnerError,
			"Not your test",
			http.StatusForbidden,
		)
		return nil, err
	}

	modelChecks, err := service.llmStorage.GetLLMChecksByTestId(ctx, testId)
	if err != nil {
		return nil, err
	}

	results := make([]ModelTestResult, len(modelChecks))
	for i := range modelChecks {
		results[i] = ModelTestResult{
			LLMSlug:    modelChecks[i].ModelSlug,
			TestId:     modelChecks[i].TargetTestId,
			LLMAnswers: modelChecks[i].Answers,
		}
	}

	return results, nil
}

func (service *Service) GetModelStatusesForTest(
	ctx context.Context,
	userId users.UserId,
	testId tests.TestId,
) ([]ModelCheckStatus, error) {
	test, err := service.testsStorage.GetTestById(ctx, testId)
	if err != nil {
		return nil, err
	}
	if test == nil {
		err := error2.NewBackendError(
			error2.TestNotFound,
			fmt.Sprintf("Test with id %s not found", test.Identifier),
			http.StatusBadRequest,
		)
		return nil, err
	}
	if test.AuthorId != userId {
		err := error2.NewBackendError(
			error2.NotOwnerError,
			"Not your test",
			http.StatusForbidden,
		)
		return nil, err
	}

	modelChecks, err := service.llmStorage.GetLLMChecksByTestId(ctx, testId)
	if err != nil {
		return nil, err
	}

	results := make([]ModelCheckStatus, len(modelChecks))
	for i := range modelChecks {
		results[i] = ModelCheckStatus{
			LLMSlug: modelChecks[i].ModelSlug,
			TestId:  modelChecks[i].TargetTestId,
			Status:  modelChecks[i].Status,
		}
	}

	return results, nil
}

func (service *Service) LaunchModelCheck(
	ctx context.Context,
	userId users.UserId,
	testId tests.TestId,
	modelSlug ModelSlug,
) (*ModelCheck, error) {
	test, err := service.testsStorage.GetTestById(ctx, testId)
	if err != nil {
		return nil, err
	}
	if test == nil {
		err := error2.NewBackendError(
			error2.TestNotFound,
			fmt.Sprintf("Test with id %s not found", test.Identifier),
			http.StatusBadRequest,
		)
		return nil, err
	}
	if test.AuthorId != userId {
		err := error2.NewBackendError(
			error2.NotOwnerError,
			"Not your test",
			http.StatusForbidden,
		)
		return nil, err
	}

	modelCheck, err := service.llmStorage.InsertNotStartedLLMCheck(ctx, modelSlug, testId, userId)
	if err != nil {
		return nil, err
	}

	return modelCheck, nil
}
