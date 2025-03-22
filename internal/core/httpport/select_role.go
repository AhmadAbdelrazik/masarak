package httpport

import (
	"errors"
	"net/http"

	"github.com/ahmadabdelrazik/masarak/internal/core/app"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/valueobject"
	"github.com/ahmadabdelrazik/masarak/internal/core/httpport/auth"
	"github.com/ahmadabdelrazik/masarak/pkg/httperr"
	"github.com/rs/zerolog/log"
)

func (h *HttpServer) registerUser(w http.ResponseWriter, r *http.Request) {
	user, err := auth.UserFromCtx(r.Context())
	if err != nil {
		log.Error().Err(err).Msg("")
		httperr.UnauthorizedResponse(w, r)
		return
	}

	var input struct {
		Role string `json:"role"`
	}

	if err := readJSON(w, r, &input); err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	cmd := app.RegisterUserType{
		User: user,
		Role: input.Role,
	}

	err = h.app.Commands.RegisterOwner.Handle(r.Context(), cmd)
	if err != nil {
		switch {
		case errors.Is(err, app.ErrUserAlreadyRegistered):
			httperr.ErrorResponse(w, r, http.StatusBadRequest, "user already has role: "+user.Role.String())
		case errors.Is(err, valueobject.ErrInvalidRole):
			httperr.ErrorResponse(w, r, http.StatusUnprocessableEntity, "invalid role")
		default:
			httperr.ServerErrorResponse(w, r, err)
		}
		return
	}

	writeJSON(w, http.StatusCreated, envelope{"message": "registered as an owner"}, nil)
}
