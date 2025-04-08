package postgres

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/internal/domain/business"
	"github.com/ahmadabdelrazik/masarak/pkg/db"
)

type BusinessRepository struct {
	db *db.DB
}

func NewBusinessRepo(db *db.DB) *BusinessRepository {
	return &BusinessRepository{db}
}

func (r *BusinessRepository) Create(
	ctx context.Context,
	name, businessEmail, ownerEmail, description, imageURL string,
) (*business.Business, error) {
	business, err := business.New(name, businessEmail, ownerEmail, description, imageURL)
	if err != nil {
		return nil, err
	}

	query := `
INSERT INTO businesses(name, email, picture_url, description)
VALUES ($1,$2,$3, $4)
RETURNING id
	`

	var id int
	if err = r.db.QueryRowContext(
		ctx,
		query,
		name,
		businessEmail,
		imageURL,
		description,
	).Scan(&id); err != nil {
		return nil, err
	}

	return business, nil
}
