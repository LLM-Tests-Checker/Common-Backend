package mappers

import (
	dto "github.com/LLM-Tests-Checker/Common-Backend/internal/generated/schema"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
)

type TestMapper struct {
}

func NewTestMapper() *TestMapper {
	return &TestMapper{}
}

func (mapper *TestMapper) MapModelToDto(model *tests.Test) dto.Test {
	mapAnswersFn := func(answers []tests.QuestionAnswer) []dto.QuestionAnswer {
		result := make([]dto.QuestionAnswer, len(answers))

		for i := range answers {
			result[i] = dto.QuestionAnswer{
				Number:    int(answers[i].Number),
				Text:      answers[i].Text,
				IsCorrect: answers[i].IsCorrect,
			}
		}

		return result
	}

	mapQuestionsFn := func(questions []tests.TestQuestion) []dto.TestQuestion {
		result := make([]dto.TestQuestion, len(questions))

		for i := range questions {
			result[i] = dto.TestQuestion{
				Number:  int(questions[i].Number),
				Text:    questions[i].Text,
				Answers: mapAnswersFn(questions[i].Answers),
			}
		}

		return result
	}

	result := dto.Test{
		Identifier:  model.Identifier,
		Name:        model.Name,
		Description: model.Description,
		Questions:   mapQuestionsFn(model.Questions),
	}

	return result
}
