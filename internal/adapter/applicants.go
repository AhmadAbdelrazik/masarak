package adapter

import (
	"context"
	"slices"
	"sync"

	"github.com/ahmadabdelrazik/linkedout/internal/domain/applicant"
)

type InMemoryApplicantRepo struct {
	applicants []*applicant.Applicant

	sync.Mutex
}

func NewInMemoryApplicantRepo() *InMemoryApplicantRepo {
	return &InMemoryApplicantRepo{
		applicants: make([]*applicant.Applicant, 0),
	}
}

func (r *InMemoryApplicantRepo) Add(_ context.Context, a *applicant.Applicant) error {
	r.Lock()
	defer r.Unlock()

	for _, aa := range r.applicants {
		if aa.GetEmail() == a.GetEmail() {
			return applicant.ErrApplicantExists
		}
	}

	r.applicants = append(r.applicants, a)

	return nil
}

func (r *InMemoryApplicantRepo) GetByEmail(_ context.Context, email string) (*applicant.Applicant, error) {
	r.Lock()
	defer r.Unlock()

	for _, aa := range r.applicants {
		if aa.GetEmail() == email {
			return aa, nil
		}
	}

	return nil, applicant.ErrApplicantNotFound
}

func (r *InMemoryApplicantRepo) GetApplicantNumbers(_ context.Context) (int, error) {
	r.Lock()
	defer r.Unlock()

	return len(r.applicants), nil
}

func (r *InMemoryApplicantRepo) GetAllApplicants(_ context.Context) ([]*applicant.Applicant, error) {
	r.Lock()
	defer r.Unlock()

	applicants := slices.Clone(r.applicants)

	return applicants, nil

}
