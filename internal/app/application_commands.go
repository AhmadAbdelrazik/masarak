package app

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/internal/domain/business"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

type ApplyToJob struct {
	User               *authuser.User
	BusinessID         int
	JobID              int
	Name               string
	Email              string
	Title              string
	YearsOfExperience  int
	HourlyRateAmount   int
	HourlyRateCurrency string
	FreelancerProfile  string
	ResumeURL          string
}

func (c *Commands) ApplyToJobHandler(ctx context.Context, cmd ApplyToJob) (JobApplication, error) {
	var applicationDTO JobApplication
	if !cmd.User.HasPermission("application.create") {
		return JobApplication{}, ErrUnauthorized
	}

	_, err := c.repo.Businesses.Update(
		ctx,
		cmd.BusinessID,
		func(ctx context.Context, business *business.Business) error {

			job, err := business.Job(cmd.JobID)
			if err != nil {
				return err
			}

			application, err := job.NewApplication(
				cmd.Name,
				cmd.Email,
				cmd.Title,
				cmd.YearsOfExperience,
				cmd.HourlyRateAmount,
				cmd.HourlyRateCurrency,
				cmd.FreelancerProfile,
				cmd.ResumeURL,
			)
			if err != nil {
				return err
			}

			applicationDTO = toApplication(application)

			return nil
		},
	)
	if err != nil {
		return JobApplication{}, err
	}

	return applicationDTO, nil
}

type UpdateApplication struct {
	User              *authuser.User
	BusinessID        int
	JobID             int
	ApplicationID     int
	Name              *string
	Title             *string
	YearsOfExperience *int
	FreelancerProfile *string
	ResumeURL         *string
}

func (c *Commands) UpdateApplicationHandler(ctx context.Context, cmd UpdateApplication) (JobApplication, error) {
	var applicationDTO JobApplication

	_, err := c.repo.Businesses.Update(
		ctx,
		cmd.BusinessID,
		func(ctx context.Context, business *business.Business) error {
			job, err := business.Job(cmd.JobID)
			if err != nil {
				return err
			}

			if job.Status() != "open" {
				return ErrUnauthorized
			}

			application, err := job.Application(cmd.ApplicationID)
			if err != nil {
				return err
			}

			if !cmd.User.HasPermission("application.update") || application.Email() != cmd.User.Email() {
				return ErrUnauthorized
			}

			if cmd.Name != nil {
				if err := application.UpdateName(*cmd.Name); err != nil {
					return err
				}
			}

			if cmd.Title != nil {
				if err := application.UpdateTitle(*cmd.Title); err != nil {
					return err
				}
			}

			if cmd.YearsOfExperience != nil {
				if err := application.UpdateYearsOfExperience(*cmd.YearsOfExperience); err != nil {
					return err
				}
			}

			if cmd.FreelancerProfile != nil {
				application.FreelancerProfile = *cmd.FreelancerProfile
			}

			if cmd.ResumeURL != nil {
				application.ResumeURL = *cmd.ResumeURL
			}

			applicationDTO = toApplication(application)

			return nil
		},
	)
	if err != nil {
		return JobApplication{}, err
	}

	return applicationDTO, nil
}

type UpdateApplicationStatus struct {
	User          *authuser.User
	BusinessID    int
	JobID         int
	ApplicationID int
	Status        string
}

func (c *Commands) UpdateApplicationStatusHandler(ctx context.Context, cmd UpdateApplicationStatus) error {
	_, err := c.repo.Businesses.Update(
		ctx,
		cmd.BusinessID,
		func(ctx context.Context, business *business.Business) error {
			if !cmd.User.HasPermission("application.update_status") || !business.IsEmployee(cmd.User.Email()) {
				return ErrUnauthorized
			}

			job, err := business.Job(cmd.JobID)
			if err != nil {
				return err
			}

			return job.UpdateApplicationStatus(cmd.ApplicationID, cmd.Status)
		},
	)

	return err
}
