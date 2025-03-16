package httpport

import (
	"errors"
	"net/http"

	"github.com/ahmadabdelrazik/linkedout/internal/app/command"
	"github.com/ahmadabdelrazik/linkedout/internal/app/query"
	"github.com/ahmadabdelrazik/linkedout/internal/common/auth/tokens"
	"github.com/ahmadabdelrazik/linkedout/internal/common/server/httperr"
	"github.com/ahmadabdelrazik/linkedout/internal/domain/applicant"
)

func (h *HttpServer) SelectPersonRole(w http.ResponseWriter, r *http.Request) {
	user, err := tokens.UserFromCtx(r.Context())
	if err != nil {
		httperr.AuthenticationErrorResponse(w, r)
		return
	}

	var input struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Role      string `json:"role"`
	}

	cmd := command.SelectPersonRole{
		User:      user,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Role:      input.Role,
	}

	err = h.app.Commands.SelectPersonRole.Handle(r.Context(), cmd)
	if err != nil {
		switch {
		case errors.Is(err, applicant.ErrApplicantExists):
			httperr.BadRequestResponse(w, r, err)
		default:
			httperr.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err := writeJSON(w, http.StatusCreated, envelope{"message": "success"}, nil); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}
}

func (h *HttpServer) GetApplicant(w http.ResponseWriter, r *http.Request) {
	user, err := tokens.UserFromCtx(r.Context())
	if err != nil {
		httperr.AuthenticationErrorResponse(w, r)
		return
	}

	cmd := query.GetApplicant{
		Email: user.Email,
	}

	a, err := h.app.Queries.GetApplicant.Handle(r.Context(), cmd)
	if err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	var input struct {
		Email    string
		FullName string
	}

	input.Email = a.GetEmail()
	input.FullName = a.GetFullName()

	writeJSON(w, http.StatusOK, envelope{"applicant": a}, nil)
}

func (h *HttpServer) ApplicantNumbers(w http.ResponseWriter, r *http.Request) {
}
