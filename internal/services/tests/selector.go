package tests

type SelectorService interface {
	SelectMyTests(currentUserId int32) (error, []Test)

	SelectTestById(currentUserId int32, testId string) (error, Test)
}
