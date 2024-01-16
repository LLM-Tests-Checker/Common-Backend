package jwt

import (
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestComponent_GenerateAccessToken(t *testing.T) {

	type TestCase struct {
		config        Config
		inputUserId   users.UserId
		tokenExpected string
		errExpected   error
	}

	tests := map[string]TestCase{
		"success": {
			config: Config{
				AccessTokenLiveTime:  60 * time.Minute,
				RefreshTokenLiveTime: 24 * 60 * time.Minute,
				AccessSecretKey:      "access_very_secret_key",
				RefreshSecretKey:     "refresh_very_secret_key",
				Issuer:               "Ferum-bot",
			},
			inputUserId:   users.UserId(228),
			tokenExpected: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			component := NewJWTComponent(test.config)

			tokenActual, err := component.GenerateAccessToken(test.inputUserId)

			if test.errExpected != nil {
				assert.ErrorIs(t, err, test.errExpected)
				return
			}
			assert.NoError(t, err)
			assert.Contains(t, tokenActual, test.tokenExpected)
		})
	}
}

func TestComponent_GenerateRefreshToken(t *testing.T) {
	type TestCase struct {
		config        Config
		inputUserId   users.UserId
		tokenExpected string
		errExpected   error
	}

	tests := map[string]TestCase{
		"success": {
			config: Config{
				AccessTokenLiveTime:  60 * time.Minute,
				RefreshTokenLiveTime: 24 * 60 * time.Minute,
				AccessSecretKey:      "access_very_secret_key",
				RefreshSecretKey:     "refresh_very_secret_key",
				Issuer:               "Ferum-bot",
			},
			inputUserId:   users.UserId(228),
			tokenExpected: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			component := NewJWTComponent(test.config)

			tokenActual, err := component.GenerateRefreshToken(test.inputUserId)

			if test.errExpected != nil {
				assert.ErrorIs(t, err, test.errExpected)
				return
			}
			assert.NoError(t, err)
			assert.Contains(t, tokenActual, test.tokenExpected)
		})
	}
}
