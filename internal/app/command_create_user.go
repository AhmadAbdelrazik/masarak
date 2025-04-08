package app

import "context"

type CreateUser struct {
}

func (c *Commands) CreateUserHandler(ctx context.Context, cmd CreateUser) error {
	c.repo.AuthUsers.Create()
}
