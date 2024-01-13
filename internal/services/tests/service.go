package tests

import (
	"errors"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type Service struct {
}

func NewTestsService() *Service {
	return &Service{}
}

func (service *Service) CreateTest(authorID users.UserId, data CreateTestData) (*Test, error) {
	return nil, errors.New("not implemented yet")
}

func (service *Service) DeleteTest(authorId users.UserId, testId TestId) error {
	return errors.New("not implemented yet")
}

func (service *Service) GetTestsByAuthorId(authorId users.UserId, pageNumber, pageSize int32) ([]Test, error) {
	return nil, errors.New("not implemented yet")
}

func (service *Service) GetTestById(userId users.UserId, testId TestId) (*Test, error) {
	return nil, errors.New("not implemented yet")
}
