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
	// Create - Creates a new freelancer Profile with unique email
	Create(
		ctx context.Context,
		name, email, pictureURL, title string,
		skills []string,
		yearsOfExperience int,
		hourlyRate *money.Money,
	) (*FreelancerProfile, error)
	GetByEmail(ctx context.Context, email string) (*FreelancerProfile, error)
	Save(ctx context.Context, profile *FreelancerProfile) error
}
