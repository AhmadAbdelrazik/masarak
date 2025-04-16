package app

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/pkg/filters"
)

type GetBusiness struct {
	BusinessID int
}

func (q *Queries) GetBusinessHandler(ctx context.Context, cmd GetBusiness) (Business, error) {
	business, err := q.repo.Businesses.GetByID(ctx, cmd.BusinessID)
	if err != nil {
		return Business{}, err
	}

	return toBusiness(business), nil
}

type SearchBusinesses struct {
	Name string
	// Filters has defaults of page number 1 and page size 20. the only
	// possible sort is by name.
	Filters filters.Filter
}

func (q *Queries) SearchBusinessesHandler(ctx context.Context, cmd SearchBusinesses) ([]Business, filters.Metadata, error) {
	businesses, metadata, err := q.repo.Businesses.Search(ctx, cmd.Name, cmd.Filters)
	if err != nil {
		return nil, filters.Metadata{}, err
	}

	businessDTO := make([]Business, 0, len(businesses))
	for _, b := range businesses {
		businessDTO = append(businessDTO, toBusiness(b))
	}

	return businessDTO, metadata, nil
}
