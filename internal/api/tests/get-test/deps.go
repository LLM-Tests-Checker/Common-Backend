package get_test

import (
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type testSelector interface {
	GetTestById(userId users.UserId, testId tests.TestId) (*tests.Test, error)
}

type testMapper interface {
	MapModelToDto(model *tests.Test) dto.Test
}
