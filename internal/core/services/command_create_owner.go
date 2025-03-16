package app

import (
	"context"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/owner"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/valueobject"
)

type CreateOwner struct {
	user entity.AuthUser
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
	o, err := owner.New(cmd.user.Name, cmd.user.Email)
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
	if err := h.userRepo.ChangeRole(ctx, cmd.user.Email, ownerRole); err != nil {
		return err
	}

	return nil
}
