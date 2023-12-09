package services

import "github.com/LLM-Tests-Checker/Common-Backend/internal/components/llm"

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

func (service *LLMCheckerService) LaunchLLMCheck() {

}

func (service *LLMCheckerService) GetLLMCHeckStatus() {

}

func (service *LLMCheckerService) GetLLMCheckResult() {

}
