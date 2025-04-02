package authuser

import (
	"context"

	"github.com/pkg/errors"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)

type Repository interface {
	GetByEmail(ctx context.Context, email string) (*AuthUser, error)
	Create(ctx context.Context, user *AuthUser) error
	Save(ctx context.Context, user *AuthUser) error
}
