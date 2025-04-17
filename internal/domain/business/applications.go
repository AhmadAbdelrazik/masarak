package business

import (
	"fmt"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/ahmadabdelrazik/masarak/internal/domain"
	"github.com/ahmadabdelrazik/masarak/internal/domain/valueobject"
	"github.com/ahmadabdelrazik/masarak/pkg/validator"
)

type Application struct {
	businessID        int
	businessName      string
	jobID             int
	jobTitle          string
	id                int
	status            *valueobject.ApplicationStatus
	name              string
	email             string
	title             string
	yearsOfExperience int
	hourlyRate        *money.Money
	FreelancerProfile string
	ResumeURL         string
	createdAt         time.Time
	updatedAt         time.Time
}

var (
	ErrInvalidApplicationProperty = fmt.Errorf("%w: application", domain.ErrInvalidProperty)
	ErrInvalidApplicationUpdate   = fmt.Errorf("%w: application", domain.ErrInvalidUpdate)
)

func newApplication(
	businessID, jobID int,
	businessName, jobTitle string,
	name, email, title string,
	yearsOfExperience, hourlyRateAmount int,
	hourlyRateCurrency string,
	freelancerProfile, resumeURL string,
) (*Application, error) {
	applicationStatus, err := valueobject.NewApplicationStatus("pending")
	if err != nil {
		panic(err)
	}

	if len(name) == 0 || len(name) > 30 {
		return nil, fmt.Errorf("%w: name must be between 0 and 30 bytes", ErrInvalidApplicationProperty)
	}

	if len(title) == 0 || len(title) > 100 {
		return nil, fmt.Errorf("%w: title must be between 0 and 100 bytes", ErrInvalidApplicationProperty)
	}

	if !validator.Matches(email, validator.EmailRX) {
		return nil, fmt.Errorf("%w: invalid email", ErrInvalidApplicationProperty)
	}

	if yearsOfExperience < 0 {
		return nil, fmt.Errorf("%w: invalid years of experience", ErrInvalidApplicationProperty)
	}

	hourlyRate := money.New(int64(hourlyRateAmount), hourlyRateCurrency)

	return &Application{
		businessID:        businessID,
		businessName:      businessName,
		jobID:             jobID,
		jobTitle:          jobTitle,
		status:            applicationStatus,
		name:              name,
		email:             email,
		title:             title,
		yearsOfExperience: yearsOfExperience,
		hourlyRate:        hourlyRate,
		FreelancerProfile: freelancerProfile,
		ResumeURL:         resumeURL,
		createdAt:         time.Now(),
		updatedAt:         time.Now(),
	}, nil
}

func InstantiateApplication(
	businessID, jobID int,
	businessName, jobTitle string,
	id int,
	status, name, email, title string,
	yearsOfExperience, hourlyRateAmount int,
	hourlyRateCurrency string,
	freelancerProfile, resumeURL string,
	createdAt, updatedAt time.Time,
) *Application {
	applicationStatus, err := valueobject.NewApplicationStatus(status)
	if err != nil {
		panic(err)
	}

	hourlyRate := money.New(int64(hourlyRateAmount), hourlyRateCurrency)

	return &Application{
		businessID:        businessID,
		businessName:      businessName,
		jobID:             jobID,
		jobTitle:          jobTitle,
		id:                id,
		status:            applicationStatus,
		name:              name,
		email:             email,
		title:             title,
		yearsOfExperience: yearsOfExperience,
		hourlyRate:        hourlyRate,
		FreelancerProfile: freelancerProfile,
		ResumeURL:         resumeURL,
		createdAt:         createdAt,
		updatedAt:         updatedAt,
	}
}

func (a *Application) ID() int {
	return a.id
}

func (a *Application) Status() string {
	return a.status.Status()
}

func (a *Application) updateStatus(status string) error {
	applicationStatus, err := valueobject.NewApplicationStatus(status)
	if err != nil {
		return fmt.Errorf("%w: invalid application status (pending - accept - reject)", ErrInvalidApplicationUpdate)
	}

	a.status = applicationStatus

	a.updatedAt = time.Now()

	return nil
}

func (a *Application) Name() string {
	return a.name
}

func (a *Application) UpdateName(name string) error {
	if len(name) == 0 || len(name) > 30 {
		return fmt.Errorf("%w: name must be between 0 and 30 bytes", ErrInvalidApplicationUpdate)
	}

	a.name = name
	a.updatedAt = time.Now()
	return nil
}

func (a *Application) Email() string {
	return a.email
}

func (a *Application) Title() string {
	return a.title
}

func (a *Application) UpdateTitle(title string) error {
	if len(title) == 0 || len(title) > 100 {
		return fmt.Errorf("%w: title must be between 0 and 100 bytes", ErrInvalidApplicationUpdate)
	}

	a.title = title
	a.updatedAt = time.Now()

	return nil
}

func (a *Application) YearsOfExperience() int {
	return a.yearsOfExperience
}

func (a *Application) UpdateYearsOfExperience(yearsOfExperience int) error {
	if yearsOfExperience < 0 {
		return fmt.Errorf("%w: invalid years of experience", ErrInvalidApplicationUpdate)
	}

	a.yearsOfExperience = yearsOfExperience
	a.updatedAt = time.Now()

	return nil
}

func (a *Application) CreatedAt() time.Time {
	return a.createdAt
}

func (a *Application) UpdatedAt() time.Time {
	return a.updatedAt
}
