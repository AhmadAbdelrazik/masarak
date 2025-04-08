package entity

import (
	"errors"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/ahmadabdelrazik/masarak/internal/domain/valueobject"
)

var (
	ErrUnableToUpdate = errors.New("updating window is closed")
)

type Application struct {
	ID                int
	status            *valueobject.ApplicationStatus
	Name              string
	Email             string
	Title             string
	YearsOfExperience int
	HourlyRate        *money.Money
	FreelancerProfile string
	ResumeURL         string
	CreatedAt         time.Time
	UpdatedAt         time.Time
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
		status:            status,
		Name:              name,
		Email:             email,
		Title:             title,
		YearsOfExperience: yearsOfExperience,
		HourlyRate:        hourlyRate,
		FreelancerProfile: freelancerProfile,
		ResumeURL:         resumeURL,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

func ReconstituteApplication(
	id int,
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
		ID:                id,
		status:            status,
		Name:              name,
		Email:             email,
		Title:             title,
		YearsOfExperience: yearsOfExperience,
		HourlyRate:        hourlyRate,
		FreelancerProfile: freelancerProfile,
		ResumeURL:         resumeURL,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

func (a *Application) Update(
	name, title string,
	yearsOfExperience int,
	hourlyRate *money.Money,
	freelancerProfile, resumeURL string,
) error {
	a.Name = name
	a.Title = title
	a.YearsOfExperience = yearsOfExperience
	a.HourlyRate = hourlyRate
	a.FreelancerProfile = freelancerProfile
	a.ResumeURL = resumeURL
	a.UpdatedAt = time.Now()

	return nil
}

// Status - return the job application status ("accepted", "rejected", "pending")
func (a *Application) Status() string {
	return a.status.Status()
}

func (a *Application) SetStatusToAccepted() error {
	return a.setStatus("accepted")
}

func (a *Application) SetStatusToRejected() error {
	return a.setStatus("rejected")
}

func (a *Application) SetStatusToPending() error {
	return a.setStatus("pending")
}

func (a *Application) setStatus(statusString string) error {
	status, err := valueobject.NewApplicationStatus(statusString)
	if err != nil {
		panic(err)
	}
	a.status = status

	return nil
}
