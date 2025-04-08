package authuser

import (
	"context"

	"github.com/pkg/errors"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidProperty   = errors.New("invalid property")
)

type UserRepository interface {
	// Create - creates a new user and save it to the database.
	// returns ErrUserAlreadyExists if email is already used
	Create(ctx context.Context, name, email, passwordText, role string) error

	// GetByEmail - returns a user by the email. returns
	// ErrUserNotFound if the user doesn't exist
	GetByEmail(ctx context.Context, email string) (*User, error)

	// GetByEmail - returns a user by the token. returns
	// ErrUserNotFound if the user doesn't exist
	GetByToken(ctx context.Context, token Token) (*User, error)

	// Save - Save changes to a user (name, password, or role
	// change). returns ErrUserNotFound if user doesn't exist or a
	// generic error.
	Save(
		ctx context.Context,
		email string,
		updateFn func(ctx context.Context, user *User) error,
	) error
}
