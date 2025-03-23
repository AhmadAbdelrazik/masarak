package app

import (
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/company"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/owner"
)

type OwnerDTO struct {
	Name           string       `json:"name"`
	Email          string       `json:"email"`
	OwnedCompanies []CompanyDTO `json:"owned_companies"`
}

func NewOwnerDTO(owner *owner.Owner, ownedCompanies []*company.Company) OwnerDTO {
	companies := make([]CompanyDTO, 0, len(ownedCompanies))

	for _, c := range ownedCompanies {
		companies = append(companies, CompanyDTO{Name: c.Name()})
	}

	return OwnerDTO{
		Name:           owner.Name(),
		Email:          owner.Email(),
		OwnedCompanies: companies,
	}
}
