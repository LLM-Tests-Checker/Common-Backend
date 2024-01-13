package delete_test

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type testDeleter interface {
	DeleteTest(authorId users.UserId, testId tests.TestId) error
}
