package services

import (
	"errors"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/components/tests"
)

type TestsService struct {
	selector tests.Selector
	updater  tests.Updater
}

func NewTestsService(
	selector tests.Selector,
	updater tests.Updater,
) *TestsService {
	return &TestsService{
		selector: selector,
		updater:  updater,
	}
}

func (service *TestsService) GetMyTests(currentUserId int32) (error, []tests.Test) {
	return errors.New("not implemented yet"), nil
}

func (service *TestsService) GetTestById(currentUserId int32, testId string) (error, *tests.Test) {
	return errors.New("not implemented yet"), nil
}

func (service *TestsService) CreateTest(currentUserId int32, createTestData tests.CreateTestData) (error, *tests.Test) {
	return errors.New("not implemented tet"), nil
}

func (service *TestsService) DeleteTestById(currentUserId int32, testId string) error {
	return errors.New("not implemented yet")
}
