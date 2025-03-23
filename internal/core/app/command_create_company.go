package app

import (
	"context"
)

type CreateCompany struct {
	OwnerEmail            string
	CompanyName           string
	CompanyDetails        string
	CompanyLineOfBusiness string
}

func (h *Commands) CreateCompanyHandler(ctx context.Context, cmd CreateCompany) error {
	owner, err := h.repo.Owner.GetByEmail(ctx, cmd.OwnerEmail)
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

	return h.repo.Companies.Create(ctx, company)
}
