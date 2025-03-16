// The user package is ued to introduce a user for the auth process only
// it should not be dealt with as an aggregate
package users

import (
	"context"

	"github.com/pkg/errors"
)

type User struct {
	Email string
	Name  string
	Role  string
}

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type Repository interface {
	Add(ctx context.Context, user *User) error
	Get(ctx context.Context, email string) (*User, error)
	UpdateRole(ctx context.Context, email, role string) error
}
