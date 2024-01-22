package llm

import (
	"context"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/llm"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/tests"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

const llmModelCheckCollectionName = "model_check"

type Storage struct {
	logger     *logrus.Logger
	collection *mongo.Collection
}

func NewLLMStorage(
	logger *logrus.Logger,
	database *mongo.Database,
) *Storage {
	return &Storage{
		logger:     logger,
		collection: database.Collection(llmModelCheckCollectionName),
	}
}

func (storage *Storage) GetLLMChecksByTestId(
	ctx context.Context,
	testId tests.TestId,
) ([]llm.ModelCheck, error) {
	cursor, err := storage.collection.Find(
		ctx,
		bson.M{
			modelCheckFieldTargetTestId: testId.String(),
		},
	)
	if err != nil {
		return nil, wrapError(err, "Can't get llm checks by test id")
	}

	rawModelChecks := make([]modelCheck, 0)
	for cursor.Next(ctx) {
		var rawModelCheck modelCheck
		err := cursor.Decode(&rawModelCheck)
		if err != nil {
			return nil, wrapError(err, "Can't get llm checks by test id")
		}
		rawModelChecks = append(rawModelChecks, rawModelCheck)
	}

	resultModelChecks := make([]llm.ModelCheck, len(rawModelChecks))
	for i := range resultModelChecks {
		resultModelChecks[i] = convertRawToModel(rawModelChecks[i])
	}

	return resultModelChecks, nil
}

func (storage *Storage) InsertNotStartedLLMCheck(
	ctx context.Context,
	modelSlug llm.ModelSlug,
	testId tests.TestId,
	authorId users.UserId,
) (*llm.ModelCheck, error) {
	now := time.Now().Format(time.RFC3339)
	rawModelCheck := modelCheck{
		Identifier:   uuid.New().String(),
		Slug:         modelSlug,
		TargetTestId: testId.String(),
		AuthorId:     authorId.Int32(),
		Status:       llm.StatusNotStarted,
		Answers:      nil,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	insertResult, err := storage.collection.InsertOne(ctx, rawModelCheck)
	if err != nil {
		return nil, wrapError(err, "Can't insert not started LLM Check")
	}

	storage.logger.Debugf("Inserted model check with id: %s", insertResult.InsertedID)

	insertedModelCheck := convertRawToModel(rawModelCheck)

	return &insertedModelCheck, nil
}

func convertRawToModel(rawModelCheck modelCheck) llm.ModelCheck {
	mapModelAnswersFn := func(rawAnswers []modelAnswer) []llm.ModelTestAnswer {
		result := make([]llm.ModelTestAnswer, len(rawAnswers))
		for i := range rawAnswers {
			result[i] = llm.ModelTestAnswer{
				QuestionNumber:       rawAnswers[i].QuestionNumber,
				SelectedAnswerNumber: rawAnswers[i].SelectedAnswerNumber,
			}
		}
		return result
	}

	return llm.ModelCheck{
		Identifier:   uuid.MustParse(rawModelCheck.Identifier),
		ModelSlug:    rawModelCheck.Slug,
		TargetTestId: uuid.MustParse(rawModelCheck.TargetTestId),
		AuthorId:     users.UserId(rawModelCheck.AuthorId),
		Status:       rawModelCheck.Status,
		Answers:      mapModelAnswersFn(rawModelCheck.Answers),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}
}

func wrapError(err error, message string) error2.BackendError {
	return error2.WrapError(
		err,
		error2.UnknownError,
		message,
		http.StatusInternalServerError,
	)
}
