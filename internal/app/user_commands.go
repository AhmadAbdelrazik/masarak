package app

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

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

type SelectRole struct {
	User *authuser.User
	ID   int
	Role string
}

// SelectRoleHandler - Selects a role for a user. Notice that this action is
// available only once per account. returns ErrUserNotFound or ErrEditConflict
// or ErrInvalidRole in case of errors
func (c *Commands) SelectRoleHandler(ctx context.Context, cmd SelectRole) error {
	if !cmd.User.HasPermission("role.select") {
		return ErrUnauthorized
	}

	err := c.repo.Users.Update(ctx, cmd.ID, func(ctx context.Context, user *authuser.User) error {
		err := user.UpdateRole(cmd.Role)

		return err
	})

	return err
}
