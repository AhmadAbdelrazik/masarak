package app

import (
	"context"

	"github.com/google/uuid"
)

type GetOwners struct {
	Limit  int
	Offset int
}

func (h *Queries) GetOwnersHandler(ctx context.Context, cmd GetOwners) ([]*OwnerDTO, error) {
	owners, err := h.repo.Owner.GetAll(ctx, cmd.Limit, cmd.Offset)
	if err != nil {
		return nil, err
	}

	ids := make([]uuid.UUID, 0, len(owners))

	for _, o := range owners {
		ids = append(ids, o.ID())
	}

}
