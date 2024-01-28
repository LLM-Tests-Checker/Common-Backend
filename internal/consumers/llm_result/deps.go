package llm_result

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
)

type llmStorage interface {
	SetLLMCheckCompleted(
		ctx context.Context,
		checkId llm.ModelCheckId,
		modelAnswers []llm.ModelTestAnswer,
	) error
}
