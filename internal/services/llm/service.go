package llm

import (
	"errors"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type Service struct {
}

func NewLLMService() *Service {
	return &Service{}
}

func (service *Service) GetAllResultsForTest(
	userId users.UserId,
	testId tests.TestId,
) ([]ModelTestResult, error) {
	return nil, errors.New("not implemented yet")
}

func (service *Service) GetModelStatusesForTest(
	userId users.UserId,
	testId tests.TestId,
) ([]ModelCheckStatus, error) {
	return nil, errors.New("not implemented yet")
}

func (service *Service) LaunchModelCheck(
	userId users.UserId,
	testId tests.TestId,
	modelSlug ModelSlug,
) (*ModelCheck, error) {
	return nil, errors.New("not implemented yet")
}
