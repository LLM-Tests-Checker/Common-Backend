package tests

import "time"

type Test struct {
	Identifier  string
	Name        string
	Description *string
	CreatedBy   int64
	CreatedAt   time.Time
	Questions   []TestQuestion
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
