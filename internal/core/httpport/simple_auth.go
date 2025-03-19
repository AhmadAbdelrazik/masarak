package httpport

import (
	"context"
	"errors"
	"net/http"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/valueobject"
	"github.com/ahmadabdelrazik/linkedout/pkg/httperr"
	"github.com/google/uuid"
)

func (h *HttpServer) Signup(w http.ResponseWriter, r *http.Request) {
	// provide sign up credentials
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := readJSON(w, r, &input); err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	role, err := valueobject.NewRole("user")
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	user, err := entity.NewAuthUser(uuid.NewString(), input.Name, input.Email, input.Password, role)
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	if err := h.userRepo.Add(r.Context(), user); err != nil {
		switch {
		case errors.Is(err, entity.ErrUserAlreadyExists):
			httperr.ErrorResponse(w, r, http.StatusUnprocessableEntity, "user already exists")
		default:
			httperr.ServerErrorResponse(w, r, err)
		}
		return
	}

	cookie, err := getTokenCookie(r, input.Email, h.tokenRepo)
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	http.SetCookie(w, cookie)
	if err := writeJSON(w, http.StatusOK, envelope{"message": "logged in successfully"}, nil); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}
}

func (h *HttpServer) login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := readJSON(w, r, &input); err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	user, err := h.userRepo.GetByEmail(r.Context(), input.Email)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrUserNotFound):
			httperr.AuthenticationErrorResponse(w, r)
		default:
			httperr.ServerErrorResponse(w, r, err)
		}
		return
	}

	if match, err := user.Password.Matches(input.Password); err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	} else if !match {
		httperr.AuthenticationErrorResponse(w, r)
		return
	}

	cookie, err := getTokenCookie(r, input.Email, h.tokenRepo)
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	http.SetCookie(w, cookie)
	if err := writeJSON(w, http.StatusOK, envelope{"message": "logged in successfully"}, nil); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}

}

func (h *HttpServer) IsAuthenticated(next http.HandlerFunc) http.Handler {
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

		user, err := h.tokenRepo.GetFromToken(r.Context(), Token(cookie.Value))

		ctx := r.Context()

		ctx = context.WithValue(ctx, UserContextKey, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
