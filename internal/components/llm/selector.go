package llm

import (
	"fmt"
	"time"
)

type StatusSelector interface {
	GetLLMCheckStatus(currentUserId int32, testId string) (error, *CheckStatus)
}

type ResultSelector interface {
	GetLLMCheckResult(currentUserId int32, testId string) (error, *CheckResult)
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

func NewStatusSelector() StatusSelector {
	return defaultStatusSelector{}
}

func NewResultSelector() ResultSelector {
	return defaultResultSelector{}
}

type defaultStatusSelector struct {
}

func (selector defaultStatusSelector) GetLLMCheckStatus(currentUserId int32, testId string) (error, *CheckStatus) {
	return fmt.Errorf("not implemented yet"), nil
}

type defaultResultSelector struct {
}

func (selector defaultResultSelector) GetLLMCheckResult(currentUserId int32, testId string) (error, *CheckResult) {
	return fmt.Errorf("not implemented yet"), nil
}
