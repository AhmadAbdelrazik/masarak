package business

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	// Create - Creates a new business with unique name, email, and imageURL
	Create(ctx context.Context, name, email, description, imageURL string) (*Business, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Business, error)
	GetByName(ctx context.Context, name string) (*Business, error)
	// Save - Save the changes applied to business.
	// check that name, email, and imageURL still unique
	Save(ctx context.Context, business *Business) error
}
