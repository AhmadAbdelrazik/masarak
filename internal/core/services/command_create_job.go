package app

import (
	"context"
	"errors"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/company"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/job"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/owner"
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

func (h *CreateJobHandler) Handle(ctx context.Context, cmd CreateJob) error {
	owner, err := h.ownerRepo.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return err
	}

	companies, err := h.companyRepo.GetByOwnerID(ctx, owner.ID())
	if err != nil {
		return err
	}

	var targetCompany *company.Company

	for _, c := range companies {
		if c.Name() == cmd.CompanyName {
			targetCompany = c
			break
		}
	}

	if targetCompany == nil {
		return errors.New("company not found")
	}

	job, err := job.New(cmd.JobTitle, cmd.JobDescription, targetCompany.ID())
	if err != nil {
		return err
	}

	return h.jobRepo.Create(ctx, job)
}
