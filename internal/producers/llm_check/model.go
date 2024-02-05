package llm_check

const CheckEventMessageVersionHeader = "Check-Event-Version"
const CheckEventVersion = "1.0.0"

type CheckEvent struct {
	ModelCheckId string         `json:"id"`
	LLMSlug      string         `json:"llm_slug"`
	TargetTest   CheckEventTest `json:"test"`
}

type CheckEventTest struct {
	Id          string               `json:"id"`
	Name        string               `json:"name"`
	Description *string              `json:"description"`
	Questions   []CheckEventQuestion `json:"questions"`
}

type CheckEventQuestion struct {
	QuestionNumber  int32                      `json:"number"`
	QuestionText    string                     `json:"text"`
	QuestionAnswers []CheckEventQuestionAnswer `json:"answers"`
}

type CheckEventQuestionAnswer struct {
	AnswerNumber int32  `json:"number"`
	AnswerText   string `json:"text"`
}
