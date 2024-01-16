package http

import (
	"context"
	"github.com/LLM-Tests-Checker/Common-Backend/internal/services/users"
)

type tokenParser interface {
	ParseUserId(ctx context.Context, accessToken string) (users.UserId, error)
}
