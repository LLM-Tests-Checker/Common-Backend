package user

import (
	"context"
	"errors"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	logger   *logrus.Logger
	database *mongo.Database
}

func NewUserStorage(
	logger *logrus.Logger,
	database *mongo.Database,
) *Storage {
	return &Storage{
		logger:   logger,
		database: database,
	}
}

func (storage *Storage) CheckUserWithLoginNotExists(
	ctx context.Context,
	login string,
) (bool, error) {
	return true, errors.New("not implemented yet")
}

func (storage *Storage) CheckUserWithIdExists(
	ctx context.Context,
	userId users.UserId,
) (bool, error) {
	return true, errors.New("not implemented yet")
}

func (storage *Storage) GetUserByLogin(
	ctx context.Context,
	login string,
) (*users.User, error) {
	return nil, errors.New("not implemented yet")
}

func (storage *Storage) CreateNewUser(
	ctx context.Context,
	name, login, passwordHash string,
) (users.User, error) {
	return users.User{}, errors.New("not implemented yet")
}
