package get_statuses

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type statusSelector interface {
	GetModelStatusesForTest(
		ctx context.Context,
		userId users.UserId,
		testId tests.TestId,
	) ([]llm.ModelCheckStatus, error)
}

type tokenParser interface {
	ParseUserId(ctx context.Context, accessToken string) (users.UserId, error)
}
