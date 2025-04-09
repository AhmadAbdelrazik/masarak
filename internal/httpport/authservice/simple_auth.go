package authservice

import (
	"errors"
	"net/http"

	"github.com/ahmadabdelrazik/masarak/internal/app"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
	"github.com/ahmadabdelrazik/masarak/pkg/httperr"
	"github.com/ahmadabdelrazik/masarak/pkg/httputils"
)

func (h *AuthService) Signup(w http.ResponseWriter, r *http.Request) {
	// provide sign up credentials
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := httputils.ReadJSON(w, r, &input); err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	cmd := app.CreateUser{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     "user",
	}

	if err := h.app.Commands.CreateUserHandler(r.Context(), cmd); err != nil {
		switch {
		case errors.Is(err, authuser.ErrUserAlreadyExists):
			httperr.ErrorResponse(w, r, http.StatusConflict, "user already exists")
		case errors.Is(err, authuser.ErrInvalidProperty):
			httperr.BadRequestResponse(w, r, err)
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

	var output struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	output.Name = input.Name
	output.Email = input.Email

	http.SetCookie(w, cookie)
	if err := httputils.WriteJSON(w, http.StatusCreated, httputils.Envelope{"message": "registered successfully", "user": output}, nil); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}
}

func (h *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := httputils.ReadJSON(w, r, &input); err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	user, err := h.app.Queries.UserLogin(r.Context(), input.Email, input.Password)
	if err != nil {
		switch {
		case errors.Is(err, authuser.ErrUserNotFound), errors.Is(err, app.ErrInvalidPassword):
			httperr.AuthenticationErrorResponse(w, r)
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
	if err := httputils.WriteJSON(
		w,
		http.StatusOK,
		httputils.Envelope{"message": "logged in successfully", "user": user},
		nil,
	); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}

}
