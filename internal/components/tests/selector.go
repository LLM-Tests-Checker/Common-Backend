package tests

import "fmt"

type Selector interface {
	SelectMyTests(currentUserId int32) (error, []*Test)

	SelectTestById(currentUserId int32, testId string) (error, *Test)
}

func NewSelector() Selector {
	return defaultSelector{}
}

type defaultSelector struct {
}

func (selector defaultSelector) SelectMyTests(currentUserId int32) (error, []*Test) {
	return fmt.Errorf("not implemeted yet"), nil
}

func (selector defaultSelector) SelectTestById(currentUserId int32, testId string) (error, *Test) {
	return fmt.Errorf("not implemeted yet"), nil
}
