package app

import "context"

// GetUser - Returns a user from the database using email. This is not used
// for login validation, use UserLogin instead. if the user is not found,
// ErrUserNotFound is returned
func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	user, err := q.repo.AuthUsers.GetByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}

	return User{
		Name:  user.Name(),
		Email: user.Email(),
		Role:  user.Role(),
	}, nil
}
