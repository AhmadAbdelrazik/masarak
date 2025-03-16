package entity

import (
	"context"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/valueobject"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type AuthUser struct {
	ID    uuid.UUID
	Name  string
	Email string
	Role  *valueobject.Role
}

func NewAuthUser(name, email string, role *valueobject.Role) (*AuthUser, error) {
	return &AuthUser{
		ID:    uuid.New(),
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
