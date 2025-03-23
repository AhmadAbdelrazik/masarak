package app

import (
	"context"
	"errors"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/job"
)

type CreateJob struct {
	Email          string
	CompanyName    string
	JobTitle       string
	JobDescription string
}

var (
	ErrInvalidOwner = errors.New("owner does not own this company")
)

func (h *Commands) CreateJobHandler(ctx context.Context, cmd CreateJob) error {
	owner, err := h.repo.Owner.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return err
	}

	company, err := h.repo.Companies.GetByName(ctx, cmd.CompanyName)
	if err != nil {
		return err
	}

	if owner.ID() != company.ID() {
		return ErrInvalidOwner
	}

	job, err := job.New(cmd.JobTitle, cmd.JobDescription, company.ID())
	if err != nil {
		return err
	}

	return h.repo.Jobs.Create(ctx, job)
}
