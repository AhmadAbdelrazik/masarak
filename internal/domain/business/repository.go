package business

import (
	"context"
	"errors"

	"github.com/ahmadabdelrazik/masarak/pkg/filters"
)

var (
	ErrDuplicateBusiness = errors.New("duplicate business")
	ErrBusinessNotFound  = errors.New("business not found")
)

type Repository interface {
	// Create Creates a new business with unique name,
	// businessEmail, and imageURL
	Create(ctx context.Context, name, businessEmail, ownerEmail,
		description, imageURL string) (*Business, error)

	GetByID(ctx context.Context, id int) (*Business, error)
	GetByIDs(ctx context.Context, ids []int) ([]*Business,
		error)

	// Search  Returns all businesses matching the filters, the name
	// argument is used for filtering, for all names use "" in the
	// name argument
	Search(ctx context.Context, name string, filter filters.Filter) ([]*Business, filters.Metadata, error)

	// SearchJobs fetches the jobs that match the filters, use the nil
	// value for each category you don't want to filter by it, except
	// yearsOfExeprience use -1 instead of 0
	SearchJobs(
		ctx context.Context,
		title, workLocation, workTime string,
		skills []string,
		yearsOfExperience int,
		filter filters.Filter,
	) ([]*Job, filters.Metadata, error)

	// Update fetches the business and apply update to the business. the
	// updateFn provides the fetched business object where domain and
	// application logic can be applied. after that the object is saved to
	// the database
	Update(
		ctx context.Context,
		businessID int,
		updateFn func(ctx context.Context, business *Business) error,
	) (Business, error)
}
