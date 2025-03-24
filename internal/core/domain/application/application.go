package application

import (
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/valueobject"
	"github.com/google/uuid"
)

type Application struct {
	id               uuid.UUID
	status           *valueobject.ApplicationStatus
	applicantDetails *entity.ApplicantDetails
}

func New(applicantDetails *entity.ApplicantDetails, status *valueobject.ApplicationStatus) *Application {
	return &Application{
		id:               uuid.New(),
		status:           status,
		applicantDetails: applicantDetails,
	}
}

func (a *Application) Accept() string {
	panic("not implemented")
}

func (a *Application) Reject() string {
	panic("not implemented")
}
