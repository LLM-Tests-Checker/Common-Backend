package tests

const (
	testFieldIdentifier = "identifier"
	testFieldAuthorId   = "author_id"
)

type test struct {
	Identifier  string  `bson:"identifier"`
	Name        string  `bson:"name"`
	Description *string `bson:"description,omitempty"`
	Questions   []testQuestion
	CreatedAt   string `bson:"created_at"`
	AuthorId    int32  `bson:"author_id"`
}

type testQuestion struct {
	Number  int32            `bson:"number"`
	Text    string           `bson:"text"`
	Answers []questionAnswer `bson:"answers"`
}

type questionAnswer struct {
	Number    int32  `bson:"number"`
	Text      string `bson:"text"`
	IsCorrect bool   `bson:"is_correct"`
}
