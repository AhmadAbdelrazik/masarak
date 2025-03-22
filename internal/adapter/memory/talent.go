package memory

import (
	"context"
	"slices"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/talent"
	"github.com/google/uuid"
)

type InMemoryTalentRepository struct {
	memory *Memory
}

func NewInMemoryTalentRepository(memory *Memory) *InMemoryTalentRepository {
	return &InMemoryTalentRepository{
		memory: memory,
	}
}

func (r *InMemoryTalentRepository) Get(ctx context.Context, uid uuid.UUID) (*talent.Talent, error) {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, o := range r.memory.talents {
		if o.ID() == uid {
			return &o, nil
		}
	}

	return nil, talent.ErrTalentNotFound
}

func (r *InMemoryTalentRepository) GetByEmail(ctx context.Context, email string) (*talent.Talent, error) {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, o := range r.memory.talents {
		if o.Email() == email {
			return &o, nil
		}
	}

	return nil, talent.ErrTalentNotFound
}

func (r *InMemoryTalentRepository) Create(ctx context.Context, o *talent.Talent) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, oo := range r.memory.talents {
		if oo.ID() == o.ID() {
			return talent.ErrAlreadyExists
		}
	}

	r.memory.talents = append(r.memory.talents, *o)
	return nil
}
func (r *InMemoryTalentRepository) Update(ctx context.Context, o *talent.Talent) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for i, oo := range r.memory.talents {
		if oo.ID() == o.ID() {
			r.memory.talents[i] = *o
			return nil
		}
	}

	return talent.ErrTalentNotFound
}
func (r *InMemoryTalentRepository) Delete(ctx context.Context, uid uuid.UUID) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for i, oo := range r.memory.talents {
		if oo.ID() == uid {
			r.memory.talents = slices.Delete(r.memory.talents, i, i+1)
			return nil
		}
	}

	return talent.ErrTalentNotFound
}
