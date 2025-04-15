package business

import (
	"time"

	"github.com/Rhymond/go-money"
	"github.com/ahmadabdelrazik/masarak/internal/domain/valueobject"
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
