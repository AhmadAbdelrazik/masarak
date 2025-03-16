package adapter

import (
	"context"
	"sync"

	"github.com/ahmadabdelrazik/linkedout/internal/domain/user"
)

type InMemoryUserRepository struct {
	users []*users.User

	sync.Mutex
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make([]*users.User, 0),
	}
}

func (r *InMemoryUserRepository) Add(ctx context.Context, user *users.User) error {
	r.Lock()
	defer r.Unlock()

	for _, u := range r.users {
		if u.Email == user.Email {
			return users.ErrUserAlreadyExists
		}
	}

	r.users = append(r.users, user)

	return nil
}
func (r *InMemoryUserRepository) Get(ctx context.Context, email string) (*users.User, error) {
	r.Lock()
	defer r.Unlock()

	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, users.ErrUserNotFound
}

func (r *InMemoryUserRepository) UpdateRole(ctx context.Context, email, role string) error {
	r.Lock()
	defer r.Unlock()

	for _, u := range r.users {
		if u.Email == email {
			u.Role = role
			return nil
		}
	}

	return users.ErrUserNotFound
}
