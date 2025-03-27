package job

import (
	"errors"
	"slices"

	"github.com/Rhymond/go-money"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity"
	"github.com/google/uuid"
)

var (
	ErrDuplicateApplication = errors.New("duplicate application")
	ErrJobNotAvailable      = errors.New("job not available")
	ErrApplicationNotFound  = errors.New("application not found")
	ErrApplicationReviewd   = errors.New("can't edit or cancel a reviewd application")
)

func (j *Job) Apply(name, email, title string,
	yearsOfExperience int,
	hourlyRate *money.Money,
	freelancerProfile, resumeURL string,
) error {
	if !j.Status.IsAvailable() {
		return ErrJobNotAvailable
	}

	for _, a := range j.applications {
		if a.Email == email {
			return ErrDuplicateApplication
		}
	}

	application := entity.NewApplication(
		name,
		email,
		title,
		yearsOfExperience,
		hourlyRate,
		freelancerProfile,
		resumeURL,
	)

	j.applications = append(j.applications, application)

	return nil
}

func (j *Job) CancelApplication(applicationID uuid.UUID) error {
	for i, a := range j.applications {
		if a.ID == applicationID {
			if a.Status() != "pending" {
				return ErrApplicationReviewd
			}

			j.applications = slices.Delete(j.applications, i, i+1)
			return nil
		}
	}

	return ErrApplicationNotFound
}

func (j *Job) UpdateApplication(
	applicationID uuid.UUID,
	name, email, title string,
	yearsOfExperience int,
	hourlyRate *money.Money,
	freelancerProfile, resumeURL string,
) error {
	for _, a := range j.applications {
		if a.ID == applicationID {
			return a.Update(
				name,
				email,
				title,
				yearsOfExperience,
				hourlyRate,
				freelancerProfile,
				resumeURL,
			)
		}
	}

	return ErrApplicationNotFound
}

func (j *Job) AcceptApplication(applicationID uuid.UUID) error {
	for _, a := range j.applications {
		if a.ID == applicationID {
			a.Accept()
			return nil
		}
	}

	return ErrApplicationNotFound
}

func (j *Job) RejectApplication(applicationID uuid.UUID) error {
	for _, a := range j.applications {
		if a.ID == applicationID {
			a.Reject()
			return nil
		}
	}

	return ErrApplicationNotFound
}

func (j *Job) GetAcceptedApplications() []entity.Application {
	return j.getApplicationsByStatus("accepted")
}

func (j *Job) GetRejectedApplications() []entity.Application {
	return j.getApplicationsByStatus("rejected")
}

func (j *Job) GetPendingApplications() []entity.Application {
	return j.getApplicationsByStatus("pending")
}

func (j *Job) GetAllApplications() []entity.Application {
	applications := make([]entity.Application, 0, len(j.applications))

	for _, a := range j.applications {
		applications = append(applications, *a)
	}

	return applications
}

func (j *Job) getApplicationsByStatus(status string) []entity.Application {
	applications := make([]entity.Application, 0, len(j.applications))

	for _, a := range j.applications {
		if a.Status() == status {
			applications = append(applications, *a)
		}
	}

	return applications
}
