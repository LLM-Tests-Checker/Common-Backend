package llm

const (
	modelCheckFieldTargetTestId = "target_test_id"
)

type modelCheck struct {
	Identifier   string        `bson:"identifier"`
	Slug         string        `bson:"slug"`
	TargetTestId string        `bson:"target_test_id"`
	AuthorId     int32         `bson:"author_id"`
	Status       string        `bson:"status"`
	Answers      []modelAnswer `bson:"answers,omitempty"`
	CreatedAt    string        `bson:"created_at"`
	UpdatedAt    string        `bson:"updated_at"`
}

type modelAnswer struct {
	QuestionNumber       int32 `bson:"question_number"`
	SelectedAnswerNumber int32 `bson:"selected_answer_number"`
}
