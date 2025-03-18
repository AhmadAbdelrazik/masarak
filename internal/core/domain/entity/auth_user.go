package entity

import (
	"context"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/valueobject"
	"github.com/pkg/errors"
)

type AuthUser struct {
	ID       string
	Name     string
	Email    string
	Password *valueobject.Password
	Role     *valueobject.Role
}

func NewAuthUser(id, name, email, passwordText string, role *valueobject.Role) (*AuthUser, error) {
	password, err := valueobject.NewPassword(passwordText)
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

type AuthUserRepository interface {
	Add(ctx context.Context, user *AuthUser) error
	GetByEmail(ctx context.Context, email string) (*AuthUser, error)
	ChangeRole(ctx context.Context, email string, role *valueobject.Role) error
}
