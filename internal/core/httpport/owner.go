package httpport

import (
	"errors"
	"net/http"

	"github.com/ahmadabdelrazik/masarak/internal/core/app"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/company"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/owner"
	"github.com/ahmadabdelrazik/masarak/internal/core/httpport/auth"
	"github.com/ahmadabdelrazik/masarak/pkg/httperr"
)

func (h *HttpServer) getOwner(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
	}

	if err := readJSON(w, r, &input); err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	cmd := app.GetOwner{Email: input.Email}

	ownerDTO, err := h.app.Queries.GetOwnerHandler(r.Context(), cmd)
	if err != nil {
		switch {
		case errors.Is(err, owner.ErrOwnerNotFound):
			httperr.NotFoundResponse(w, r)
		default:
			httperr.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err := writeJSON(w, http.StatusOK, envelope{"owner": ownerDTO}, nil); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}
}

func (h *HttpServer) getOwners(w http.ResponseWriter, r *http.Request) {
	owners, err := h.app.Queries.GetOwnersHandler(r.Context())
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, envelope{"owners": owners}, nil); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}
}

func (h *HttpServer) postJob(w http.ResponseWriter, r *http.Request) {
	user, err := auth.UserFromCtx(r.Context())
	if err != nil {
		httperr.UnauthorizedResponse(w, r)
		return
	}

	var input struct {
		companyName    string
		jobTitle       string
		jobDescription string
	}

	if err := readJSON(w, r, &input); err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	cmd := app.CreateJob{
		Email:          user.Email,
		CompanyName:    input.companyName,
		JobTitle:       input.jobTitle,
		JobDescription: input.jobDescription,
	}

	err = h.app.Commands.CreateJobHandler(r.Context(), cmd)
	if err != nil {
		switch {
		case errors.Is(err, company.ErrCompanyNotFound):
			httperr.ErrorResponse(w, r, http.StatusNotFound,
				"company with the given name doesn't exist")
		case errors.Is(err, app.ErrInvalidOwner):
			httperr.UnauthorizedResponse(w, r)
		default:
			httperr.ServerErrorResponse(w, r, err)
		}
		return
	}

	if err := writeJSON(w, http.StatusCreated, envelope{"message": "job posted successfully"}, nil); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}
}
