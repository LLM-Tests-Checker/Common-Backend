package auth

import (
	"errors"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type Service struct {
}

func NewAuthService() *Service {
	return &Service{}
}

func (service *Service) RefreshAccessToken(
	refreshToken string,
) (string, error) {
	return "", errors.New("not implemented yet")
}

func (service *Service) PerformSignIn(
	login, passwordHash string,
) (*UserTokens, error) {
	return nil, errors.New("not implemented yet")
}

func (service *Service) PerformSignUp(
	name, login, passwordHash string,
) (*UserTokens, error) {
	return nil, errors.New("not implemented yet")
}

func (service *Service) ParseUserId(accessToken string) (users.UserId, error) {
	return 0, errors.New("not implemented yet")
}
