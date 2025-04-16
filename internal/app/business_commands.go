package app

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/internal/domain/business"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

type CreateBusiness struct {
	User          *authuser.User
	Name          string
	BusinessEmail string
	Description   string
	ImageURL      string
}

func (c *Commands) CreateBusinessHandler(ctx context.Context, cmd CreateBusiness) (Business, error) {
	if !cmd.User.HasPermission("business.create") {
		return Business{}, ErrUnauthorized
	}

	business, err := c.repo.Businesses.Create(
		ctx,
		cmd.Name,
		cmd.BusinessEmail,
		cmd.User.Email(),
		cmd.Description,
		cmd.ImageURL,
	)
	if err != nil {
		return Business{}, err
	}

	return toBusiness(business), nil
}

type UpdateBusiness struct {
	User          *authuser.User
	BusinessID    int
	Name          *string
	BusinessEmail *string
	Description   *string
	ImageURL      *string
}

func (c *Commands) UpdateBusinessHandler(ctx context.Context, cmd UpdateBusiness) (Business, error) {
	if !cmd.User.HasPermission("business.update") {
		return Business{}, ErrUnauthorized
	}

	business, err := c.repo.Businesses.Update(
		ctx,
		cmd.BusinessID,
		func(ctx context.Context, business *business.Business) error {
			if !business.IsEmployee(cmd.User.Email()) {
				return ErrUnauthorized
			}

			if cmd.Name != nil {
				if err := business.UpdateName(*cmd.Name); err != nil {
					return err
				}
			}

			if cmd.Description != nil {
				if err := business.UpdateDescription(*cmd.Description); err != nil {
					return err
				}
			}

			if cmd.BusinessEmail != nil {
				if err := business.UpdateBusinessEmail(*cmd.BusinessEmail); err != nil {
					return err
				}
			}

			if cmd.ImageURL != nil {
				business.ImageURL = *cmd.ImageURL
			}

			return nil
		},
	)

	if err != nil {
		return Business{}, err
	}

	return toBusiness(&business), nil
}

type AddEmployeeToBusiness struct {
	User          authuser.User
	BusinessID    int
	EmployeeEmail string
}

func (c *Commands) AddEmployeeToBusinessHandler(ctx context.Context, cmd AddEmployeeToBusiness) error {
	if !cmd.User.HasPermission("business.add_employee") {
		return ErrUnauthorized
	}

	_, err := c.repo.Businesses.Update(
		ctx,
		cmd.BusinessID,
		func(ctx context.Context, business *business.Business) error {
			if !business.IsEmployee(cmd.User.Email()) {
				return ErrUnauthorized
			}

			return business.AddEmployee(cmd.EmployeeEmail)
		},
	)

	return err
}

type RemoveEmployeeToBusiness struct {
	User          authuser.User
	BusinessID    int
	EmployeeEmail string
}

func (c *Commands) RemoveEmployeeFromBusinessHandler(ctx context.Context, cmd RemoveEmployeeToBusiness) error {
	if !cmd.User.HasPermission("business.delete_employee") {
		return ErrUnauthorized
	}

	_, err := c.repo.Businesses.Update(
		ctx,
		cmd.BusinessID,
		func(ctx context.Context, business *business.Business) error {
			if !business.IsEmployee(cmd.User.Email()) {
				return ErrUnauthorized
			}

			return business.RemoveEmployee(cmd.EmployeeEmail)
		},
	)

	return err
}
