package freelancerprofile

import (
	"context"

	"github.com/Rhymond/go-money"
)

type Repository interface {
	// Create - Creates a new freelancer Profile with unique email
	Create(
		ctx context.Context,
		name, email, pictureURL, title string,
		skills []string,
		yearsOfExperience int,
		hourlyRate *money.Money,
	) error
	GetByEmail(ctx context.Context, email string) *FreelancerProfile
	Save(ctx context.Context, applicationHistory *FreelancerProfile) error
}
