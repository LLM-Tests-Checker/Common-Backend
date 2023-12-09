package auth

import "fmt"

type SignUpProvider interface {
	PerformSignUp(name, login, passwordHash string) (error, *User)
}

func NewSignUpProvider() SignUpProvider {
	return defaultSignUpProvider{}
}

type defaultSignUpProvider struct {
}

func (provider defaultSignUpProvider) PerformSignUp(name, login, passwordHash string) (error, *User) {
	return fmt.Errorf("not implemented yet"), nil
}
