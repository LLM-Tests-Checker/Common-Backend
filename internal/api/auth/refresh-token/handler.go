package refresh_token

import (
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/platform/logger"
	"net/http"
)

type Handler struct {
	logger         logger.Logger
	tokenRefresher tokenRefresher
}

func New(
	logger logger.Logger,
	tokenRefresher tokenRefresher,
) *Handler {
	return &Handler{
		logger:         logger,
		tokenRefresher: tokenRefresher,
	}
}

func (handler *Handler) AuthRefreshToken(response http.ResponseWriter, r *http.Request, params dto.AuthRefreshTokenParams) {
	refreshToken := params.XLLMCheckerRefreshToken
	if refreshToken == "" {
		err := error2.NewBackendError(
			error2.InvalidRefreshToken,
			"Refresh token is missing",
			http.StatusUnauthorized,
		)
		http2.ReturnError(response, err)
		return
	}

	accessToken, err := handler.tokenRefresher.RefreshAccessToken(r.Context(), refreshToken)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	response.Header().Set(http2.AccessTokenHeaderName, accessToken)
}
