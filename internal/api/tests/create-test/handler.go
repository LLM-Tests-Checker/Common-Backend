package create_test

import (
	"encoding/json"
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	http2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/http"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/platform/logger"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"net/http"
)

type Handler struct {
	logger      logger.Logger
	creator     testCreator
	mapper      testMapper
	tokenParser tokenParser
}

func New(
	logger logger.Logger,
	creator testCreator,
	mapper testMapper,
	tokenParser tokenParser,
) *Handler {
	return &Handler{
		logger:      logger,
		creator:     creator,
		mapper:      mapper,
		tokenParser: tokenParser,
	}
}

func (handler *Handler) TestCreate(response http.ResponseWriter, r *http.Request) {
	request := dto.CreateTestRequest{}

	userId, err := http2.GetUserIdFromAccessToken(r, handler.tokenParser)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http2.ReturnErrorWithStatusCode(response, http.StatusBadRequest, err)
		return
	}

	createData := mapRequestDtoToModel(request)
	createdTest, err := handler.creator.CreateTest(r.Context(), userId, createData)
	if err != nil {
		http2.ReturnError(response, err)
		return
	}

	responseDto := handler.mapper.MapModelToDto(createdTest)
	err = json.NewEncoder(response).Encode(responseDto)
	if err != nil {
		http2.ReturnErrorWithStatusCode(response, http.StatusBadRequest, err)
		return
	}
}

func mapRequestDtoToModel(createDto dto.CreateTestRequest) tests.CreateTestData {

	answersMapFn := func(answerDto []dto.CreateTestQuestionAnswerPayload) []tests.CreateQuestionAnswerData {
		result := make([]tests.CreateQuestionAnswerData, len(answerDto))

		for i := range answerDto {
			answerData := tests.CreateQuestionAnswerData{
				Number:    int32(answerDto[i].Number),
				Text:      answerDto[i].Text,
				IsCorrect: answerDto[i].IsCorrect,
			}

			result[i] = answerData
		}

		return result
	}

	questionsMapFn := func(questionDto []dto.CreateTestQuestionPayload) []tests.CreateTestQuestionData {
		result := make([]tests.CreateTestQuestionData, len(questionDto))

		for i := range questionDto {
			questionData := tests.CreateTestQuestionData{
				Number:  int32(questionDto[i].Number),
				Text:    questionDto[i].Text,
				Answers: answersMapFn(questionDto[i].Answers),
			}

			result[i] = questionData
		}

		return result
	}

	return tests.CreateTestData{
		Name:        createDto.Name,
		Description: createDto.Description,
		Questions:   questionsMapFn(createDto.Questions),
	}
}
