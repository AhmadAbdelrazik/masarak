package port

import (
	"errors"
	"net/http"

	"github.com/ahmadabdelrazik/masarak/internal/app"
	"github.com/ahmadabdelrazik/masarak/internal/domain/valueobject"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
	"github.com/ahmadabdelrazik/masarak/pkg/httperr"
	"github.com/ahmadabdelrazik/masarak/pkg/httputils"
)

func (h *HttpServer) selectRole(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r.Context())
	if err != nil {
		httperr.UnauthorizedResponse(w, r)
		return
	}

	var input struct {
		Role string `json:"role"`
	}

	if err := httputils.ReadJSON(w, r, &input); err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	cmd := app.SelectRole{
		User: user,
		ID:   user.ID(),
		Role: input.Role,
	}

	if err := h.app.Commands.SelectRoleHandler(r.Context(), cmd); err != nil {
		switch {
		case errors.Is(err, authuser.ErrUserNotFound):
			httperr.NotFoundResponse(w, r)
		case errors.Is(err, app.ErrEditConflict):
			httperr.EditConflictResponse(w, r)
		case errors.Is(err, app.ErrUnauthorized):
			httperr.UnauthorizedResponse(w, r)
		case errors.Is(err, valueobject.ErrInvalidRole):
			httperr.BadRequestResponse(w, r, errors.New("invalid role"))
		default:
			httperr.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err := httputils.WriteJSON(w, http.StatusOK, httputils.Envelope{"new_role": input.Role}, nil); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}
}
