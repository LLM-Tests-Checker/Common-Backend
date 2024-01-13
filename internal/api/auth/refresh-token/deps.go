package refresh_token

type tokenRefresher interface {
	RefreshAccessToken(refreshToken string) (string, error)
}
