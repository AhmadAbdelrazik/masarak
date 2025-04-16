package app

import (
	"context"
	"errors"
)

type GetUser struct {
	Email string
}

// GetUserHandler - Returns a user from the database using email. This is not used
// for login validation, use UserLogin instead. if the user is not found,
// ErrUserNotFound is returned
func (q *Queries) GetUserHandler(ctx context.Context, cmd GetUser) (User, error) {
	user, err := q.repo.Users.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return User{}, err
	}

	return toUserDTO(user), nil
}

type UserLogin struct {
	Email    string
	Password string
}

var (
	ErrInvalidPassword = errors.New("invalid password")
)

// UserLogin - Get user for login purposes. this method validates password
// before returning the User.
func (q *Queries) UserLogin(ctx context.Context, cmd UserLogin) (User, error) {
	user, err := q.repo.Users.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return User{}, err
	}

	if match, err := user.Password.Matches(cmd.Password); err != nil {
		return User{}, err
	} else if !match {
		return User{}, ErrInvalidPassword
	}

	return toUserDTO(user), nil
}
