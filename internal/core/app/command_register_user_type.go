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

type RegisterUserTypeHandler struct {
	ownerRepo  owner.Repository
	talentRepo talent.Repository
	userRepo   authuser.Repository
}

func NewRegisterUserTypeHandler(ownerRepo owner.Repository, userRepo authuser.Repository, talentRepo talent.Repository) *RegisterUserTypeHandler {
	if ownerRepo == nil {
		panic("owner repo not found")
	}
	if talentRepo == nil {
		panic("talent repo not found")
	}

	if userRepo == nil {
		panic("user repo not found")
	}

	return &RegisterUserTypeHandler{
		ownerRepo:  ownerRepo,
		userRepo:   userRepo,
		talentRepo: talentRepo,
	}
}

var (
	ErrUserAlreadyRegistered = errors.New("user already registered")
)

func (h *RegisterUserTypeHandler) Handle(ctx context.Context, cmd RegisterUserType) error {
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
		if err := h.ownerRepo.Create(ctx, o); err != nil {
			return err
		}
	case role.Is("talent"):
		t, err := talent.New(cmd.User)
		if err != nil {
			return err
		}
		if err := h.talentRepo.Create(ctx, t); err != nil {
			return err
		}
	default:
		return valueobject.ErrInvalidRole
	}

	if err := h.userRepo.ChangeRole(ctx, cmd.User.Email, role); err != nil {
		return err
	}

	return nil
}
