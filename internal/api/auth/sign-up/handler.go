package sign_up

import (
	"encoding/json"
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/platform/logger"
	"net/http"
)

type Handler struct {
	logger         logger.Logger
	authenticator  authenticator
	inputValidator inputValidator
}

func New(
	logger logger.Logger,
	authenticator authenticator,
	inputValidator inputValidator,
) *Handler {
	return &Handler{
		logger:         logger,
		authenticator:  authenticator,
		inputValidator: inputValidator,
	}
}

func (handler *Handler) AuthSignUp(response http.ResponseWriter, r *http.Request) {
	request := dto.SignUpRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http2.ReturnErrorWithStatusCode(response, http.StatusBadRequest, err)
		return
	}

	err = handler.validateInput(request)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	userTokens, err := handler.authenticator.PerformSignUp(
		r.Context(),
		request.UserName,
		request.UserLogin,
		request.UserPasswordHash,
	)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	response.Header().Set(http2.AccessTokenHeaderName, userTokens.AccessToken)
	response.Header().Set(http2.RefreshTokenHeaderName, userTokens.RefreshToken)
}

func (handler *Handler) validateInput(request dto.SignUpRequest) error {
	err := handler.inputValidator.ValidateName(request.UserName)
	if err != nil {
		return err
	}

	err = handler.inputValidator.ValidateLogin(request.UserLogin)
	if err != nil {
		return err
	}

	err = handler.inputValidator.ValidatePasswordHash(request.UserPasswordHash)
	if err != nil {
		return err
	}

	return nil
}
