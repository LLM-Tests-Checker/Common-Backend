package test

import (
	"context"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
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

const testCollectionName = "test"

type Storage struct {
	logger     *logrus.Logger
	collection *mongo.Collection
}

func NewTestsStorage(
	logger *logrus.Logger,
	database *mongo.Database,
) *Storage {
	return &Storage{
		logger:     logger,
		collection: database.Collection(testCollectionName),
	}
}

func (storage *Storage) GetTestById(
	ctx context.Context,
	testId tests.TestId,
) (*tests.Test, error) {
	cursor, err := storage.collection.Find(
		ctx,
		bson.M{
			testFieldIdentifier: testId.String(),
		},
	)
	if err != nil {
		return nil, wrapError(err, "Can't get test by id")
	}

	rawTests := make([]test, 0)
	err = cursor.All(ctx, &rawTests)
	if err != nil {
		return nil, wrapError(err, "Can't get test by id")
	}

	if len(rawTests) == 0 {
		return nil, nil
	}

	targetTest := convertRawToModel(rawTests[0])

	return &targetTest, nil
}

func (storage *Storage) GetTestsByIds(
	ctx context.Context,
	testsIds []tests.TestId,
) ([]tests.Test, error) {
	testsIdsString := make([]string, len(testsIds))
	for i := range testsIds {
		testsIdsString[i] = testsIds[i].String()
	}

	cursor, err := storage.collection.Find(
		ctx,
		bson.M{
			testFieldIdentifier: bson.M{
				"$in": testsIdsString,
			},
		},
	)
	if err != nil {
		return nil, wrapError(err, "Can't get tests by id")
	}

	rawTests := make([]test, 0, len(testsIds))
	err = cursor.All(ctx, &rawTests)
	if err != nil {
		return nil, wrapError(err, "Can't get tests by id")
	}

	resultTests := make([]tests.Test, len(rawTests))
	for i := range rawTests {
		resultTests[i] = convertRawToModel(rawTests[i])
	}

	return resultTests, nil
}

func (storage *Storage) GetTestsByAuthorId(
	ctx context.Context,
	authorId users.UserId,
	offset int32,
	size int32,
) ([]tests.Test, error) {
	limit := int64(size)
	skip := int64(offset)
	options := options2.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	}
	cursor, err := storage.collection.Find(
		ctx,
		bson.M{
			testFieldAuthorId: authorId,
		},
		&options,
	)
	if err != nil {
		return nil, wrapError(err, "Can't get tests by author id")
	}

	rawTests := make([]test, 0, size)
	for cursor.Next(ctx) {
		var rawTest test
		err := cursor.Decode(&rawTest)
		if err != nil {
			return nil, wrapError(err, "Can't get tests by author id")
		}
		rawTests = append(rawTests, rawTest)
	}

	resultTests := make([]tests.Test, len(rawTests))
	for i := range rawTests {
		resultTests[i] = convertRawToModel(rawTests[i])
	}

	return resultTests, nil
}

func (storage *Storage) CreateTest(
	ctx context.Context,
	authorId users.UserId,
	createData tests.CreateTestData,
) (*tests.Test, error) {
	answersMapFn := func(answers []tests.CreateQuestionAnswerData) []questionAnswer {
		rawAnswers := make([]questionAnswer, len(answers))
		for i := range answers {
			rawAnswers[i] = questionAnswer{
				Number:    answers[i].Number,
				Text:      answers[i].Text,
				IsCorrect: answers[i].IsCorrect,
			}
		}
		return rawAnswers
	}

	questionsMapFn := func(questions []tests.CreateTestQuestionData) []testQuestion {
		rawQuestions := make([]testQuestion, len(questions))
		for i := range questions {
			rawQuestions[i] = testQuestion{
				Number:  questions[i].Number,
				Text:    questions[i].Text,
				Answers: answersMapFn(questions[i].Answers),
			}
		}
		return rawQuestions
	}

	rawTest := test{
		Identifier:  uuid.New().String(),
		Name:        createData.Name,
		Description: createData.Description,
		Questions:   questionsMapFn(createData.Questions),
		CreatedAt:   time.Now().Format(time.RFC3339),
		AuthorId:    authorId.Int32(),
	}

	insertResult, err := storage.collection.InsertOne(ctx, rawTest)
	if err != nil {
		return nil, wrapError(err, "Can't create test")
	}
	storage.logger.Debugf("Created test with id: %s", insertResult.InsertedID)

	resultTest := convertRawToModel(rawTest)

	return &resultTest, nil
}

func (storage *Storage) DeleteTestById(
	ctx context.Context,
	testId tests.TestId,
) error {
	deleteResult, err := storage.collection.DeleteOne(
		ctx,
		bson.M{
			testFieldIdentifier: testId.String(),
		},
	)
	if err != nil {
		return wrapError(err, "Can't delete test by id")
	}

	storage.logger.Debugf("Deleted number of tests: %d", deleteResult.DeletedCount)

	return nil
}

func convertRawToModel(rawTest test) tests.Test {
	mapAnswersFn := func(rawAnswers []questionAnswer) []tests.QuestionAnswer {
		answers := make([]tests.QuestionAnswer, len(rawAnswers))
		for i := range rawAnswers {
			answers[i] = tests.QuestionAnswer{
				Number:    rawAnswers[i].Number,
				Text:      rawAnswers[i].Text,
				IsCorrect: rawAnswers[i].IsCorrect,
			}
		}
		return answers
	}

	mapQuestionsFn := func(rawQuestions []testQuestion) []tests.TestQuestion {
		questions := make([]tests.TestQuestion, len(rawQuestions))
		for i := range rawQuestions {
			questions[i] = tests.TestQuestion{
				Number:  rawQuestions[i].Number,
				Text:    rawQuestions[i].Text,
				Answers: mapAnswersFn(rawQuestions[i].Answers),
			}
		}
		return questions
	}

	createdAt, err := time.Parse(time.RFC3339, rawTest.CreatedAt)
	if err != nil {
		panic(err)
	}
	return tests.Test{
		Identifier:  uuid.MustParse(rawTest.Identifier),
		Name:        rawTest.Name,
		Description: rawTest.Description,
		Questions:   mapQuestionsFn(rawTest.Questions),
		CreatedAt:   createdAt,
		AuthorId:    users.UserId(rawTest.AuthorId),
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
