package services

import "github.com/LLM-Tests-Checker/Common-Backend/internal/components/tests"

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

func (service *TestsService) GetMyTests() {

}

func (service *TestsService) GetTestById() {

}

func (service *TestsService) CreateTest() {

}

func (service *TestsService) DeleteTestById() {

}
