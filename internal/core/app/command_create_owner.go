package app

import (
	"context"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/owner"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/valueobject"
)

type CreateOwner struct {
	User entity.AuthUser
}

type CreateOwnerHandler struct {
	ownerRepo owner.Repository
	userRepo  entity.AuthUserRepository
}

func NewCreateOwnerHandler(ownerRepo owner.Repository, userRepo entity.AuthUserRepository) *CreateOwnerHandler {
	if ownerRepo == nil {
		panic("owner repo not found")
	}

	if userRepo == nil {
		panic("user repo not found")
	}
	return &CreateOwnerHandler{
		ownerRepo: ownerRepo,
		userRepo:  userRepo,
	}
}

func (h *CreateOwnerHandler) Handle(ctx context.Context, cmd CreateOwner) error {
	o, err := owner.New(cmd.User.Name, cmd.User.Email)
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
	if err := h.userRepo.ChangeRole(ctx, cmd.User.Email, ownerRole); err != nil {
		return err
	}

	return nil
}
