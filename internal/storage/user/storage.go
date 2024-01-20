package user

import (
	"context"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

const userCollectionName = "user"

type Storage struct {
	logger     *logrus.Logger
	collection *mongo.Collection
}

func NewUserStorage(
	logger *logrus.Logger,
	database *mongo.Database,
) *Storage {
	return &Storage{
		logger:     logger,
		collection: database.Collection(userCollectionName),
	}
}

func (storage *Storage) CheckUserWithLoginNotExists(
	ctx context.Context,
	login string,
) (bool, error) {
	targetUser, err := storage.GetUserByLogin(ctx, login)
	if err != nil {
		return false, err
	}

	if targetUser == nil {
		return true, nil
	} else {
		return false, nil
	}
}

func (storage *Storage) CheckUserWithIdExists(
	ctx context.Context,
	userId users.UserId,
) (bool, error) {
	cursor, err := storage.collection.Find(
		ctx,
		bson.M{
			userFieldIdentifier: userId.Int32(),
		},
	)
	if err != nil {
		return false, wrapError(err, "Can't check user with id exists")
	}

	rawUsers := make([]user, 0)
	err = cursor.All(ctx, &rawUsers)
	if err != nil {
		return false, wrapError(err, "Can't get user by login")
	}

	if len(rawUsers) == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (storage *Storage) GetUserByLogin(
	ctx context.Context,
	login string,
) (*users.User, error) {
	cursor, err := storage.collection.Find(
		ctx,
		bson.M{
			userFieldLogin: login,
		},
	)
	if err != nil {
		return nil, wrapError(err, "Can't get user by login")
	}

	rawUsers := make([]user, 0)
	err = cursor.All(ctx, &rawUsers)
	if err != nil {
		return nil, wrapError(err, "Can't get user by login")
	}

	if len(rawUsers) == 0 {
		return nil, nil
	}

	rawUser := rawUsers[0]
	resultUser := convertRawToModel(rawUser)

	return &resultUser, nil
}

func (storage *Storage) CreateNewUser(
	ctx context.Context,
	name, login, passwordHash string,
) (*users.User, error) {
	rawUser := user{
		Identifier:   int32(uuid.New().ID()),
		Name:         name,
		Login:        login,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now().Format(time.RFC3339),
	}

	insertResult, err := storage.collection.InsertOne(ctx, rawUser)
	if err != nil {
		return nil, wrapError(err, "Can't create new user")
	}

	storage.logger.Debugf("Inseted user with id: %s", insertResult.InsertedID)

	createdUser := convertRawToModel(rawUser)

	return &createdUser, nil
}

func convertRawToModel(rawUser user) users.User {
	createdAt, err := time.Parse(time.RFC3339, rawUser.CreatedAt)
	if err != nil {
		panic(err)
	}
	return users.User{
		Identifier:   users.UserId(rawUser.Identifier),
		Name:         rawUser.Name,
		Login:        rawUser.Login,
		PasswordHash: rawUser.PasswordHash,
		CreatedAt:    createdAt,
	}
}

func wrapError(err error, message string) error2.BackendError {
	return error2.Wrap(
		err,
		error2.UnknownError,
		message,
		http.StatusInternalServerError,
	)
}
