package app

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/owner"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/valueobject"
)

type RegisterOwner struct {
	User *authuser.AuthUser
}

type RegisterOwnerHandler struct {
	ownerRepo owner.Repository
	userRepo  authuser.Repository
}

func NewRegisterOwnerHandler(ownerRepo owner.Repository, userRepo authuser.Repository) *RegisterOwnerHandler {
	if ownerRepo == nil {
		panic("owner repo not found")
	}

	if userRepo == nil {
		panic("user repo not found")
	}
	return &RegisterOwnerHandler{
		ownerRepo: ownerRepo,
		userRepo:  userRepo,
	}
}

func (h *RegisterOwnerHandler) Handle(ctx context.Context, cmd RegisterOwner) error {
	o, err := owner.New(cmd.User)
	if err != nil {
		return err
	}

	if err := h.ownerRepo.Create(ctx, o); err != nil {
		return err
	}

	ownerRole, err := valueobject.NewRole("owner")
	if err != nil {
		return err
	}
	if err := h.userRepo.ChangeRole(ctx, o.Email(), ownerRole); err != nil {
		return err
	}

	return nil
}
