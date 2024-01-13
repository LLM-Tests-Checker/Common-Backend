package create_test

import (
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type testCreator interface {
	CreateTest(authorID users.UserId, data tests.CreateTestData) (*tests.Test, error)
}

type testMapper interface {
	MapModelToDto(model *tests.Test) dto.Test
}

type tokenParser interface {
	ParseUserId(accessToken string) (users.UserId, error)
}
