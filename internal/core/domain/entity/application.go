package entity

import (
	"errors"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/valueobject"
	"github.com/google/uuid"
)

var (
	ErrUnableToUpdate = errors.New("updating window is closed")
)

type Application struct {
	ID                uuid.UUID
	status            *valueobject.ApplicationStatus
	Name              string
	Email             string
	Title             string
	YearsOfExperience int
	HourlyRate        *money.Money
	FreelancerProfile string
	ResumeURL         string
	CreatedAt         time.Time
}

func NewApplication(
	name, email, title string,
	yearsOfExperience int,
	hourlyRate *money.Money,
	freelancerProfile, resumeURL string,
) *Application {
	return &Application{
		ID:                uuid.New(),
		status:            valueobject.NewApplicationStatus("pending"),
		Name:              name,
		Email:             email,
		Title:             title,
		YearsOfExperience: yearsOfExperience,
		HourlyRate:        hourlyRate,
		FreelancerProfile: freelancerProfile,
		ResumeURL:         resumeURL,
		CreatedAt:         time.Now(),
	}
}

func (a *Application) Accept() {
	a.status = valueobject.NewApplicationStatus("accepted")
}

func (a *Application) Reject() {
	a.status = valueobject.NewApplicationStatus("rejected")
}

func (a *Application) Status() string {
	return a.status.Status()
}

func (a *Application) Update(
	name, email, title string,
	yearsOfExperience int,
	hourlyRate *money.Money,
	freelancerProfile, resumeURL string,
) error {
	if !a.status.IsPending() {
		return ErrUnableToUpdate
	}

	a.Name = name
	a.Email = email
	a.Title = title
	a.YearsOfExperience = yearsOfExperience
	a.HourlyRate = hourlyRate
	a.FreelancerProfile = freelancerProfile
	a.ResumeURL = resumeURL

	return nil
}
