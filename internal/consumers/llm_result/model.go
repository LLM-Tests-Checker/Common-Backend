package llm_result

const ResultEventMessageVersionHeader = "Check-Result-Version"
const ResultEventVersion = "1.0.0"

type ResultEvent struct {
	ModelCheckId string    `json:"id"`
	TestId       string    `json:"target_test_id"`
	ModelAnswers []Answers `json:"answers"`
}

type Answers struct {
	QuestionNumber int32 `json:"question_number"`
	AnswerNumber   int32 `json:"selected_answer_number"`
}
