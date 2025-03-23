package app

import (
	"context"
)

type GetOwner struct {
	Email string
}

func (h *Queries) GetOwnerHandler(ctx context.Context, cmd GetOwner) (*OwnerDTO, error) {
	owner, err := h.repo.Owner.GetByEmail(ctx, cmd.Email)
	if err != nil {
		return nil, err
	}

	companies, err := h.repo.Companies.GetByOwnerID(ctx, owner.ID())
	if err != nil {
		return nil, err
	}

	ownerDTO := NewOwnerDTO(owner, companies)
	return &ownerDTO, nil
}
