package auth

import (
	"errors"
	"net/http"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/valueobject"
	"github.com/ahmadabdelrazik/masarak/pkg/httperr"
)

func (h *AuthService) Signup(w http.ResponseWriter, r *http.Request) {
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

	user, err := authuser.New(input.Name, input.Email, input.Password, role)
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	if err := h.userRepo.Create(r.Context(), user); err != nil {
		switch {
		case errors.Is(err, authuser.ErrUserAlreadyExists):
			httperr.ErrorResponse(w, r, http.StatusForbidden, "user already exists")
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

	output.Name = user.Name()
	output.Email = user.Email()

	http.SetCookie(w, cookie)
	if err := writeJSON(w, http.StatusCreated, envelope{"message": "registered successfully", "user": output}, nil); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}
}

func (h *AuthService) Login(w http.ResponseWriter, r *http.Request) {
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
		case errors.Is(err, authuser.ErrUserNotFound):
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

	var output struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	output.Name = user.Name()
	output.Email = user.Email()

	http.SetCookie(w, cookie)
	if err := writeJSON(w, http.StatusOK, envelope{"message": "logged in successfully", "user": output}, nil); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}

}
