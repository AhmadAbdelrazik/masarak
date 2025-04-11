package port

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ahmadabdelrazik/masarak/internal/app"
	"github.com/ahmadabdelrazik/masarak/internal/domain/freelancerprofile"
	"github.com/ahmadabdelrazik/masarak/pkg/httperr"
	"github.com/ahmadabdelrazik/masarak/pkg/httputils"
)

func (h *HttpServer) CreateFreelancerProfileHandler(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r.Context())
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	if err := r.ParseMultipartForm((1 << 20) * 10); err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	email := user.Email()
	name := r.FormValue("name")
	title := r.FormValue("title")
	skills := r.Form["skills"]

	yearsOfExperienceStr := r.FormValue("years_of_experience")
	yearsOfExperience, err := strconv.ParseInt(yearsOfExperienceStr, 10, 64)
	if err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	hourlyRateAmountStr := r.FormValue("hourly_rate_amount")
	hourlyRateAmount, err := strconv.ParseInt(hourlyRateAmountStr, 10, 64)
	if err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	hourlyRateCurrency := r.FormValue("hourly_rate_currency")

	pictureURL, err := httputils.SaveFile(r, "picture", filepath.Join(".", "uploads", "images"), httputils.ImagesMime...)
	if err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	resumeURL, err := httputils.SaveFile(r, "resume", filepath.Join(".", "uploads", "resumes"), httputils.PDFmime)
	if err != nil {
		os.Remove(pictureURL)
		httperr.BadRequestResponse(w, r, err)
		return
	}

	cmd := app.CreateFreelancerProfile{
		User:               user,
		Email:              email,
		Name:               name,
		Title:              title,
		PictureURL:         pictureURL,
		Skills:             skills,
		YearsOfExperience:  int(yearsOfExperience),
		HourlyRateAmount:   int(hourlyRateAmount),
		HourlyRateCurrency: hourlyRateCurrency,
		ResumeURL:          resumeURL,
	}

	err = h.app.Commands.CreateFreelancerProfileHandler(r.Context(), cmd)
	if err != nil {
		switch {
		case errors.Is(err, app.ErrUnauthorized):
			httperr.UnauthorizedResponse(w, r)
		case errors.Is(err, freelancerprofile.ErrDuplicateProfile),
			errors.Is(err, freelancerprofile.ErrInvalidYearsOfExperience),
			errors.Is(err, freelancerprofile.ErrSkillLimitReached),
			errors.Is(err, freelancerprofile.ErrInvalidHourlyRate):
			httperr.BadRequestResponse(w, r, err)
		default:
			httperr.ServerErrorResponse(w, r, err)
		}
		os.Remove(resumeURL)
		os.Remove(pictureURL)
		return
	}

	err = httputils.WriteJSON(w, http.StatusCreated, httputils.Envelope{"message": "created"}, nil)
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}
}
