package get_statuses

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type statusSelector interface {
	GetModelStatusesForTest(userId users.UserId, testId tests.TestId) ([]llm.ModelCheckStatus, error)
}

type tokenParser interface {
	ParseUserId(accessToken string) (users.UserId, error)
}
