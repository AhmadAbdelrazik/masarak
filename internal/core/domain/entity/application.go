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
	status, err := valueobject.NewApplicationStatus("pending")
	if err != nil {
		panic(err)
	}

	return &Application{
		ID:                uuid.New(),
		status:            status,
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

func (a *Application) IsPending() bool {
	return a.status.IsPending()
}

func (a *Application) IsAccepted() bool {
	return a.status.IsAccepted()
}

func (a *Application) IsRejected() bool {
	return a.status.IsRejected()
}

func (a *Application) Accept() {
	a.setStatus("accepted")
}

func (a *Application) Reject() {
	a.setStatus("rejected")
}

func (a *Application) setStatus(statusString string) {
	status, err := valueobject.NewApplicationStatus(statusString)
	if err != nil {
		panic(err)
	}
	a.status = status
}

func (a *Application) Update(
	name, email, title string,
	yearsOfExperience int,
	hourlyRate *money.Money,
	freelancerProfile, resumeURL string,
) error {
	if !a.IsPending() {
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
