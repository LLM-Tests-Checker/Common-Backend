package drop_old_in_progress_llm_check

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"time"
)

type llmStorage interface {
	GetInProgressModelChecks(
		ctx context.Context,
		updatedLaterThen time.Time,
		maxCount int32,
	) ([]llm.ModelCheck, error)

	UpdateModelChecksStatus(
		ctx context.Context,
		modelCheckIds []llm.ModelCheckId,
		newStatus llm.CheckStatus,
	) error
}
