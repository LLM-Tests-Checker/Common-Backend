package get_test

import (
	"context"
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type testSelector interface {
	GetTestById(
		ctx context.Context,
		userId users.UserId,
		testId tests.TestId,
	) (*tests.Test, error)
}

type testMapper interface {
	MapModelToDto(model *tests.Test) dto.Test
}

type tokenParser interface {
	ParseUserId(ctx context.Context, accessToken string) (users.UserId, error)
}
