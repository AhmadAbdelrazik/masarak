package httpport

import (
	"context"
	"errors"
	"net/http"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
)

type Token string

type TokenRepository interface {
	GenerateToken(ctx context.Context, email string) (Token, error)
	GetFromToken(ctx context.Context, token Token) (*entity.AuthUser, error)
}

func getTokenCookie(r *http.Request, email string, tokenRepo TokenRepository) (*http.Cookie, error) {

	userToken, err := tokenRepo.GenerateToken(r.Context(), email)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    string(userToken),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	return cookie, nil
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
