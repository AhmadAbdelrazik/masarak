package authservice

import (
	"context"
	"errors"
	"net/http"

	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

func getTokenCookie(r *http.Request, email string, tokenRepo authuser.TokenRepository) (*http.Cookie, error) {

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

func UserFromCtx(ctx context.Context) (*authuser.User, error) {
	user, ok := ctx.Value(UserContextKey).(*authuser.User)
	if ok {
		return user, nil
	}

	return nil, NoUserInContextError
}
