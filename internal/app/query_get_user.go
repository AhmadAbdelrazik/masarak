package app

import (
	"context"
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
