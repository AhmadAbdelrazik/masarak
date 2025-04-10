package freelancerprofile

import (
	"context"
	"errors"
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
		email, name, title, pictureURL string,
		skills []string,
		yearsOfExperience int,
		hourlyRateAmount int,
		hourlyRateCurrency string,
		resumeURL string,
	) (*FreelancerProfile, error)

	// GetByEmail - Returns freelancer profile by email. returns
	// ErrProfileNotFound if it doesn't exist
	GetByEmail(ctx context.Context, email string) (*FreelancerProfile, error)

	// Update - Returns a freelancer profile by email for editing. the
	// fetched profile would be available in the updateFn for updating it
	// based on the domain logic. after that it will be saved.
	Update(
		ctx context.Context,
		email string,
		updateFn func(ctx context.Context, profile *FreelancerProfile) error,
	) error
}
