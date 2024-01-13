package llm

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type testsStorage interface {
	GetTestById(ctx context.Context, testId tests.TestId) (*tests.Test, error)
}

type llmStorage interface {
	GetLLMChecksByTestId(
		ctx context.Context,
		testId tests.TestId,
	) ([]ModelCheck, error)

	InsertNotStartedLLMCheck(
		ctx context.Context,
		modelSlug ModelSlug,
		testId tests.TestId,
		authorId users.UserId,
	) (*ModelCheck, error)
}
