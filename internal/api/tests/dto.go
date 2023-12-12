package tests

type CreateTestRequest struct {
	Name        string                      `json:"name"`
	Description *string                     `json:"description"`
	Questions   []CreateTestQuestionPayload `json:"questions"`
}

type CreateTestQuestionPayload struct {
	Number  int32                             `json:"number"`
	Text    string                            `json:"text"`
	Answers []CreateTestQuestionAnswerPayload `json:"answers"`
}

type CreateTestQuestionAnswerPayload struct {
	Number    int32  `json:"number"`
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}

type CreateTestResponse struct {
	Test TestDTO `json:"test"`
}

type GetMyTestsResponse struct {
	Tests []TestDTO `json:"tests"`
}

type GetTestByIdResponse struct {
	Test TestDTO `json:"test"`
}

type TestDTO struct {
	Identifier  string            `json:"identifier"`
	Name        string            `json:"name"`
	Description *string           `json:"description"`
	Questions   []TestQuestionDTO `json:"questions"`
}

type TestQuestionDTO struct {
	Number  int32               `json:"number"`
	Text    string              `json:"text"`
	Answers []QuestionAnswerDTO `json:"answers"`
}

type QuestionAnswerDTO struct {
	Number    int32  `json:"number"`
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}
