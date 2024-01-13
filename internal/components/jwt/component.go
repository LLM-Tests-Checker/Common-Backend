package jwt

import (
	"errors"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type Component struct {
	config Config
}

func NewJWTComponent(
	config Config,
) *Component {
	return &Component{
		config: config,
	}
}

func (component *Component) GenerateAccessToken(
	userId users.UserId,
) (string, error) {
	return "", errors.New("not implemented yet")
}

func (component *Component) GenerateRefreshToken(
	userId users.UserId,
) (string, error) {
	return "", errors.New("not implemented yet")
}

func (component *Component) ValidateAndParseAccessToken(
	accessToken string,
) (users.UserId, error) {
	return 0, errors.New("not implemented yet")
}

func (component *Component) ValidateAndParseRefreshToken(
	refreshToken string,
) (users.UserId, error) {
	return 0, errors.New("not implemented yet")
}
