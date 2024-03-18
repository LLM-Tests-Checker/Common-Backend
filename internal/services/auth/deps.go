package auth

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type userStorage interface {
	CheckUserWithLoginNotExists(ctx context.Context, login string) (bool, error)

	CheckUserWithIdExists(ctx context.Context, userId users.UserId) (bool, error)

	GetUserByLogin(ctx context.Context, login string) (*users.User, error)

	CreateNewUser(ctx context.Context, login, passwordHash string) (*users.User, error)
}

type jwtProvider interface {
	GenerateAccessToken(userId users.UserId) (string, error)

	GenerateRefreshToken(userId users.UserId) (string, error)

	ValidateAndParseAccessToken(accessToken string) (users.UserId, error)

	ValidateAndParseRefreshToken(refreshToken string) (users.UserId, error)
}
