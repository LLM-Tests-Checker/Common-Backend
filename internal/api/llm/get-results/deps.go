package get_results

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type resultsSelector interface {
	GetAllResultsForTest(userId users.UserId, testId tests.TestId) ([]llm.ModelTestResult, error)
}

type tokenParser interface {
	ParseUserId(accessToken string) (users.UserId, error)
}
