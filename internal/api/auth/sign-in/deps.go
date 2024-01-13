package sign_in

import "github.com/LLM-Tests-Checker/Common-Backend/internal/services/auth"

type authenticator interface {
	PerformSignIn(login, passwordHash string) (*auth.UserTokens, error)
}

type inputValidator interface {
	ValidateLogin(login string) error
	ValidatePasswordHash(passwordHash string) error
}
