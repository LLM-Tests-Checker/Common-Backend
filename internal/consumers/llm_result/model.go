package llm_result

const ResultEventMessageVersionHeader = "Check-Result-Version"
const ResultEventVersion = "1.0.0"

type ResultEvent struct {
	ModelCheckId string    `json:"model_check_id"`
	TestId       string    `json:"test_id"`
	ModelAnswers []Answers `json:"model_answers"`
}

type Answers struct {
	QuestionNumber int32 `json:"question_number"`
	AnswerNumber   int32 `json:"answer_number"`
}
