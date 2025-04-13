package port

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/ahmadabdelrazik/masarak/internal/app"
	"github.com/ahmadabdelrazik/masarak/internal/domain/freelancerprofile"
	"github.com/ahmadabdelrazik/masarak/pkg/filters"
	"github.com/ahmadabdelrazik/masarak/pkg/httperr"
	"github.com/ahmadabdelrazik/masarak/pkg/httputils"
)

func (h *HttpServer) createFreelancerProfileHandler(w http.ResponseWriter, r *http.Request) {
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

func (h *HttpServer) getFreelancerProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form.Get("email")

	user, err := getUser(r.Context())
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	cmd := app.GetFreelancerProfile{
		User:  user,
		Email: email,
	}

	profile, err := h.app.Queries.GetFreelancerProfileHandler(r.Context(), cmd)
	if err != nil {
		switch {
		case errors.Is(err, freelancerprofile.ErrProfileNotFound):
			httperr.NotFoundResponse(w, r)
		default:
			httperr.ServerErrorResponse(w, r, err)
		}
		return
	}

	profile.ResumeURL = fmt.Sprintf("http://%v/%v", h.cfg.HostURL, profile.ResumeURL)
	profile.PictureURL = fmt.Sprintf("http://%v/%v", h.cfg.HostURL, profile.PictureURL)

	httputils.WriteJSON(w, http.StatusOK, httputils.Envelope{"profile": profile}, nil)
}

func (h *HttpServer) updateFreelancerProfile(w http.ResponseWriter, r *http.Request) {
	var cmd app.UpdateFreelancerProfile

	user, err := getUser(r.Context())
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	cmd.User = user

	if err := r.ParseMultipartForm((1 << 20) * 10); err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	name := r.FormValue("name")
	if name != "" {
		cmd.Name = &name
	}

	title := r.FormValue("title")
	if title != "" {
		cmd.Title = &title
	}

	skills := r.Form["skills"]
	if len(skills) != 0 {
		cmd.Skills = skills
	}

	yearsOfExperienceStr := r.FormValue("years_of_experience")
	if yearsOfExperienceStr != "" {
		yoe, err := strconv.ParseInt(yearsOfExperienceStr, 10, 64)
		if err != nil {
			httperr.BadRequestResponse(w, r, err)
			return
		}

		yearsOfExperience := int(yoe)

		cmd.YearsOfExperience = &yearsOfExperience
	}

	hourlyRateAmountStr := r.FormValue("hourly_rate_amount")
	if hourlyRateAmountStr != "" {
		hra, err := strconv.ParseInt(hourlyRateAmountStr, 10, 64)
		if err != nil {
			httperr.BadRequestResponse(w, r, err)
			return
		}

		hourlyRateAmount := int(hra)

		cmd.HourlyRateAmount = &hourlyRateAmount
	}

	hourlyRateCurrency := r.FormValue("hourly_rate_currency")
	if hourlyRateCurrency != "" {
		cmd.HourlyRateCurrency = &hourlyRateCurrency
	}

	pictureURL, err := httputils.SaveFile(r, "picture", filepath.Join(".", "uploads", "images"), httputils.ImagesMime...)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrMissingFile):
		default:
			httperr.BadRequestResponse(w, r, err)
			return
		}
	} else {
		cmd.PictureURL = &pictureURL
	}

	resumeURL, err := httputils.SaveFile(r, "resume", filepath.Join(".", "uploads", "resumes"), httputils.PDFmime)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrMissingFile):
		default:
			os.Remove(pictureURL)
			httperr.BadRequestResponse(w, r, err)
			return
		}
	} else {
		cmd.ResumeURL = &resumeURL
	}

	oldProfile, err := h.app.Queries.GetFreelancerProfileHandler(r.Context(), app.GetFreelancerProfile{User: user, Email: user.Email()})
	if err != nil {
		switch {
		case errors.Is(err, freelancerprofile.ErrProfileNotFound):
			httperr.NotFoundResponse(w, r)
		default:
			httperr.ServerErrorResponse(w, r, err)
		}
		return
	}

	err = h.app.Commands.UpdateFreelancerProfileHandler(r.Context(), cmd)
	if err != nil {
		switch {
		case errors.Is(err, app.ErrUnauthorized):
			httperr.UnauthorizedResponse(w, r)
		case errors.Is(err, freelancerprofile.ErrDuplicateProfile),
			errors.Is(err, freelancerprofile.ErrInvalidYearsOfExperience),
			errors.Is(err, freelancerprofile.ErrSkillLimitReached),
			errors.Is(err, freelancerprofile.ErrInvalidHourlyRate):
			httperr.BadRequestResponse(w, r, err)
		case errors.Is(err, app.ErrEditConflict):
			httperr.EditConflictResponse(w, r)
		default:
			httperr.ServerErrorResponse(w, r, err)
		}
		os.Remove(resumeURL)
		os.Remove(pictureURL)
		return
	}

	if cmd.ResumeURL != nil {
		os.Remove(oldProfile.ResumeURL)
	}
	if cmd.PictureURL != nil {
		os.Remove(oldProfile.PictureURL)
	}

	err = httputils.WriteJSON(w, http.StatusCreated, httputils.Envelope{"message": "updated"}, nil)
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}
}

