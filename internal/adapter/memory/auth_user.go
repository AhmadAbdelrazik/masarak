package memory

import (
	"context"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
)

type InMemoryAuthUserRepository struct {
	memory *Memory
}

func NewInMemoryAuthUserRepository(memory *Memory) *InMemoryAuthUserRepository {
	return &InMemoryAuthUserRepository{
		memory: memory,
	}
}

func (r *InMemoryAuthUserRepository) Add(ctx context.Context, user *entity.AuthUser) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, u := range r.memory.authUsers {
		if u.Email == user.Email {
			return entity.ErrUserAlreadyExists
		}
	}

	return nil
}

func (r *InMemoryAuthUserRepository) GetByEmail(ctx context.Context, email string) (*entity.AuthUser, error) {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, u := range r.memory.authUsers {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, entity.ErrUserNotFound
}
