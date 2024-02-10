package drop_old_in_progress_llm_check

import (
	"context"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/platform/logger"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"time"
)

type Worker struct {
	logger     logger.Logger
	llmStorage llmStorage
}

func NewWorker(
	logger logger.Logger,
	llmStorage llmStorage,
) *Worker {
	return &Worker{
		logger:     logger,
		llmStorage: llmStorage,
	}
}

func (worker *Worker) Start(ctx context.Context) error {
	const maxModelChecksCount = 100
	const dayOffset = -1

	timeBefore := time.Now().AddDate(0, 0, dayOffset)
	llmChecks, err := worker.llmStorage.GetInProgressModelChecks(ctx, timeBefore, maxModelChecksCount)
	if err != nil {
		return error2.Wrap(err, "worker.llmStorage.GetInProgressModelChecks")
	}

	worker.logger.Infof("Received %d old model checks with status in progress", len(llmChecks))

	checksIds := make([]llm.ModelCheckId, len(llmChecks))
	for i := range llmChecks {
		checksIds[i] = llmChecks[i].Identifier
	}

	err = worker.llmStorage.UpdateModelChecksStatus(ctx, checksIds, llm.StatusError)
	if err != nil {
		return error2.Wrap(err, "worker.llmStorage.UpdateModelChecksStatus")
	}

	return nil
}
