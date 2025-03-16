package httpport

import (
	"context"
	"net/http"

	"github.com/ahmadabdelrazik/linkedout/internal/common/auth/tokens"
	users "github.com/ahmadabdelrazik/linkedout/internal/domain/user"
)

type TokenRepository interface {
	GenerateToken(ctx context.Context, email string) (string, error)
	GetFromToken(ctx context.Context, token string) (*users.User, error)
}

type OAuthService interface {
	AuthMiddleware(next http.HandlerFunc) http.Handler
	GoogleLogin(w http.ResponseWriter, r *http.Request)
	GoogleCallback(w http.ResponseWriter, r *http.Request)
}

func UserFromCtx(ctx context.Context) (users.User, error) {
	u, ok := ctx.Value(tokens.UserContextKey).(users.User)
	if ok {
		return u, nil
	}

	return users.User{}, tokens.NoUserInContextError
}
