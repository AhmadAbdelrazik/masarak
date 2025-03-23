package app

import (
	"context"
	"errors"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/owner"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/talent"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/valueobject"
)

type RegisterUserType struct {
	User *authuser.AuthUser
	Role string
}

var (
	ErrUserAlreadyRegistered = errors.New("user already registered")
)

func (h *Commands) RegisterUserTypeHandler(ctx context.Context, cmd RegisterUserType) error {
	if !cmd.User.Role.Is("user") {
		return ErrUserAlreadyRegistered
	}

	role, err := valueobject.NewRole(cmd.Role)
	if err != nil {
		// ErrInvalidRole
		return err
	}

	switch {
	case role.Is("owner"):
		o, err := owner.New(cmd.User)
		if err != nil {
			return err
		}
		if err := h.repo.Owner.Create(ctx, o); err != nil {
			return err
		}
	case role.Is("talent"):
		t, err := talent.New(cmd.User)
		if err != nil {
			return err
		}
		if err := h.repo.Talents.Create(ctx, t); err != nil {
			return err
		}
	default:
		return valueobject.ErrInvalidRole
	}

	if err := h.repo.AuthUsers.ChangeRole(ctx, cmd.User.Email, role); err != nil {
		return err
	}

	return nil
}
