package auth

type TokenService interface {
	GenerateTokens(userId int32) (error, JWTTokens)

	ValidateAccessToken(accessToken string) error

	GenerateAccessToken(refreshToken string) (error, string)

	ProvideUserId(accessToken string) (error, int32)
}

type JWTTokens struct {
	AccessToken  string
	RefreshToken string
}
