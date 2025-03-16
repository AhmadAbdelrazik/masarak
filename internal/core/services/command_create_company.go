package app

import (
	"context"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/company"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/owner"
)

type CreateCompany struct {
	OwnerEmail            string
	CompanyName           string
	CompanyDetails        string
	CompanyLineOfBusiness string
}

type CreateCompanyHandler struct {
	companyRepo company.Repository
	ownerRepo   owner.Repository
}

func NewCreateCompanyHandler(companyRepo company.Repository, ownerRepo owner.Repository) *CreateCompanyHandler {
	if companyRepo == nil {
		panic("company repo not found")
	}

	if ownerRepo == nil {
		panic("user repo not found")
	}
	return &CreateCompanyHandler{
		companyRepo: companyRepo,
		ownerRepo:   ownerRepo,
	}
}

func (h *CreateCompanyHandler) Handle(ctx context.Context, cmd CreateCompany) error {
	owner, err := h.ownerRepo.GetByEmail(ctx, cmd.OwnerEmail)
	if err != nil {
		return err
	}

	company, err := owner.CreateCompany(
		cmd.CompanyName,
		cmd.CompanyDetails,
		cmd.CompanyLineOfBusiness,
	)
	if err != nil {
		return err
	}

	return h.companyRepo.Create(ctx, company)
}
