package users

import (
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	"net/http"
	"regexp"
)

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

var symbolsRegexp = regexp.MustCompile(`^[A-Za-z0-9]+$`)

func (validator *Validator) ValidateLogin(login string) error {
	const minLength = 3
	const maxLength = 20

	if len(login) < minLength || len(login) > maxLength {
		return error2.NewBackendError(
			error2.InputValidationError,
			"Invalid login length",
			http.StatusBadRequest,
		)
	}

	matched := symbolsRegexp.MatchString(login)
	if !matched {
		return error2.NewBackendError(
			error2.InputValidationError,
			"Invalid login symbols",
			http.StatusBadRequest,
		)
	}

	return nil
}

func (validator *Validator) ValidatePasswordHash(passwordHash string) error {
	const length = 128
	if len(passwordHash) != length {
		return error2.NewBackendError(
			error2.InputValidationError,
			"Invalid password hash length",
			http.StatusBadRequest,
		)
	}

	return nil
}
