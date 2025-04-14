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
		Username string `json:"username"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := httputils.ReadJSON(w, r, &input); err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	cmd := app.CreateUser{
		Username: input.Username,
		Email:    input.Email,
		Name:     input.Name,
		Password: input.Password,
		Role:     "user",
	}

	user, err := h.app.Commands.CreateUserHandler(r.Context(), cmd)
	if err != nil {
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

	cookie, err := getTokenCookie(r, user.ID, h.tokenRepo)
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	http.SetCookie(w, cookie)
	if err := httputils.WriteJSON(w, http.StatusCreated, httputils.Envelope{"message": "registered successfully", "user": user}, nil); err != nil {
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

	cmd := app.UserLogin{
		Email:    input.Email,
		Password: input.Password,
	}

	user, err := h.app.Queries.UserLogin(r.Context(), cmd)
	if err != nil {
		switch {
		case errors.Is(err, authuser.ErrUserNotFound), errors.Is(err, app.ErrInvalidPassword):
			httperr.AuthenticationErrorResponse(w, r)
		default:
			httperr.ServerErrorResponse(w, r, err)
		}
		return
	}

	cookie, err := getTokenCookie(r, user.ID, h.tokenRepo)
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

func (h *AuthService) Logout(w http.ResponseWriter, r *http.Request) {
	user, err := UserFromCtx(r.Context())
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	if err := h.tokenRepo.DeleteTokensByID(r.Context(), user.ID()); err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	}

	http.SetCookie(w, cookie)

	if err := httputils.WriteJSON(w, http.StatusOK, httputils.Envelope{"message": "logged out successfully"}, nil); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}
}
