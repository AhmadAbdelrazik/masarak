package business

import (
	"github.com/Rhymond/go-money"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity"
	"github.com/google/uuid"
)

func (b *Business) ApplyToJob(
	jobID uuid.UUID,
	name, email, title string,
	yearsOfExperience int,
	hourlyRate *money.Money,
	freelancerProfile, resumeURL string,
) error {
	job, err := b.getJobByID(jobID)
	if err != nil {
		return err
	}

	return job.Apply(
		name,
		email,
		title,
		yearsOfExperience,
		hourlyRate,
		freelancerProfile,
		resumeURL,
	)
}

// CancelJobApplication - Cancel the application and deletes it as long as it's still pending
func (b *Business) CancelJobApplication(jobID, applicationID uuid.UUID) error {
	job, err := b.getJobByID(jobID)
	if err != nil {
		return err
	}

	return job.CancelApplication(applicationID)
}

func (b *Business) UpdateJobApplication(
	jobID, applicationID uuid.UUID,
	name, title string,
	yearsOfExperience int,
	hourlyRate *money.Money,
	freelancerProfile, resumeURL string,
) error {
	job, err := b.getJobByID(jobID)
	if err != nil {
		return err
	}

	return job.UpdateApplication(
		applicationID,
		name,
		title,
		yearsOfExperience,
		hourlyRate,
		freelancerProfile,
		resumeURL,
	)
}

func (b *Business) AcceptJobApplication(jobID, applicationID uuid.UUID) error {
	job, err := b.getJobByID(jobID)
	if err != nil {
		return err
	}

	return job.AcceptApplication(applicationID)
}

func (b *Business) RejectJobApplication(jobID, applicationID uuid.UUID) error {
	job, err := b.getJobByID(jobID)
	if err != nil {
		return err
	}

	return job.RejectApplication(applicationID)
}

func (b *Business) SetApplicationStatusToPending(jobID, applicationID uuid.UUID) error {
	job, err := b.getJobByID(jobID)
	if err != nil {
		return err
	}

	return job.SetApplicationStatusToPending(applicationID)

}

func (b *Business) GetApplicationByID(jobID, applicationID uuid.UUID) (entity.Application, error) {
	job, err := b.getJobByID(jobID)
	if err != nil {
		return entity.Application{}, err
	}

	return job.GetApplicationByID(applicationID)
}

func (b *Business) GetApplicationByEmail(jobID uuid.UUID, applicantEmail string) (entity.Application, error) {
	job, err := b.getJobByID(jobID)
	if err != nil {
		return entity.Application{}, err
	}

	return job.GetApplicationByEmail(applicantEmail)

}

func (b *Business) GetAcceptedApplications(jobID uuid.UUID) ([]entity.Application, error) {
	job, err := b.getJobByID(jobID)
	if err != nil {
		return []entity.Application{}, err
	}

	return job.GetAcceptedApplications(), nil
}

func (b *Business) GetRejectedApplications(jobID uuid.UUID) ([]entity.Application, error) {
	job, err := b.getJobByID(jobID)
	if err != nil {
		return []entity.Application{}, err
	}

	return job.GetRejectedApplications(), nil
}

func (b *Business) GetPendingApplications(jobID uuid.UUID) ([]entity.Application, error) {
	job, err := b.getJobByID(jobID)
	if err != nil {
		return []entity.Application{}, err
	}

	return job.GetPendingApplications(), nil
}

func (b *Business) GetAllApplications(jobID uuid.UUID) ([]entity.Application, error) {
	job, err := b.getJobByID(jobID)
	if err != nil {
		return []entity.Application{}, err
	}

	return job.GetAllApplications(), nil
}
