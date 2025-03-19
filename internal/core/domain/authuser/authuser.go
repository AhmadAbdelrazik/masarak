package authuser

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/valueobject"
	"github.com/pkg/errors"
)

type AuthUser struct {
	ID       string
	Name     string
	Email    string
	Password *Password
	Role     *valueobject.Role
}

func New(id, name, email, passwordText string, role *valueobject.Role) (*AuthUser, error) {
	password, err := newPassword(passwordText)
	if err != nil {
		return nil, err
	}
	return &AuthUser{
		ID:       id,
		Name:     name,
		Email:    email,
		Role:     role,
		Password: password,
	}, nil
}

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)

type Repository interface {
	Add(ctx context.Context, user *AuthUser) error
	GetByEmail(ctx context.Context, email string) (*AuthUser, error)
	ChangeRole(ctx context.Context, email string, role *valueobject.Role) error
}
