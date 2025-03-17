package entity

import (
	"context"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/valueobject"
	"github.com/pkg/errors"
)

type AuthUser struct {
	ID    string
	Name  string
	Email string
	Role  *valueobject.Role
}

func NewAuthUser(id, name, email string, role *valueobject.Role) (*AuthUser, error) {
	return &AuthUser{
		ID:    id,
		Name:  name,
		Email: email,
		Role:  role,
	}, nil
}

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)

type AuthUserRepository interface {
	Add(ctx context.Context, user *AuthUser) error
	GetByEmail(ctx context.Context, email string) (*AuthUser, error)
	ChangeRole(ctx context.Context, email string, role *valueobject.Role) error
}
