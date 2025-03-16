package company

import (
	"context"
	"errors"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
	"github.com/google/uuid"
)

type Company struct {
	company      *entity.Company
	ownerID      uuid.UUID
	postedJobIDs uuid.UUIDs
}

func (c *Company) ID() uuid.UUID {
	return c.company.ID
}

func New(name, details, lineOfBusiness string, ownerID uuid.UUID) (*Company, error) {
	company, err := entity.NewCompany(name, details, lineOfBusiness)
	if err != nil {
		return nil, err
	}

	return &Company{
		company:      company,
		ownerID:      ownerID,
		postedJobIDs: uuid.UUIDs{},
	}, nil
}

var (
	ErrCompanyNotFound = errors.New("company not found")
	ErrAlreadyExists   = errors.New("company already exists")
)

type Repository interface {
	Get(ctx context.Context, uid uuid.UUID) (*Company, error)
	Create(ctx context.Context, company *Company) error
	Update(ctx context.Context, company *Company) error
	Delete(ctx context.Context, uid uuid.UUID) error
}
