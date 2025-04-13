package authservice

import (
	"context"
	"errors"
	"net/http"

	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

// getTokenCookie - Generates a new cookie that contains the token session. the
// newly created token will be stored in the database.
func getTokenCookie(r *http.Request, id int, tokenRepo authuser.TokenRepository) (*http.Cookie, error) {
	userToken, err := tokenRepo.GenerateToken(r.Context(), id)
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
