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
	options2 "go.mongodb.org/mongo-driver/mongo/options"
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

func (storage *Storage) GetNotStartedModelChecks(
	ctx context.Context,
	maxCount int32,
) ([]llm.ModelCheck, error) {
	limit := int64(maxCount)
	options := options2.FindOptions{
		Limit: &limit,
	}
	options.SetSort(bson.D{
		{modelCheckFieldCreatedAt, 1},
	})

	cursor, err := storage.collection.Find(
		ctx,
		bson.M{
			modelCheckFieldStatus: llm.StatusNotStarted,
		},
		&options,
	)
	if err != nil {
		return nil, wrapError(err, "Can't get not started model checks")
	}

	rawModelChecks := make([]modelCheck, 0, maxCount)
	err = cursor.All(ctx, &rawModelChecks)
	if err != nil {
		return nil, wrapError(err, "Can't get not started model checks")
	}

	resultModelChecks := make([]llm.ModelCheck, len(rawModelChecks))
	for i := range rawModelChecks {
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

func (storage *Storage) UpdateModelChecksStatus(
	ctx context.Context,
	modelCheckIds []llm.ModelCheckId,
	newStatus llm.CheckStatus,
) error {
	modelCheckIdsString := make([]string, len(modelCheckIds))
	for i := range modelCheckIds {
		modelCheckIdsString[i] = modelCheckIds[i].String()
	}

	updateResult, err := storage.collection.UpdateMany(
		ctx,
		bson.D{
			{
				modelCheckFieldIdentifier,
				bson.M{
					"$in": modelCheckIdsString,
				},
			},
		},
		bson.D{
			{
				"$set",
				bson.D{
					{modelCheckFieldStatus, newStatus},
				},
			},
		},
	)
	if err != nil {
		return wrapError(err, "Can't update model checks status")
	}

	storage.logger.Debugf("Updated model checks status count: %d", updateResult.ModifiedCount)

	return nil
}

func (storage *Storage) SetLLMCheckCompleted(
	ctx context.Context,
	checkId llm.ModelCheckId,
	modelAnswers []llm.ModelTestAnswer,
) error {
	checkIdString := checkId.String()
	rawAnswers := make([]modelAnswer, len(modelAnswers))
	for i := range modelAnswers {
		rawAnswers[i] = modelAnswer{
			QuestionNumber:       modelAnswers[i].QuestionNumber,
			SelectedAnswerNumber: modelAnswers[i].SelectedAnswerNumber,
		}
	}

	updateResult, err := storage.collection.UpdateOne(
		ctx,
		bson.D{
			{
				modelCheckFieldIdentifier,
				checkIdString,
			},
		},
		bson.D{
			{
				"$set",
				bson.D{
					{modelCheckFieldStatus, llm.StatusCompleted},
				},
			},
			{
				"$set",
				bson.D{
					{modelCheckFieldAnswers, rawAnswers},
				},
			},
		},
	)
	if err != nil {
		return wrapError(err, "Can't set llm check completed")
	}

	storage.logger.Debugf("Set completed model checks count: %d", updateResult.ModifiedCount)

	return nil
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
