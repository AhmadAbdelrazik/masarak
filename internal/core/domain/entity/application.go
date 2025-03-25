package entity

import (
	"github.com/Rhymond/go-money"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/valueobject"
	"github.com/google/uuid"
)

type Application struct {
	ID                uuid.UUID
	Status            *valueobject.ApplicationStatus
	Name              string
	Email             string
	Title             string
	YearsOfExperience int
	HourlyRate        *money.Money
	FreelancerProfile string
	ResumeURL         string
}

func NewApplication(
	name, email, title string,
	yearsOfExperience int,
	hourlyRate *money.Money,
	freelancerProfile, resumeURL string,
) *Application {
	return &Application{
		ID:                uuid.New(),
		Status:            valueobject.NewApplicationStatus("pending"),
		Name:              name,
		Email:             email,
		Title:             title,
		YearsOfExperience: yearsOfExperience,
		HourlyRate:        hourlyRate,
		FreelancerProfile: freelancerProfile,
		ResumeURL:         resumeURL,
	}
}

func (a *Application) Accept() {
	a.Status = valueobject.NewApplicationStatus("accepted")
}

func (a *Application) Reject() {
	a.Status = valueobject.NewApplicationStatus("rejected")
}
