package freelancerprofile

import (
	"context"
	"errors"

	"github.com/Rhymond/go-money"
)

var (
	ErrDuplicateProfile = errors.New("profile already exists")
	ErrProfileNotFound  = errors.New("profile doesn't exist")
)

type Repository interface {
	// Create - Creates a new freelancer Profile with unique email and save
	// it in the database. returns ErrDuplicateProfile if it already exists
	Create(
		ctx context.Context,
		name, email, pictureURL, title string,
		skills []string,
		yearsOfExperience int,
		hourlyRate *money.Money,
	) (*FreelancerProfile, error)

	// GetByEmail - Returns freelancer profile by email. returns
	// ErrProfileNotFound if it doesn't exist
	GetByEmail(ctx context.Context, email string) (*FreelancerProfile, error)

	// Save - Returns a freelancer profile by email for editing. the
	// fetched profile would be available in the updateFn for updating it
	// based on the domain logic. after that it will be saved.
	Save(
		ctx context.Context,
		email string,
		updateFn func(ctx context.Context, profile *FreelancerProfile) error,
	) error
}
