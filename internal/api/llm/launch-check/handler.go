package launch_check

import (
	"encoding/json"
	"fmt"
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	logger      *logrus.Logger
	launcher    checkLauncher
	tokenParser tokenParser
}

func New(
	logger *logrus.Logger,
	launcher checkLauncher,
	tokenParser tokenParser,
) *Handler {
	return &Handler{
		logger:      logger,
		launcher:    launcher,
		tokenParser: tokenParser,
	}
}

func (handler *Handler) LlmLaunch(response http.ResponseWriter, r *http.Request, testId dto.TestId) {
	userId, err := http2.GetUserIdFromAccessToken(r, handler.tokenParser)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	request := dto.LaunchLLMCheckRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http2.ReturnErrorWithStatusCode(response, http.StatusBadRequest, err)
		return
	}

	llmSlug, err := mapDtoSlugToModelSlug(request.LlmSlug)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	modelCheck, err := handler.launcher.LaunchModelCheck(r.Context(), userId, testId, llmSlug)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	responseCheck := dto.LaunchLLMCheckResponse{
		LaunchIdentifier: modelCheck.Identifier,
	}

	err = json.NewEncoder(response).Encode(responseCheck)
	if err != nil {
		http2.ReturnErrorWithStatusCode(response, http.StatusBadRequest, err)
		return
	}
}

func mapDtoSlugToModelSlug(dtoSlug dto.LaunchLLMCheckRequestLlmSlug) (llm.ModelSlug, error) {
	switch dtoSlug {
	case dto.GPT3:
		return llm.ModelGPT3, nil
	case dto.YandexGPT2:
		return llm.ModelYandexGPT, nil
	case dto.Dummy:
		return llm.Dummy, nil
	}

	return "", error2.NewBackendError(
		error2.InvalidLLMSlug,
		fmt.Sprintf("LLM slug %s not found", dtoSlug),
		http.StatusBadRequest,
	)
}
