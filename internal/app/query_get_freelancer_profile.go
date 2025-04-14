package app

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

type GetFreelancerProfile struct {
	User     *authuser.User
	Username string
}

// GetFreelancerProfileHandler - returns a freelancer profile if found, returns
// ErrProfileNotFound if not found
func (q *Queries) GetFreelancerProfileHandler(ctx context.Context, cmd GetFreelancerProfile) (FreelancerProfile, error) {
	if !cmd.User.HasPermission("freelancer_profile.read") {
		return FreelancerProfile{}, ErrUnauthorized
	}

	profile, err := q.repo.FreelancerProfile.GetByUsername(ctx, cmd.Username)
	if err != nil {
		return FreelancerProfile{}, err
	}

	return toFreelancerProfile(profile), nil
}
