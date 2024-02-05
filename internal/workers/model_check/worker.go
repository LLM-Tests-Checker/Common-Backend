package model_check

import (
	"context"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/producers/llm_check"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/sirupsen/logrus"
	"time"
)

type Worker struct {
	logger           *logrus.Logger
	llmStorage       llmStorage
	testsStorage     testsStorage
	producer         producer
	currentSleepTime time.Duration
}

func NewWorker(
	logger *logrus.Logger,
	llmStorage llmStorage,
	testsStorage testsStorage,
	producer producer,
) *Worker {
	return &Worker{
		logger:           logger,
		llmStorage:       llmStorage,
		testsStorage:     testsStorage,
		producer:         producer,
		currentSleepTime: 0,
	}
}

func (worker *Worker) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			handledChecksCount, err := worker.work(ctx)
			if err != nil {
				worker.logger.Errorf("Worker tick finished with error: %s", err)
				worker.sleep(0, err)
				continue
			}

			worker.logger.Infof("Model checks handled count: %d", handledChecksCount)

			worker.sleep(handledChecksCount, nil)
		}
	}
}

// Method is not transactional. Updating database and producing event execute not consistently.
func (worker *Worker) work(ctx context.Context) (int32, error) {
	const maxModelChecksCount = 10

	notStartedModelChecks, err := worker.llmStorage.GetNotStartedModelChecks(ctx, maxModelChecksCount)
	if err != nil {
		return 0, error2.Wrap(err, "worker.llmStorage.GetNotStartedModelChecks")
	}

	targetTestsIds := make([]tests.TestId, 0, len(notStartedModelChecks))
	for _, modelCheck := range notStartedModelChecks {
		targetTestsIds = append(targetTestsIds, modelCheck.TargetTestId)
	}

	targetTests, err := worker.testsStorage.GetTestsByIds(ctx, targetTestsIds)
	if err != nil {
		return 0, error2.Wrap(err, "worker.testsStorage.GetTestsByIds")
	}

	testIdToTest := make(map[tests.TestId]tests.Test, len(targetTests))
	for _, targetTest := range targetTests {
		testIdToTest[targetTest.Identifier] = targetTest
	}

	events := make([]llm_check.CheckEvent, len(notStartedModelChecks))
	for i := range notStartedModelChecks {
		modelCheck := notStartedModelChecks[i]
		targetTest := testIdToTest[modelCheck.TargetTestId]

		events[i] = mapModelCheckToEvent(modelCheck, targetTest)
	}

	err = worker.producer.ProduceEvents(ctx, events)
	if err != nil {
		return 0, error2.Wrap(err, "worker.producer.ProduceEvents")
	}

	modelCheckIds := make([]llm.ModelCheckId, len(notStartedModelChecks))
	for i := range notStartedModelChecks {
		modelCheckIds[i] = notStartedModelChecks[i].Identifier
	}

	err = worker.llmStorage.UpdateModelChecksStatus(ctx, modelCheckIds, llm.StatusInProgress)
	if err != nil {
		return 0, error2.Wrap(err, "worker.llmStorage.UpdateModelChecksStatus")
	}

	return int32(len(events)), nil
}

func (worker *Worker) sleep(handledChecksCount int32, err error) {
	const maxSleepDuration = 5 * time.Second
	const minSleepDuration time.Duration = 0
	const sleepStep = 250 * time.Millisecond

	if err != nil {
		worker.currentSleepTime = maxSleepDuration
	} else {
		if handledChecksCount != 0 {
			worker.currentSleepTime = minSleepDuration
		} else {
			newSleepTime := worker.currentSleepTime + sleepStep
			newSleepTime = min(newSleepTime, maxSleepDuration)
			worker.currentSleepTime = newSleepTime
		}
	}

	worker.logger.Infof("Sleeping for %d", worker.currentSleepTime)

	time.Sleep(worker.currentSleepTime)
}

func mapModelCheckToEvent(
	modelCheck llm.ModelCheck,
	targetTest tests.Test,
) llm_check.CheckEvent {
	mapAnswersFn := func(answers []tests.QuestionAnswer) []llm_check.CheckEventQuestionAnswer {
		result := make([]llm_check.CheckEventQuestionAnswer, len(answers))
		for i := range answers {
			result[i] = llm_check.CheckEventQuestionAnswer{
				AnswerNumber: answers[i].Number,
				AnswerText:   answers[i].Text,
			}
		}
		return result
	}

	mapQuestionsFn := func(questions []tests.TestQuestion) []llm_check.CheckEventQuestion {
		result := make([]llm_check.CheckEventQuestion, len(questions))
		for i := range questions {
			result[i] = llm_check.CheckEventQuestion{
				QuestionNumber:  questions[i].Number,
				QuestionText:    questions[i].Text,
				QuestionAnswers: mapAnswersFn(questions[i].Answers),
			}
		}
		return result
	}

	eventTest := llm_check.CheckEventTest{
		Id:          targetTest.Identifier.String(),
		Name:        targetTest.Name,
		Description: targetTest.Description,
		Questions:   mapQuestionsFn(targetTest.Questions),
	}

	return llm_check.CheckEvent{
		ModelCheckId: modelCheck.Identifier.String(),
		LLMSlug:      modelCheck.ModelSlug,
		TargetTest:   eventTest,
	}
}
