package app

import (
	"context"
	"errors"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
)

// UserLogin - Get user for login purposes. this method validates password
// before returning the User.
func (q *Queries) UserLogin(ctx context.Context, email, password string) (User, error) {
	user, err := q.repo.Users.GetByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}

	return User{
		Name:  user.Name(),
		Email: user.Email(),
		Role:  user.Role(),
	}, nil
}
