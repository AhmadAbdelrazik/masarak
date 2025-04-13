package app

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

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
