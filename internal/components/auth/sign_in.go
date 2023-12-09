package auth

import "fmt"

type SignInProvider interface {
	PerformSignIn(login, passwordHash string) (error, *User)
}

func NewSignInProvider() SignInProvider {
	return defaultSignInProvider{}
}

type defaultSignInProvider struct {
}

func (provider defaultSignInProvider) PerformSignIn(login, passwordHash string) (error, *User) {
	return fmt.Errorf("not implemented yet"), nil
}
