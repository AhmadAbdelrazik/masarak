package app

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/internal/domain/business"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

type AddJob struct {
	User           *authuser.User
	BusinessID     int
	Title          string
	JobDescription string
	WorkLocation   string
	WorkTime       string
	Skills         []string
}

func (c *Commands) AddJobHandler(ctx context.Context, cmd AddJob) (Job, error) {
	if !cmd.User.HasPermission("job.create") {
		return Job{}, ErrUnauthorized
	}

	var jobDTO Job
	_, err := c.repo.Businesses.Update(
		ctx,
		cmd.BusinessID,
		func(ctx context.Context, business *business.Business) error {
			if !business.IsEmployee(cmd.User.Email()) {
				return ErrUnauthorized
			}

			job, err := business.NewJob(
				cmd.Title,
				cmd.JobDescription,
				cmd.WorkLocation,
				cmd.WorkTime,
				cmd.Skills,
			)

			jobDTO = toJob(job)
			return err
		},
	)

	if err != nil {
		return Job{}, err
	}

	return jobDTO, nil
}

type UpdateJob struct {
	User                *authuser.User
	BusinessID          int
	JobID               int
	Title               *string
	JobDescription      *string
	WorkLocation        *string
	WorkTime            *string
	Skills              []string
	ExpectedSalaryRange *struct {
		From     int
		To       int
		Currency string
	}
	YearsOfExperienceRange *struct {
		From int
		To   int
	}
	Status *string
}

func (c *Commands) UpdateJobHandler(ctx context.Context, cmd UpdateJob) (Job, error) {
	if !cmd.User.HasPermission("job.update") {
		return Job{}, ErrUnauthorized
	}

	var jobDTO Job

	_, err := c.repo.Businesses.Update(
		ctx,
		cmd.BusinessID,
		func(ctx context.Context, business *business.Business) error {
			if !business.IsEmployee(cmd.User.Email()) {
				return ErrUnauthorized
			}

			job, err := business.Job(cmd.JobID)
			if err != nil {
				return err
			}

			if cmd.Title != nil {
				if err := business.UpdateJobTitle(cmd.JobID, *cmd.Title); err != nil {
					return err
				}
			}

			if cmd.JobDescription != nil {
				if err := job.UpdateDescription(*cmd.JobDescription); err != nil {
					return err
				}
			}

			if cmd.WorkLocation != nil {
				if err := job.UpdateWorkLocation(*cmd.WorkLocation); err != nil {
					return err
				}
			}

			if cmd.WorkTime != nil {
				if err := job.UpdateWorkTime(*cmd.WorkTime); err != nil {
					return err
				}
			}

			if cmd.ExpectedSalaryRange != nil {
				if err := job.UpdateExpectedSalary(
					cmd.ExpectedSalaryRange.From,
					cmd.ExpectedSalaryRange.To,
					cmd.ExpectedSalaryRange.Currency,
				); err != nil {
					return err
				}
			}

			if cmd.YearsOfExperienceRange != nil {
				if err := job.UpdateYearsOfExperience(
					cmd.YearsOfExperienceRange.From,
					cmd.YearsOfExperienceRange.To,
				); err != nil {
					return err
				}
			}

			if cmd.Status != nil {
				if err := job.UpdateStatus(*cmd.Status); err != nil {
					return nil
				}
			}

			return nil
		},
	)
	if err != nil {
		return Job{}, err
	}

	return jobDTO, nil
}
