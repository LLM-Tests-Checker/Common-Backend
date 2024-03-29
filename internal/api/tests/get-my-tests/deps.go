package get_my_tests

import (
	"context"
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type testsSelector interface {
	GetTestsByAuthorId(
		ctx context.Context,
		authorId users.UserId,
		pageNumber, pageSize int32,
	) ([]tests.Test, error)
}

type testMapper interface {
	MapModelToDto(model *tests.Test) dto.Test
}

type tokenParser interface {
	ParseUserId(ctx context.Context, accessToken string) (users.UserId, error)
}
