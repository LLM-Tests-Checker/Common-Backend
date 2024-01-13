package tests

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"time"
)

type TestId = string

type Test struct {
	Identifier  TestId
	Name        string
	Description *string
	Questions   []TestQuestion

	CreatedAt time.Time
	AuthorId  users.UserId
}

type TestQuestion struct {
	Number  int32
	Text    string
	Answers []QuestionAnswer
}

type QuestionAnswer struct {
	Number    int32
	Text      string
	IsCorrect bool
}

type CreateTestData struct {
	Name        string
	Description *string
	Questions   []CreateTestQuestionData
}

type CreateTestQuestionData struct {
	Number  int32
	Text    string
	Answers []CreateQuestionAnswerData
}

type CreateQuestionAnswerData struct {
	Number    int32
	Text      string
	IsCorrect bool
}
