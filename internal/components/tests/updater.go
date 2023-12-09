package tests

import "fmt"

type Updater interface {
	CreateTest(currentUserId int32, testData CreateTestData) (error, *Test)

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

func NewUpdater() Updater {
	return defaultUpdater{}
}

type defaultUpdater struct {
}

func (updater defaultUpdater) CreateTest(currentUserId int32, testData CreateTestData) (error, *Test) {
	return fmt.Errorf("not implemented yet"), nil
}

func (updater defaultUpdater) DeleteTest(currentUserId int32, testId int32) error {
	return fmt.Errorf("not implemeted yet")
}
