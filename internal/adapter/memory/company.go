package memory

import (
	"context"
	"slices"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/company"
	"github.com/google/uuid"
)

type InMemoryCompanyRepository struct {
	memory *Memory
}

func NewInMemoryCompanyRepository(memory *Memory) *InMemoryCompanyRepository {
	return &InMemoryCompanyRepository{
		memory: memory,
	}
}

func (r *InMemoryCompanyRepository) Get(ctx context.Context, uid uuid.UUID) (*company.Company, error) {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, c := range r.memory.companies {
		if c.ID() == uid {
			return c, nil
		}
	}

	return nil, company.ErrCompanyNotFound
}

func (r *InMemoryCompanyRepository) Create(ctx context.Context, c *company.Company) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, cc := range r.memory.companies {
		if cc.ID() == c.ID() {
			return company.ErrAlreadyExists
		}
	}

	r.memory.companies = append(r.memory.companies, c)
	return nil
}
func (r *InMemoryCompanyRepository) Update(ctx context.Context, c *company.Company) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for i, cc := range r.memory.companies {
		if cc.ID() == c.ID() {
			r.memory.companies[i] = c
			return nil
		}
	}

	return company.ErrCompanyNotFound
}
func (r *InMemoryCompanyRepository) Delete(ctx context.Context, uid uuid.UUID) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for i, cc := range r.memory.companies {
		if cc.ID() == uid {
			r.memory.companies = slices.Delete(r.memory.companies, i, i+1)
			return nil
		}
	}

	return company.ErrCompanyNotFound
}

func (r *InMemoryCompanyRepository) GetByOwnerID(ctx context.Context, ownerID uuid.UUID) ([]*company.Company, error) {
	r.memory.Lock()
	defer r.memory.Unlock()

	companies := make([]*company.Company, 0)

	for _, c := range r.memory.companies {
		if c.OwnerID() == ownerID {
			companies = append(companies, c)
		}
	}

	return companies, nil
}
