package company

import (
	"context"
	"errors"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity"
	"github.com/google/uuid"
)

type Company struct {
	company *entity.Company
	ownerID uuid.UUID
}

func (c *Company) OwnerID() uuid.UUID {
	return c.ownerID
}

func (c *Company) Name() string {
	return c.company.Name
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
		company: company,
		ownerID: ownerID,
	}, nil
}

var (
	ErrCompanyNotFound = errors.New("company not found")
	ErrAlreadyExists   = errors.New("company already exists")
)

type Repository interface {
	Get(ctx context.Context, uid uuid.UUID) (*Company, error)
	GetByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]*Company, error)
	GetByName(ctx context.Context, name string) (*Company, error)
	Create(ctx context.Context, company *Company) error
	Update(ctx context.Context, company *Company) error
	Delete(ctx context.Context, uid uuid.UUID) error
}
