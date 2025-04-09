package app

import "context"

type CreateUser struct {
	Name     string
	Email    string
	Password string
	Role     string
}

// CreateUserHandler - Creates a new user in the system. if the user already exists, ErrUserAlreadyExists will return. if any field fails validation check, ErrInvalidProperty will return.
func (c *Commands) CreateUserHandler(ctx context.Context, cmd CreateUser) error {
	err := c.repo.Users.Create(
		ctx,
		cmd.Name,
		cmd.Email,
		cmd.Password,
		cmd.Role,
	)

	return err
}
