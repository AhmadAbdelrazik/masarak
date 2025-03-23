package app

import "github.com/ahmadabdelrazik/masarak/internal/core/domain/company"

type CompanyDTO struct {
	Name string `json:"name"`
}

func NewCompanyDTO(company *company.Company) CompanyDTO {
	return CompanyDTO{
		Name: company.Name(),
	}
}
