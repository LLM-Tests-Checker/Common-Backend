package services

import "github.com/LLM-Tests-Checker/Common-Backend/internal/components/auth"

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

func (service *AuthService) UserSignIn() {

}

func (service *AuthService) UserSignUp() {

}

func (service *AuthService) RefreshAccessToken() {

}
