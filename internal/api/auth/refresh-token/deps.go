package refresh_token

import "context"

type tokenRefresher interface {
	RefreshAccessToken(ctx context.Context, refreshToken string) (string, error)
}
