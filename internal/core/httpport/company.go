package httpport

import (
	"errors"
	"net/http"

	"github.com/ahmadabdelrazik/masarak/internal/core/app"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/company"
	"github.com/ahmadabdelrazik/masarak/internal/core/httpport/auth"
	"github.com/ahmadabdelrazik/masarak/pkg/httperr"
)

func (h *HttpServer) postCompany(w http.ResponseWriter, r *http.Request) {
	user, err := auth.UserFromCtx(r.Context())
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

	err = h.app.Commands.CreateCompanyHandler(r.Context(), cmd)
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
