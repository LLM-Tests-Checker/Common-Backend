package get_results

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type resultsSelector interface {
	GetAllResultsForTest(
		ctx context.Context,
		userId users.UserId,
		testId tests.TestId,
	) ([]llm.ModelTestResult, error)
}

type tokenParser interface {
	ParseUserId(ctx context.Context, accessToken string) (users.UserId, error)
}
