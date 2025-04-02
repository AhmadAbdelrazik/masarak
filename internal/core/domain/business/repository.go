package business

import (
	"context"
	"errors"

	"github.com/ahmadabdelrazik/masarak/pkg/filters"
	"github.com/google/uuid"
)

var (
	ErrDuplicateBusiness = errors.New("duplicate business")
	ErrBusinessNotFound  = errors.New("business not found")
)

type Repository interface {
	// Create - Creates a new business with unique name,
	// businessEmail, and imageURL
	Create(ctx context.Context, name, businessEmail, ownerEmail,
		description, imageURL string) (*Business, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Business, error)
	GetByIDs(ctx context.Context, ids []uuid.UUID) ([]*Business,
		error)
	// GetAll - Returns all businesses matching the filters, the name
	// argument is used for filtering, for all names use "" in the
	// name argument
	GetAll(ctx context.Context, name string, filter filters.Filter) ([]*Business, error)
	// Save - Save the changes applied to business. check that name,
	// email, and imageURL still unique
	Save(ctx context.Context, business *Business) error
}
