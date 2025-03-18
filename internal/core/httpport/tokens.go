package httpport

import (
	"context"
	"errors"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
)

type Token string

type TokenRepository interface {
	GenerateToken(ctx context.Context, email string) (Token, error)
	GetFromToken(ctx context.Context, token Token) (*entity.AuthUser, error)
}
type ctxKey int

const (
	UserContextKey ctxKey = iota
)

var (
	NoUserInContextError = errors.New("no user in context error")
)

func userFromCtx(ctx context.Context) (entity.AuthUser, error) {
	u, ok := ctx.Value(UserContextKey).(entity.AuthUser)
	if ok {
		return u, nil
	}

	return entity.AuthUser{}, NoUserInContextError
}
