package memory

import (
	"context"

	"github.com/Rhymond/go-money"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/freelancerprofile"
)

type InMemoryFreelancerProfileRepository struct {
	memory *Memory
}

func NewInMemoryFreelancerProfileRepository(mem *Memory) *InMemoryFreelancerProfileRepository {
	return &InMemoryFreelancerProfileRepository{
		memory: mem,
	}
}

func (r *InMemoryFreelancerProfileRepository) Create(
	ctx context.Context,
	name, email, pictureURL, title string,
	skills []string,
	yearsOfExperience int,
	hourlyRate *money.Money,
) (*freelancerprofile.FreelancerProfile, error) {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, profile := range r.memory.freelancerProfiles {
		if profile.Email() == email {
			return nil, freelancerprofile.ErrDuplicateProfile
		}
	}

	profile, err := freelancerprofile.New(
		name,
		email,
		pictureURL,
		title,
		skills,
		yearsOfExperience,
		hourlyRate,
	)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (r *InMemoryFreelancerProfileRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*freelancerprofile.FreelancerProfile, error) {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, profile := range r.memory.freelancerProfiles {
		if profile.Email() == email {
			return profile, nil
		}
	}

	return nil, freelancerprofile.ErrProfileNotFound
}

func (r *InMemoryFreelancerProfileRepository) Save(
	ctx context.Context,
	profile *freelancerprofile.FreelancerProfile,
) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for i, p := range r.memory.freelancerProfiles {
		if p.Email() == profile.Email() {
			r.memory.freelancerProfiles[i] = profile
			return nil
		}
	}

	r.memory.freelancerProfiles = append(r.memory.freelancerProfiles, profile)

	return nil
}
