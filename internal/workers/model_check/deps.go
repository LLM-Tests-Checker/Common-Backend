package model_check

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/producers/llm_check"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
)

type testsStorage interface {
	GetTestsByIds(
		ctx context.Context,
		testIds []tests.TestId,
	) ([]tests.Test, error)
}

type llmStorage interface {
	GetNotStartedModelChecks(
		ctx context.Context,
		maxCount int32,
	) ([]llm.ModelCheck, error)

	UpdateModelChecksStatus(
		ctx context.Context,
		modelCheckIds []llm.ModelCheckId,
		newStatus llm.CheckStatus,
	) error
}

type producer interface {
	ProduceEvents(ctx context.Context, events []llm_check.CheckEvent) error
}
