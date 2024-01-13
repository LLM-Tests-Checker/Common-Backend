package launch_check

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type checkLauncher interface {
	LaunchModelCheck(
		ctx context.Context,
		userId users.UserId,
		testId tests.TestId,
		modelSlug llm.ModelSlug,
	) (*llm.ModelCheck, error)
}

type tokenParser interface {
	ParseUserId(ctx context.Context, accessToken string) (users.UserId, error)
}
