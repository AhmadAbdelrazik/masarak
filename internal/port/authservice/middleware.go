package authservice

import (
	"context"
	"errors"
	"net/http"

	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
	"github.com/ahmadabdelrazik/masarak/pkg/httperr"
)

// IsAuthenticated - Check if user is authenticated by validating the
// session_id token in http request
func (a *AuthService) IsAuthenticated(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("session_id")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				httperr.AuthenticationErrorResponse(w, r)
			default:
				httperr.BadRequestResponse(w, r, err)
			}
			return
		}

		user, err := a.userRepo.GetByToken(r.Context(), authuser.Token(cookie.Value))
		if err != nil {
			switch {
			case errors.Is(err, authuser.ErrUserNotFound):
				httperr.AuthenticationErrorResponse(w, r)
			default:
				httperr.ServerErrorResponse(w, r, err)
			}
			return
		}

		ctx := r.Context()

		ctx = context.WithValue(ctx, UserContextKey, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// HasPermission - Check if authenticated user has the required permission to
// access the resources. This acts as RBAC based permissions, you might use
// another authorization layer for differentiating between users with the same
// role
func (a *AuthService) HasPermission(permission string, next http.HandlerFunc) http.Handler {
	return a.IsAuthenticated(func(w http.ResponseWriter, r *http.Request) {
		user, err := UserFromCtx(r.Context())
		if err != nil {
			httperr.ServerErrorResponse(w, r, err)
			return
		}

		if !user.HasPermission(permission) {
			httperr.UnauthorizedResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
