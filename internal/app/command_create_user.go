package app

import "context"

type CreateUser struct {
	Username string
	Email    string
	Name     string
	Password string
	Role     string
}

// CreateUserHandler - Creates a new user in the system. if the user already
// exists, ErrUserAlreadyExists will return. if any field fails validation
// check, ErrInvalidProperty will return.
func (c *Commands) CreateUserHandler(ctx context.Context, cmd CreateUser) (User, error) {
	user, err := c.repo.Users.Create(
		ctx,
		cmd.Username,
		cmd.Email,
		cmd.Name,
		cmd.Password,
		cmd.Role,
	)

	return toUserDTO(user), err
}
