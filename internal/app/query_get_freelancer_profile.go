package app

import (
	"context"
)

type GetFreelancerProfile struct {
	Username string
}

// GetFreelancerProfileHandler - returns a freelancer profile if found, returns
// ErrProfileNotFound if not found
func (q *Queries) GetFreelancerProfileHandler(ctx context.Context, cmd GetFreelancerProfile) (FreelancerProfile, error) {
	profile, err := q.repo.FreelancerProfile.GetByUsername(ctx, cmd.Username)
	if err != nil {
		return FreelancerProfile{}, err
	}

	return toFreelancerProfile(profile), nil
}
