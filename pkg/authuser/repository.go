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
	Create(ctx context.Context, username, email, name, passwordText, role string) (*User, error)

	// GetByEmail - returns a user by the email. returns
	// ErrUserNotFound if the user doesn't exist
	GetByEmail(ctx context.Context, email string) (*User, error)

	// GetByUsername - returns a user by Username. returns
	// ErrUserNotFound if the user doesn't exist
	GetByUsername(ctx context.Context, username string) (*User, error)
	// GetByID - returns a user by the id. returns
	// ErrUserNotFound if the user doesn't exist
	GetByID(ctx context.Context, id int) (*User, error)

	// GetByEmail - returns a user by the token. returns
	// ErrUserNotFound if the user doesn't exist
	GetByToken(ctx context.Context, token Token) (*User, error)

	// Update - Update changes to a user (name, password, or role
	// change). returns ErrUserNotFound if user doesn't exist or a
	// generic error.
	Update(
		ctx context.Context,
		id int,
		updateFn func(ctx context.Context, user *User) error,
	) error
}
