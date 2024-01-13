package error

const (
	UnknownError ErrorCode = iota

	InputValidationError

	InvalidAccessToken
	InvalidRefreshToken
	AccessTokenExpired
	RefreshTokenExpired

	InvalidLLMSlug
)
