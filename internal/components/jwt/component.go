package jwt

import (
	"fmt"
	error2 "github.com/LLM-Tests-Checker/Common-Backend/internal/platform/error"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Component struct {
	config Config
}

var tokenSigningMethod = jwt.SigningMethodHS256

func NewJWTComponent(
	config Config,
) *Component {
	return &Component{
		config: config,
	}
}

func (component *Component) GenerateAccessToken(
	userId users.UserId,
) (string, error) {
	now := time.Now()
	expires := now.Add(component.config.AccessTokenLiveTime)
	tokenPayload := jwt.MapClaims{
		"iat": jwt.NewNumericDate(now),
		"exp": jwt.NewNumericDate(expires),
		"sub": strconv.Itoa(userId.Int()),
		"iss": component.config.Issuer,
	}

	token := jwt.NewWithClaims(tokenSigningMethod, tokenPayload)
	secretKey := []byte(component.config.AccessSecretKey)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		backendErr := error2.WrapError(
			err,
			error2.ServerError,
			"Failed to create access token",
			http.StatusInternalServerError,
		)
		return "", backendErr
	}

	return signedToken, nil
}

func (component *Component) GenerateRefreshToken(
	userId users.UserId,
) (string, error) {
	now := time.Now()
	expires := now.Add(component.config.RefreshTokenLiveTime)
	tokenPayload := jwt.MapClaims{
		"iat": jwt.NewNumericDate(now),
		"exp": jwt.NewNumericDate(expires),
		"sub": strconv.Itoa(userId.Int()),
		"iss": component.config.Issuer,
	}

	token := jwt.NewWithClaims(tokenSigningMethod, tokenPayload)
	secretKey := []byte(component.config.RefreshSecretKey)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		backendErr := error2.WrapError(
			err,
			error2.ServerError,
			"Failed to create refresh token",
			http.StatusInternalServerError,
		)
		return "", backendErr
	}

	return signedToken, nil
}

func (component *Component) ValidateAndParseAccessToken(
	accessToken string,
) (users.UserId, error) {
	options := []jwt.ParserOption{
		jwt.WithIssuedAt(),
		jwt.WithIssuer(component.config.Issuer),
	}

	parser := jwt.NewParser(options...)

	token, err := parser.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secretKey := []byte(component.config.AccessSecretKey)
		return secretKey, nil
	})
	if err != nil {
		var backendErr = error2.BackendError{}
		if strings.Contains(err.Error(), "token is expired") {
			backendErr = error2.WrapError(
				err,
				error2.AccessTokenExpired,
				"Access token expired",
				http.StatusUnauthorized,
			)
		} else {
			backendErr = error2.WrapError(
				err,
				error2.InvalidAccessToken,
				"Failed to parse access token",
				http.StatusUnauthorized,
			)
		}
		return 0, backendErr
	}

	expiredAt, err := token.Claims.GetExpirationTime()
	if err != nil {
		return 0, provideDefaultAccessTokenError(err)
	}

	if expiredAt.Before(time.Now()) {
		backendErr := error2.NewBackendError(
			error2.AccessTokenExpired,
			"Access token expired",
			http.StatusUnauthorized,
		)
		return 0, backendErr
	}

	userIdString, err := token.Claims.GetSubject()
	if err != nil {
		return 0, provideDefaultAccessTokenError(err)
	}
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return 0, provideDefaultAccessTokenError(err)
	}

	return users.UserId(userId), nil
}

func (component *Component) ValidateAndParseRefreshToken(
	refreshToken string,
) (users.UserId, error) {
	options := []jwt.ParserOption{
		jwt.WithIssuedAt(),
		jwt.WithIssuer(component.config.Issuer),
	}

	parser := jwt.NewParser(options...)

	token, err := parser.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		secretKey := []byte(component.config.RefreshSecretKey)
		return secretKey, nil
	})
	if err != nil {
		var backendErr = error2.BackendError{}
		if strings.Contains(err.Error(), "token is expired") {
			backendErr = error2.WrapError(
				err,
				error2.RefreshTokenExpired,
				"Refresh token expired",
				http.StatusUnauthorized,
			)
		} else {
			backendErr = error2.WrapError(
				err,
				error2.InvalidRefreshToken,
				"Failed to parse refresh token",
				http.StatusUnauthorized,
			)
		}
		return 0, backendErr
	}

	expiredAt, err := token.Claims.GetExpirationTime()
	if err != nil {
		return 0, provideDefaultRefreshTokenError(err)
	}

	if expiredAt.Before(time.Now()) {
		backendErr := error2.NewBackendError(
			error2.RefreshTokenExpired,
			"Refresh token expired",
			http.StatusUnauthorized,
		)
		return 0, backendErr
	}

	userIdString, err := token.Claims.GetSubject()
	if err != nil {
		return 0, provideDefaultRefreshTokenError(err)
	}
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return 0, provideDefaultRefreshTokenError(err)
	}

	return users.UserId(userId), nil
}

func provideDefaultAccessTokenError(err error) error {
	return error2.WrapError(
		err,
		error2.InvalidAccessToken,
		"Invalid access token",
		http.StatusUnauthorized,
	)
}

func provideDefaultRefreshTokenError(err error) error {
	return error2.WrapError(
		err,
		error2.InvalidRefreshToken,
		"Invalid refresh token",
		http.StatusUnauthorized,
	)
}
