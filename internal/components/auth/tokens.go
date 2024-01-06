package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	AccessTokenLiveTimeSeconds  = 60 * 60
	RefreshTokenLiveTimeSeconds = AccessTokenLiveTimeSeconds * 24
)

type TokensProvider interface {
	GenerateTokens(userId int32) (error, *Tokens)

	GenerateAccessToken(refreshToken string) (error, string)
}

type TokenUserIdProvider interface {
	ProvideUserId(accessToken string) (error, int32)
}

type TokensValidator interface {
	ValidateAccessToken(accessToken string) error
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func NewTokensProvider() TokensProvider {
	return jwtTokensProvider{}
}

func NewTokensValidator() TokensValidator {
	return jwtTokensProvider{}
}

func NewTokensUserIdProvider() TokenUserIdProvider {
	return jwtTokensProvider{}
}

type jwtTokensProvider struct {
	key     []byte
	issuer  string
	subject string
}

func (provider jwtTokensProvider) GenerateTokens(userId int32) (error, *Tokens) {
	accessTokenClaims := jwt.MapClaims{
		"iat":  jwt.NewNumericDate(time.Now()),
		"iss":  provider.issuer,
		"sub":  provider.subject,
		"exp":  AccessTokenLiveTimeSeconds,
		"user": userId,
		"type": "access",
	}
	refreshTokenClaims := jwt.MapClaims{
		"iat":  jwt.NewNumericDate(time.Now()),
		"iss":  provider.issuer,
		"sub":  provider.subject,
		"exp":  RefreshTokenLiveTimeSeconds,
		"user": userId,
		"type": "refresh",
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	accessTokenSigned, err := accessToken.SignedString(provider.key)
	if err != nil {
		return err, nil
	}
	refreshTokenSigned, err := refreshToken.SignedString(provider.key)
	if err != nil {
		return err, nil
	}

	tokens := Tokens{
		AccessToken:  accessTokenSigned,
		RefreshToken: refreshTokenSigned,
	}

	return nil, &tokens
}

func (provider jwtTokensProvider) ValidateAccessToken(accessToken string) error {
	token, err := jwt.Parse(accessToken)
	return fmt.Errorf("not implemented yet")
}

func (provider jwtTokensProvider) GenerateAccessToken(refreshToken string) (error, string) {
	return fmt.Errorf("not implemented yet"), ""
}

func (provider jwtTokensProvider) ProvideUserId(accessToken string) (error, int32) {
	return fmt.Errorf("not implemented yet"), 0
}
