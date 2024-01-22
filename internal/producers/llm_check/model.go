package llm_check

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
)

type CheckEvent struct {
	ModelCheckId  llm.ModelCheckId
	TestId        tests.TestId
	TestQuestions []CheckEventQuestion
}

type CheckEventQuestion struct {
	QuestionNumber  int32
	QuestionText    string
	QuestionAnswers []CheckEventQuestionAnswer
}

type CheckEventQuestionAnswer struct {
	AnswerNumber int32
	AnswerText   string
}
