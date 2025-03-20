package memory

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/valueobject"
)

type InMemoryAuthUserRepository struct {
	memory *Memory
}

func NewInMemoryAuthUserRepository(memory *Memory) *InMemoryAuthUserRepository {
	return &InMemoryAuthUserRepository{
		memory: memory,
	}
}

func (r *InMemoryAuthUserRepository) Add(ctx context.Context, user *authuser.AuthUser) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, u := range r.memory.authUsers {
		if u.Email == user.Email {
			return authuser.ErrUserAlreadyExists
		}
	}

	r.memory.authUsers = append(r.memory.authUsers, *user)

	return nil
}

func (r *InMemoryAuthUserRepository) GetByEmail(ctx context.Context, email string) (*authuser.AuthUser, error) {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, u := range r.memory.authUsers {
		if u.Email == email {
			return &u, nil
		}
	}

	return nil, authuser.ErrUserNotFound
}

func (r *InMemoryAuthUserRepository) ChangeRole(ctx context.Context, email string, role *valueobject.Role) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for i, u := range r.memory.authUsers {
		if u.Email == email {
			r.memory.authUsers[i].Role = role
			return nil
		}
	}

	return authuser.ErrUserNotFound
}
