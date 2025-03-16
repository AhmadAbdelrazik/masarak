package memory

import (
	"context"
	"slices"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/owner"
	"github.com/google/uuid"
)

type InMemoryOwnerRepository struct {
	memory *Memory
}

func NewInMemoryOwnerRepository(memory *Memory) *InMemoryOwnerRepository {
	return &InMemoryOwnerRepository{
		memory: memory,
	}
}

func (r *InMemoryOwnerRepository) Get(ctx context.Context, uid uuid.UUID) (*owner.Owner, error) {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, o := range r.memory.owners {
		if o.ID() == uid {
			return o, nil
		}
	}

	return nil, owner.ErrOwnerNotFound
}

func (r *InMemoryOwnerRepository) GetByEmail(ctx context.Context, email string) (*owner.Owner, error) {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, o := range r.memory.owners {
		if o.Email() == email {
			return o, nil
		}
	}

	return nil, owner.ErrOwnerNotFound
}

func (r *InMemoryOwnerRepository) Create(ctx context.Context, o *owner.Owner) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, oo := range r.memory.owners {
		if oo.ID() == o.ID() {
			return owner.ErrAlreadyExists
		}
	}

	r.memory.owners = append(r.memory.owners, o)
	return nil
}
func (r *InMemoryOwnerRepository) Update(ctx context.Context, o *owner.Owner) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for i, oo := range r.memory.owners {
		if oo.ID() == o.ID() {
			r.memory.owners[i] = o
			return nil
		}
	}

	return owner.ErrOwnerNotFound
}
func (r *InMemoryOwnerRepository) Delete(ctx context.Context, uid uuid.UUID) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for i, oo := range r.memory.owners {
		if oo.ID() == uid {
			r.memory.owners = slices.Delete(r.memory.owners, i, i+1)
			return nil
		}
	}

	return owner.ErrOwnerNotFound
}
