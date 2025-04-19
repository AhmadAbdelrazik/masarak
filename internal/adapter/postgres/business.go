package postgres

import (
	"context"
	"database/sql"

	"github.com/ahmadabdelrazik/masarak/internal/domain/business"
	"github.com/ahmadabdelrazik/masarak/pkg/filters"
)

type BusinessRepo struct {
	db *sql.DB
}

func (r *BusinessRepo) Create(
	ctx context.Context,
	name, businessEmail, ownerEmail,
	description, imageURL string,
) (*business.Business, error) {
	return nil, nil
}

func (r *BusinessRepo) GetByID(ctx context.Context, id int) (*business.Business, error) {
	return nil, nil
}

func (r *BusinessRepo) GetByIDs(ctx context.Context, ids []int) ([]*business.Business, error) {
	return nil, nil
}

func (r *BusinessRepo) Search(ctx context.Context, name string, filter filters.Filter) ([]*business.Business, filters.Metadata, error) {
	return nil, filters.Metadata{}, nil
}

func (r *BusinessRepo) SerachJobs(
	ctx context.Context,
	title, workLocation, workTime string,
	skills []string,
	yearsOfExperience int,
	filter filters.Filter,
) ([]*business.Job, filters.Metadata, error) {
	return nil, filters.Metadata{}, nil
}

func (r *BusinessRepo) Update(
	ctx context.Context,
	businessID int,
	updateFn func(ctx context.Context, business *business.Business) error,
) (business.Business, error) {
	return business.Business{}, nil
}
