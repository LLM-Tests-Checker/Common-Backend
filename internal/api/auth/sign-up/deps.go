package sign_up

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/auth"
)

type authenticator interface {
	PerformSignUp(ctx context.Context, name, login, passwordHash string) (*auth.UserTokens, error)
}

type inputValidator interface {
	ValidateName(name string) error
	ValidateLogin(login string) error
	ValidatePasswordHash(passwordHash string) error
}
