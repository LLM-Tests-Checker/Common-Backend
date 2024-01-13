package jwt

import "time"

type Config struct {
	AccessTokenLiveTime  time.Duration
	RefreshTokenLiveTime time.Duration
	AccessSecretKey      string
	RefreshSecretKey     string
}
