package error

const (
	UnknownError ErrorCode = iota

	InputValidationError
	UserWithLoginAlreadyExists
	UserNotFound
	UserInvalidPassword

	NotOwnerError

	TestNotFound

	InvalidAccessToken
	InvalidRefreshToken
	AccessTokenExpired
	RefreshTokenExpired

	InvalidLLMSlug
)
