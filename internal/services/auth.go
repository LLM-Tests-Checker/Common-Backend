package services

import (
	"fmt"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/components/auth"
)

type AuthService struct {
	signInProvider auth.SignInProvider
	signUpProvider auth.SignUpProvider
	tokensProvider auth.TokensProvider
	userProvider   auth.UserProvider
}

func NewAuthService(
	signInProvider auth.SignInProvider,
	signUpProvider auth.SignUpProvider,
	tokensProvider auth.TokensProvider,
	userProvider auth.UserProvider,
) *AuthService {
	return &AuthService{
		signInProvider: signInProvider,
		signUpProvider: signUpProvider,
		tokensProvider: tokensProvider,
		userProvider:   userProvider,
	}
}

func (service *AuthService) UserSignIn(login, passwordHash string) (error, *auth.Tokens) {
	return fmt.Errorf("not implementer yet"), nil
}

func (service *AuthService) UserSignUp(name, login, passwordHash string) (error, *auth.Tokens) {
	return fmt.Errorf("not implemented yet"), nil
}

func (service *AuthService) RefreshAccessToken(refreshToken string) (error, string) {
	return fmt.Errorf("not implemented yet"), ""
}
