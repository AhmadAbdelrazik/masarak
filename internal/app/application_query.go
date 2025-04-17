package app

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/internal/domain/business"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

type GetApplication struct {
	User       *authuser.User
	BusinessID int
	JobID      int
	Email      string
}

func (q *Queries) GetApplicationHandler(ctx context.Context, cmd GetApplication) (JobApplication, error) {
	var applicationDTO JobApplication

	_, err := q.repo.Businesses.Update(
		ctx,
		cmd.BusinessID,
		func(ctx context.Context, business *business.Business) error {
			job, err := business.Job(cmd.JobID)
			if err != nil {
				return err
			}

			application, err := job.ApplicationByEmail(cmd.User.Email())
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

type GetApplicationsByStatus struct {
	User       *authuser.User
	BusinessID int
	JobID      int
	Status     string // can be open - closed - archived
}

func (q *Queries) GetApplicationsByStatusHandler(ctx context.Context, cmd GetApplicationsByStatus) ([]JobApplication, error) {
	var applicationDTO []JobApplication

	_, err := q.repo.Businesses.Update(
		ctx,
		cmd.BusinessID,
		func(ctx context.Context, business *business.Business) error {
			if !cmd.User.HasPermission("application.read") || !business.IsEmployee(cmd.User.Email()) {
				return ErrUnauthorized
			}
			job, err := business.Job(cmd.JobID)
			if err != nil {
				return err
			}

			applications, err := job.ApplicationsByStatus(cmd.Status)
			if err != nil {
				return err
			}

			applicationDTO = make([]JobApplication, 0, len(applications))

			for _, a := range applications {
				applicationDTO = append(applicationDTO, toApplication(a))
			}

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return applicationDTO, nil
}
