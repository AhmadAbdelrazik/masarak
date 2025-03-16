package entity

import "github.com/google/uuid"

type Company struct {
	ID             uuid.UUID
	Name           string
	Details        string
	LineOfBusiness string
}

func NewCompany(name, details, lineOfBusiness string) (*Company, error) {
	return &Company{
		ID:             uuid.New(),
		Name:           name,
		Details:        details,
		LineOfBusiness: lineOfBusiness,
	}, nil
}
