package llm

import (
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"github.com/google/uuid"
	"time"
)

type ModelCheckId = uuid.UUID

type ModelSlug = string

type CheckStatus = string

const (
	ModelGigaChat ModelSlug = "gigachat"
	ModelGPT4     ModelSlug = "gpt4"
	Dummy         ModelSlug = "dummy"
)

const (
	StatusNotStarted CheckStatus = "NOT_STARTED"
	StatusInProgress CheckStatus = "IN_PROGRESS"
	StatusCompleted  CheckStatus = "COMPLETED"
	StatusError      CheckStatus = "ERROR"
)

type ModelCheck struct {
	Identifier   ModelCheckId
	ModelSlug    ModelSlug
	TargetTestId tests.TestId
	AuthorId     users.UserId
	Status       CheckStatus
	Answers      []ModelTestAnswer
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ModelCheckStatus struct {
	LLMSlug ModelSlug
	TestId  tests.TestId
	Status  CheckStatus
}

type ModelTestResult struct {
	LLMSlug    ModelSlug
	TestId     dto.TestId
	LLMAnswers []ModelTestAnswer
}

type ModelTestAnswer struct {
	QuestionNumber       int32
	SelectedAnswerNumber int32
}
