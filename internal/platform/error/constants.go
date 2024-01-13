package error

const (
	UnknownError ErrorCode = iota

	InvalidAccessToken
	InvalidRefreshToken
	AccessTokenExpired
	RefreshTokenExpired
)
