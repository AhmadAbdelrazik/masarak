package applicationshistory

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type AppliedJobs struct {
	ID        uuid.UUID
	ApplyDate time.Time
}

type ApplicationHistory struct {
	freelancerEmail string
	appliedJobs     []AppliedJobs
}

func New(freelancerEmail string) *ApplicationHistory {
	return &ApplicationHistory{
		freelancerEmail: freelancerEmail,
		appliedJobs:     make([]AppliedJobs, 0),
	}
}

var (
	ErrAlreadySubmittedToJob = errors.New("already submitted to this job")
)

func (a *ApplicationHistory) Add(jobID uuid.UUID) error {
	for _, j := range a.appliedJobs {
		if j.ID == jobID {
			return ErrAlreadySubmittedToJob
		}
	}

	job := AppliedJobs{
		ID:        jobID,
		ApplyDate: time.Now(),
	}

	a.appliedJobs = append(a.appliedJobs, job)
	return nil
}

func (a *ApplicationHistory) IsAppliedTo(jobID uuid.UUID) bool {
	for _, j := range a.appliedJobs {
		if j.ID == jobID {
			return true
		}
	}

	return false
}
