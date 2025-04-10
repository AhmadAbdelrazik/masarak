package app

import "context"

// GetFreelancerProfileHandler - returns a freelancer profile if found, returns
// ErrProfileNotFound if not found
func (q *Queries) GetFreelancerProfileHandler(ctx context.Context, email string) (FreelancerProfile, error) {
	profile, err := q.repo.FreelancerProfile.GetByEmail(ctx, email)
	if err != nil {
		return FreelancerProfile{}, err
	}

	return toFreelancerProfile(profile), nil
}
