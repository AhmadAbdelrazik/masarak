package job

import (
	"errors"
	"slices"

	"github.com/Rhymond/go-money"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity"
)

var ErrAlreadyApplied = errors.New("applicant already applied to this job")

func (j *AvailableJob) Apply(
	name, email, title string,
	yearsOfExperience int,
	hourlyRate *money.Money,
	freelancerProfile, resumeURL string) error {
	// Check that the user hasn't already applied
	for _, a := range j.applications {
		if a.Email == email {
			return ErrAlreadyApplied
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

func (j *AvailableJob) GetPendingApplications() []entity.Application {
	return j.getApplicationsWithStatus("pending")
}

func (j *AvailableJob) GetAcceptedApplications() []entity.Application {
	return j.getApplicationsWithStatus("accepted")
}

func (j *AvailableJob) getApplicationsWithStatus(status string) []entity.Application {
	appsWithStatus := make([]entity.Application, 0, len(j.applications))

	for _, a := range j.applications {
		if a.Status.Status() == status {
			appsWithStatus = append(appsWithStatus, *a)
		}
	}

	return appsWithStatus
}

var ErrApplicationNotFound = errors.New("application not found")

func (j *AvailableJob) Get(applicantEmail string) (entity.Application, error) {
	for _, a := range j.applications {
		if a.Email == applicantEmail {
			return *a, nil
		}
	}

	return entity.Application{}, ErrApplicationNotFound
}

func (j *AvailableJob) DeleteApplication(email string) error {
	for i, a := range j.applications {
		if a.Email == email && a.Status.IsPending() {
			j.applications = slices.Delete(j.applications, i, i+1)
			return nil
		}
	}

	return ErrApplicationNotFound
}
