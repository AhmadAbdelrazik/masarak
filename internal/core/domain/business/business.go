package business

import (
	"errors"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/job"
	"github.com/google/uuid"
)

type Business struct {
	id          uuid.UUID
	name        string
	email       string
	description string
	imageURL    string
	jobs        []*job.AvailableJob
}

func NewBusiness(name, email, description, imageURL string) (*Business, error) {
	business := &Business{
		name:        name,
		email:       email,
		description: description,
		imageURL:    imageURL,
	}

	return business, nil
}

func (b Business) ID() uuid.UUID {
	return b.id
}

func (b Business) Name() string {
	return b.name
}

func (b Business) Email() string {
	return b.email
}

var ErrJobAlreadyPosted = errors.New("job already posted")

func (b *Business) AddJob(job *job.AvailableJob) error {
	for _, j := range b.jobs {
		if j.ID() == job.ID() {
			return ErrJobAlreadyPosted
		}
	}

	return nil
}

func (b *Business) GetJobs() []*job.AvailableJob {
	return b.jobs
}
