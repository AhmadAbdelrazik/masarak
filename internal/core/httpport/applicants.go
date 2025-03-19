package httpport

import (
	"errors"
	"net/http"

	"github.com/ahmadabdelrazik/masarak/internal/core/app"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/company"
	"github.com/ahmadabdelrazik/masarak/pkg/httperr"
)

func (h *HttpServer) postOwner(w http.ResponseWriter, r *http.Request) {
	user, err := userFromCtx(r.Context())
	if err != nil {
		httperr.UnauthorizedResponse(w, r)
		return
	}
	if !user.Role.Is("user") {
		httperr.ErrorResponse(w, r, http.StatusBadRequest, "user already has role: "+user.Role.String())
		return
	}
	cmd := app.CreateOwner{Name: user.Name, Email: user.Email}

	err = h.app.Commands.CreateOwner.Handle(r.Context(), cmd)
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	writeJSON(w, http.StatusCreated, envelope{"message": "registered as an owner"}, nil)
}

func (h *HttpServer) postCompany(w http.ResponseWriter, r *http.Request) {
	user, err := userFromCtx(r.Context())
	if err != nil {
		httperr.UnauthorizedResponse(w, r)
		return
	}

	var input struct {
		CompanyName    string `json:"company_name"`
		CompanyDetails string `json:"company_details"`
		LineOfBusiness string `json:"line_of_business"`
	}

	if err := readJSON(w, r, &input); err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	cmd := app.CreateCompany{
		OwnerEmail:            user.Email,
		CompanyName:           input.CompanyName,
		CompanyDetails:        input.CompanyDetails,
		CompanyLineOfBusiness: input.LineOfBusiness,
	}

	err = h.app.Commands.CreateCompany.Handle(r.Context(), cmd)
	if err != nil {
		switch {
		case errors.Is(err, company.ErrAlreadyExists):
			httperr.ErrorResponse(
				w,
				r,
				http.StatusUnprocessableEntity,
				"company with this name already exists",
			)
		default:
			httperr.ServerErrorResponse(w, r, err)
		}
		return
	}

}

func (h *HttpServer) postJob(w http.ResponseWriter, r *http.Request) {
	user, err := userFromCtx(r.Context())
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

	err = h.app.Commands.CreateJob.Handle(r.Context(), cmd)
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
