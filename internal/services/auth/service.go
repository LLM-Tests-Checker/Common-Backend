package auth

import (
	"context"
	"fmt"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"net/http"
)

type Service struct {
	userStorage userStorage
	jwtProvider jwtProvider
}

func NewAuthService(
	userStorage userStorage,
	jwtProvider jwtProvider,
) *Service {
	return &Service{
		userStorage: userStorage,
		jwtProvider: jwtProvider,
	}
}

func (service *Service) RefreshAccessToken(
	ctx context.Context,
	refreshToken string,
) (string, error) {
	userId, err := service.jwtProvider.ValidateAndParseRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	userExists, err := service.userStorage.CheckUserWithIdExists(ctx, userId)
	if err != nil {
		return "", err
	}

	if !userExists {
		err := error2.NewBackendError(
			error2.InvalidRefreshToken,
			"Refresh token is invalid",
			http.StatusUnauthorized,
		)
		return "", err
	}

	accessToken, err := service.jwtProvider.GenerateAccessToken(userId)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (service *Service) PerformSignIn(
	ctx context.Context,
	login, passwordHash string,
) (*UserTokens, error) {
	targetUser, err := service.userStorage.GetUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}
	if targetUser == nil {
		err := error2.NewBackendError(
			error2.UserNotFound,
			fmt.Sprintf("User with login %s not found", login),
			http.StatusBadRequest,
		)
		return nil, err
	}

	if targetUser.PasswordHash != passwordHash {
		err := error2.NewBackendError(
			error2.UserInvalidPassword,
			"Invalid password",
			http.StatusBadRequest,
		)
		return nil, err
	}

	accessToken, err := service.jwtProvider.GenerateAccessToken(targetUser.Identifier)
	if err != nil {
		return nil, err
	}
	refreshToken, err := service.jwtProvider.GenerateRefreshToken(targetUser.Identifier)
	if err != nil {
		return nil, err
	}

	userTokens := UserTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &userTokens, nil
}

func (service *Service) PerformSignUp(
	ctx context.Context,
	name, login, passwordHash string,
) (*UserTokens, error) {
	notExists, err := service.userStorage.CheckUserWithLoginNotExists(ctx, login)
	if err != nil {
		return nil, err
	}
	if !notExists {
		err := error2.NewBackendError(
			error2.UserWithLoginAlreadyExists,
			fmt.Sprintf("User with login %s already exists", login),
			http.StatusBadRequest,
		)
		return nil, err
	}

	createdUser, err := service.userStorage.CreateNewUser(ctx, name, login, passwordHash)
	if err != nil {
		return nil, err
	}

	accessToken, err := service.jwtProvider.GenerateAccessToken(createdUser.Identifier)
	if err != nil {
		return nil, err
	}
	refreshToken, err := service.jwtProvider.GenerateRefreshToken(createdUser.Identifier)
	if err != nil {
		return nil, err
	}

	userTokens := UserTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &userTokens, nil
}

func (service *Service) ParseUserId(
	ctx context.Context,
	accessToken string,
) (users.UserId, error) {
	userId, err := service.jwtProvider.ValidateAndParseAccessToken(accessToken)
	if err != nil {
		return 0, err
	}

	userExists, err := service.userStorage.CheckUserWithIdExists(ctx, userId)
	if err != nil {
		return 0, err
	}

	if !userExists {
		err := error2.NewBackendError(
			error2.InvalidAccessToken,
			"Access token is invalid",
			http.StatusUnauthorized,
		)
		return 0, err
	}

	return userId, nil
}
