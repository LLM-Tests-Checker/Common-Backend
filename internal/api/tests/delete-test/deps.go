package delete_test

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type testDeleter interface {
	DeleteTest(
		ctx context.Context,
		authorId users.UserId,
		testId tests.TestId,
	) error
}

type tokenParser interface {
	ParseUserId(ctx context.Context, accessToken string) (users.UserId, error)
}
