package sign_in

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/auth"
)

type authenticator interface {
	PerformSignIn(ctx context.Context, login, passwordHash string) (*auth.UserTokens, error)
}

type inputValidator interface {
	ValidateLogin(login string) error
	ValidatePasswordHash(passwordHash string) error
}