func (h *HttpServer) searchFreelancerProfiles(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	title := r.FormValue("title")
	skills := r.Form["skills"]
	yoeStr := r.FormValue("years_of_experience")
	yearsOfExperience := -1

	if yoeStr != "" {
		yoe, err := strconv.ParseInt(yoeStr, 10, 64)
		if err != nil || yoe < 0 {
			httperr.BadRequestResponse(w, r, err)
			return
		}
	}

	hraStr := r.FormValue("hourly_rate_amount")
	hourlyRateAmount := -1
	if hraStr != "" {
		hra, err := strconv.ParseInt(hraStr, 10, 64)
		if err != nil || hra < 0 {
			httperr.BadRequestResponse(w, r, err)
			return
		}
	}

	hourlyRateCurrency := r.FormValue("hourly_rate_currency")

	sort := r.FormValue("sort")

	pageSize := 20
	pageSizeStr := r.FormValue("page_size")
	if pageSizeStr != "" {
		ps, err := strconv.ParseInt(pageSizeStr, 10, 64)
		if err != nil || ps <= 0 {
			httperr.BadRequestResponse(w, r, err)
			return
		}

		pageSize = int(ps)
	}

	pageNumber := 1
	pageNumberStr := r.FormValue("page_number")
	if pageNumberStr != "" {
		pn, err := strconv.ParseInt(pageNumberStr, 10, 64)
		if err != nil || pn <= 0 {
			httperr.BadRequestResponse(w, r, err)
			return
		}

		pageNumber = int(pn)
	}

	filter, err := filters.NewSQLFilter(
		filters.WithPage(pageNumber),
		filters.WithPageSize(pageSize),
		filters.WithSort(sort, "name", []string{
			"name",
			"title",
			"years_of_experience",
			"hourly_rate_amount",
		}),
	)
	if err != nil {
		httperr.BadRequestResponse(w, r, err)
		return
	}

	cmd := app.SearchFreelancerProfiles{
		Name:               name,
		Title:              title,
		Skills:             skills,
		YearsOfExperience:  yearsOfExperience,
		HourlyRateAmount:   hourlyRateAmount,
		HourlyRateCurrency: hourlyRateCurrency,
		Filters:            filter,
	}
	profiles, meta, err := h.app.Queries.SearchFreelancerProfilesHandler(r.Context(), cmd)
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	httputils.WriteJSON(w, http.StatusOK, httputils.Envelope{"profiles": profiles, "metadata": meta}, nil)
}
