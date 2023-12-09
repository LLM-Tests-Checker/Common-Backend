package auth

import "fmt"

type TokensProvider interface {
	GenerateTokens(userId int32) (error, *Tokens)

	ValidateAccessToken(accessToken string) error

	GenerateAccessToken(refreshToken string) (error, string)

	ProvideUserId(accessToken string) (error, int32)
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func NewTokensProvider() TokensProvider {
	return jwtTokensProvider{}
}

type jwtTokensProvider struct {
}

func (provider jwtTokensProvider) GenerateTokens(userId int32) (error, *Tokens) {
	return fmt.Errorf("not implemented yet"), nil
}

func (provider jwtTokensProvider) ValidateAccessToken(accessToken string) error {
	return fmt.Errorf("not implemented yet")
}

func (provider jwtTokensProvider) GenerateAccessToken(refreshToken string) (error, string) {
	return fmt.Errorf("not implemented yet"), ""
}

func (provider jwtTokensProvider) ProvideUserId(accessToken string) (error, int32) {
	return fmt.Errorf("not implemented yet"), 0
}
