package app

import (
	"context"
	"errors"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/company"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/job"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/owner"
)

type CreateJob struct {
	Email          string
	CompanyName    string
	JobTitle       string
	JobDescription string
}

type CreateJobHandler struct {
	ownerRepo   owner.Repository
	companyRepo company.Repository
	jobRepo     job.Repository
}

func NewCreateJobHandler(
	ownerRepo owner.Repository,
	companyRepo company.Repository,
	jobRepo job.Repository,
) *CreateJobHandler {
	if companyRepo == nil {
		panic("company repo not found")
	}

	if ownerRepo == nil {
		panic("user repo not found")
	}

	if jobRepo == nil {
		panic("job repo not found")
	}

	return &CreateJobHandler{}
}

var (
	ErrInvalidOwner = errors.New("owner does not own this company")
)

func (h *CreateJobHandler) Handle(ctx context.Context, cmd CreateJob) error {
	owner, err := h.ownerRepo.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return err
	}

	company, err := h.companyRepo.GetByName(ctx, cmd.CompanyName)
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

	return h.jobRepo.Create(ctx, job)
}
