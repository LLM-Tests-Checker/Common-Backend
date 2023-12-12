package services

import (
	"errors"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/components/llm"
)

type LLMCheckerService struct {
	launcher       llm.Launcher
	statusSelector llm.StatusSelector
	resultSelector llm.ResultSelector
}

func NewLLMCheckerService(
	launcherService llm.Launcher,
	statusSelector llm.StatusSelector,
	resultSelector llm.ResultSelector,
) *LLMCheckerService {
	return &LLMCheckerService{
		launcher:       launcherService,
		statusSelector: statusSelector,
		resultSelector: resultSelector,
	}
}

func (service *LLMCheckerService) LaunchLLMCheck(currentUserId int32, testId, llmSlug string) (error, *llm.LaunchResult) {
	return errors.New("not implemented yet"), nil
}

func (service *LLMCheckerService) GetLLMCHeckStatus(currentUserId int32, testId string) (error, *llm.CheckStatus) {
	return errors.New("not implemented yet"), nil
}

func (service *LLMCheckerService) GetLLMCheckResult(currentUserId int32, testId string) (error, *llm.CheckResult) {
	return errors.New("not implemented yet"), nil
}
