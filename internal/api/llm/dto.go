package llm

type LaunchLLMCheckRequest struct {
	LLMSlug string `json:"llm_slug"`
}

type GetLLMCheckStatusResponse struct {
	Statuses []GetLLMCheckStatusValue `json:"statuses"`
}

type GetLLMCheckStatusValue struct {
	LLMSlug string `json:"llm_slug"`
	Status  string `json:"status"`
}

type GetLLMCheckResultResponse struct {
	Results []GetLLMCheckResultValue `json:"results"`
}

type GetLLMCheckResultValue struct {
	LLMSlug string                       `json:"llm_slug"`
	Answers []GetLLMCheckResultLLMAnswer `json:"answers"`
}

type GetLLMCheckResultLLMAnswer struct {
	QuestionNumber       int32 `json:"question_number"`
	SelectedAnswerNumber int32 `json:"selected_answer_number"`
}
