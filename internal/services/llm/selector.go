package llm

import "time"

type StatusSelector interface {
	GetLLMCheckStatus(currentUserId int32, testId string) (error, CheckStatus)
}

type ResultSelector interface {
	GetLLMCheckResult(currentUserId int32, testId string) (error, CheckResult)
}

type CheckStatus struct {
	Statuses []ModelStatus
}

type ModelStatus struct {
	LLMSlug string
	Status  string
}

type CheckResult struct {
	Identifier   string
	TargetTestId string
	ModelResults []ModelResult
}

type ModelResult struct {
	LLMSlug   string
	CreatedAt time.Time
	Results   []ModelQuestionResult
}

type ModelQuestionResult struct {
	QuestionNumber       int32
	SelectedAnswerNumber int32
}
