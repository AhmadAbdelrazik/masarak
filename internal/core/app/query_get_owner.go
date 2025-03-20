package app

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/company"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/owner"
)

type GetOwner struct {
	Email string
}

type GetOwnerHandler struct {
	ownerRepo   owner.Repository
	companyRepo company.Repository
}

func NewGetOwnerHandler(ownerRepo owner.Repository, companyRepo company.Repository) *GetOwnerHandler {
	if ownerRepo == nil {
		panic("owner repo not found")
	} else if companyRepo == nil {
		panic("company repo not found")
	}

	return &GetOwnerHandler{
		ownerRepo:   ownerRepo,
		companyRepo: companyRepo,
	}
}

func (h *GetOwnerHandler) Handle(ctx context.Context, cmd GetOwner) (*owner.Owner, []*company.Company, error) {
	owner, err := h.ownerRepo.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return nil, nil, err
	}

	companies, err := h.companyRepo.GetByOwnerID(ctx, owner.ID())
	if err != nil {
		return nil, nil, err
	}

	return owner, companies, nil
}
