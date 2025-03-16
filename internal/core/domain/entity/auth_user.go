package entity

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type AuthUser struct {
	ID    uuid.UUID
	Name  string
	Email string
	Role  string
}

func NewAuthUser(name, email, role string) (*AuthUser, error) {
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
}
