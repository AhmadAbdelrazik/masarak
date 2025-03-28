package applicationshistory

import (
	"errors"

	"github.com/google/uuid"
)

type ApplicationHistory struct {
	freelancerEmail string
	businesses      []uuid.UUID
}

func New(freelancerEmail string) *ApplicationHistory {
	return &ApplicationHistory{
		freelancerEmail: freelancerEmail,
		businesses:      make([]uuid.UUID, 0),
	}
}

var (
	ErrAlreadySubmittedToJob = errors.New("already submitted to this job")
)

func (a *ApplicationHistory) Add(businessID uuid.UUID) error {
	a.businesses = append(a.businesses, businessID)
	return nil
}
