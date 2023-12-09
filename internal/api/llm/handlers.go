package llm

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/components/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services"
	"net/http"
)

var (
	llmService *services.LLMCheckerService
)

func init() {
	launcher := llm.NewLauncher()
	statusSelector := llm.NewStatusSelector()
	resultSelector := llm.NewResultSelector()

	llmService = services.NewLLMCheckerService(launcher, statusSelector, resultSelector)
}

func LaunchLLMCheckHandler(responseWriter http.ResponseWriter, request *http.Request) {
	llmService.LaunchLLMCheck()

}
func GetLLMCheckStatusHandler(responseWriter http.ResponseWriter, request *http.Request) {
	llmService.GetLLMCHeckStatus()
}

func GetLLMCheckResultHandler(responseWriter http.ResponseWriter, request *http.Request) {
	llmService.GetLLMCheckResult()
}
