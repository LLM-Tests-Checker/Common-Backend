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
