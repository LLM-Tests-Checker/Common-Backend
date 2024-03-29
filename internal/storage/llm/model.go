package llm

const (
	modelCheckFieldIdentifier   = "identifier"
	modelCheckFieldTargetTestId = "target_test_id"
	modelCheckFieldCreatedAt    = "created_at"
	modelCheckFieldUpdatedAt    = "updated_at"
	modelCheckFieldStatus       = "status"
	modelCheckFieldAnswers      = "answers"
)

type modelCheck struct {
	Identifier   string        `bson:"identifier"`
	Slug         string        `bson:"slug"`
	TargetTestId string        `bson:"target_test_id"`
	AuthorId     int32         `bson:"author_id"`
	Status       string        `bson:"status"`
	Answers      []modelAnswer `bson:"answers,omitempty"`
	CreatedAt    int64         `bson:"created_at"`
	UpdatedAt    int64         `bson:"updated_at"`
}

type modelAnswer struct {
	QuestionNumber       int32 `bson:"question_number"`
	SelectedAnswerNumber int32 `bson:"selected_answer_number"`
}
