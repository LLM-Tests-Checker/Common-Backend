package auth

import (
	"encoding/json"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/common"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/api/constants"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/components/auth"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
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
	var requestBody SignInRequest
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		apiError := common.ApiError{
			ErrorCode:    constants.ErrorInvalidCredentials,
			ErrorMessage: err.Error(),
		}
		http2.ReturnApiError(responseWriter, apiError, http.StatusBadRequest)
		return
	}

	err, tokens := authService.UserSignIn(requestBody.UserLogin, requestBody.UserPasswordHash)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusBadRequest)
		return
	}

	responseWriter.Header().Set(constants.AccessTokenHeaderName, tokens.AccessToken)
	responseWriter.Header().Set(constants.RefreshTokenHeaderName, tokens.RefreshToken)
	responseWriter.WriteHeader(http.StatusOK)
}

func SignUpHandler(responseWriter http.ResponseWriter, request *http.Request) {
	var requestBody SignUpRequest
	err := json.NewDecoder(request.Body).Decode(&requestBody)
	if err != nil {
		apiError := common.ApiError{
			ErrorCode:    constants.ErrorInvalidCredentials,
			ErrorMessage: err.Error(),
		}
		http2.ReturnApiError(responseWriter, apiError, http.StatusBadRequest)
		return
	}

	err, tokens := authService.UserSignUp(requestBody.UserName, requestBody.UserLogin, requestBody.UserPasswordHash)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusBadRequest)
	}

	responseWriter.Header().Set(constants.AccessTokenHeaderName, tokens.AccessToken)
	responseWriter.Header().Set(constants.RefreshTokenHeaderName, tokens.RefreshToken)
	responseWriter.WriteHeader(http.StatusOK)
}

func RefreshAccessTokenHandler(responseWriter http.ResponseWriter, request *http.Request) {
	refreshToken := request.Header.Get(constants.RefreshTokenHeaderName)
	if refreshToken == "" {
		apiError := common.ApiError{
			ErrorCode:    constants.ErrorInvalidRefreshToken,
			ErrorMessage: "Refresh token is missing",
		}
		http2.ReturnApiError(responseWriter, apiError, http.StatusBadRequest)
		return
	}

	err, newAccessToken := authService.RefreshAccessToken(refreshToken)
	if err != nil {
		http2.ReturnError(responseWriter, err, http.StatusBadRequest)
		return
	}

	responseWriter.Header().Set(constants.AccessTokenHeaderName, newAccessToken)
	responseWriter.WriteHeader(http.StatusOK)
}
