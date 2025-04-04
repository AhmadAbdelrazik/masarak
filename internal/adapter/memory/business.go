package memory

import (
	"context"
	"fmt"
	"strings"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/business"
	"github.com/ahmadabdelrazik/masarak/pkg/filters"
	"github.com/google/uuid"
)

type InMemoryBusinessRepository struct {
	memory *Memory
}

func NewInMemoryBusinessRepository(mem *Memory) *InMemoryBusinessRepository {
	return &InMemoryBusinessRepository{
		memory: mem,
	}
}

func (r *InMemoryBusinessRepository) Create(
	ctx context.Context,
	name, businessEmail, ownerEmail, description, imageURL string,
) (*business.Business, error) {
	b, err := business.NewBusiness(name, businessEmail, ownerEmail, description, imageURL)
	if err != nil {
		return nil, err
	}

	r.memory.Lock()
	defer r.memory.Unlock()

	err = r.checkValid(ctx, b)
	if err != nil {
		return nil, err
	}

	r.memory.businesses = append(r.memory.businesses, b)

	return b, nil
}

func (r *InMemoryBusinessRepository) GetByID(ctx context.Context, businessID uuid.UUID) (*business.Business, error) {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, b := range r.memory.businesses {
		if b.ID() == businessID {
			return b, nil
		}
	}

	return nil, business.ErrBusinessNotFound
}

func (r *InMemoryBusinessRepository) GetIDs(ids []uuid.UUID) ([]*business.Business, error) {
	idSet := make(map[uuid.UUID]bool)

	for _, id := range ids {
		idSet[id] = true
	}

	r.memory.Lock()
	defer r.memory.Unlock()

	businesses := make([]*business.Business, 0, len(ids))

	for _, b := range r.memory.businesses {
		if _, ok := idSet[b.ID()]; ok {
			businesses = append(businesses, b)
			delete(idSet, b.ID())
		}
	}

	if len(idSet) != 0 {
		return nil, fmt.Errorf("invalid businesses %v", idSet)
	}

	return businesses, nil
}

func (r *InMemoryBusinessRepository) GetAll(ctx context.Context, name string, filter filters.Filter) ([]*business.Business, error) {
	r.memory.Lock()
	defer r.memory.Unlock()

	result := make([]*business.Business, 0, len(r.memory.businesses))

	for _, b := range r.memory.businesses {
		if strings.Contains(b.Name(), name) {
			result = append(result, b)
		}
	}

	from := min(len(result), filter.Offset())
	to := min(len(result), filter.Offset()+filter.Offset()+filter.Limit())

	return result[from:to], nil
}

func (r *InMemoryBusinessRepository) Save(ctx context.Context, b *business.Business) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	err := r.checkValid(ctx, b)
	if err != nil {
		return err
	}

	for i, bb := range r.memory.businesses {
		if bb.ID() == b.ID() {
			r.memory.businesses[i] = b
			return nil
		}
	}

	r.memory.businesses = append(r.memory.businesses, b)
	return nil
}

func (r *InMemoryBusinessRepository) checkValid(_ context.Context, b *business.Business) error {
	for _, bb := range r.memory.businesses {
		if bb.ID() == b.ID() {
			continue
		}
		if bb.Name() == b.Name() || bb.Email() == b.Email() {
			return business.ErrDuplicateBusiness
		}
	}

	return nil
}
