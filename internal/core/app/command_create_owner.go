package app

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/owner"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/valueobject"
)

type CreateOwner struct {
	Name  string
	Email string
}

type CreateOwnerHandler struct {
	ownerRepo owner.Repository
	userRepo  authuser.Repository
}

func NewCreateOwnerHandler(ownerRepo owner.Repository, userRepo authuser.Repository) *CreateOwnerHandler {
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
	o, err := owner.New(cmd.Name, cmd.Email)
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
	if err := h.userRepo.ChangeRole(ctx, cmd.Email, ownerRole); err != nil {
		return err
	}

	return nil
}
