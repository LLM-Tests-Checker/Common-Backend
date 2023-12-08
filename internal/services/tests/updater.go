package tests

type UpdaterService interface {
	CreateTest(currentUserId int32, testData CreateTestData) (error, Test)

	DeleteTest(currentUserId int32, testId int32) error
}

type CreateTestData struct {
	Name        string
	Description *string
	Questions   []CreateTestQuestionData
}

type CreateTestQuestionData struct {
	Number  int32
	Text    string
	Answers []CreateQuestionAnswerData
}

type CreateQuestionAnswerData struct {
	Number    int32
	Text      string
	IsCorrect bool
}
