package llm_check

type CheckEvent struct {
	ModelCheckId  string               `json:"model_check_id"`
	TestId        string               `json:"test_id"`
	TestQuestions []CheckEventQuestion `json:"test_questions"`
}

type CheckEventQuestion struct {
	QuestionNumber  int32                      `json:"question_number"`
	QuestionText    string                     `json:"question_text"`
	QuestionAnswers []CheckEventQuestionAnswer `json:"question_answers"`
}

type CheckEventQuestionAnswer struct {
	AnswerNumber int32  `json:"answer_number"`
	AnswerText   string `json:"answer_text"`
}
