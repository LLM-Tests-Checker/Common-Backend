package auth

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/components/auth"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services"
	"net/http"
)

var (
	authService *services.AuthService
)

func init() {
	signInProvider := auth.NewSignInProvider()
	signUpProvider := auth.NewSignUpProvider()
	tokensProvider := auth.NewTokensProvider()
	userProvider := auth.NewUserProvider()

	authService = services.NewAuthService(signInProvider, signUpProvider, tokensProvider, userProvider)
}

func SignInHandler(responseWriter http.ResponseWriter, request *http.Request) {
	authService.UserSignIn()
}

func SignUpHandler(responseWriter http.ResponseWriter, request *http.Request) {
	authService.UserSignUp()
}

func RefreshAccessTokenHandler(w http.ResponseWriter, request *http.Request) {
	authService.RefreshAccessToken()
}
