package app

import (
	"context"
	"errors"
)

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
