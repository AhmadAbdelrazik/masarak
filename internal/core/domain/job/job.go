package job

import (
	"context"
	"errors"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity"
	"github.com/google/uuid"
)

type AvailableJob struct {
	job *entity.Job
}

func (j *AvailableJob) ID() uuid.UUID {
	return j.job.ID
}

func New(title, description, yearsOfExperience, workLocation, expectedSalary string) (*AvailableJob, error) {
	job, err := entity.NewJob(title, description, yearsOfExperience, workLocation, expectedSalary)
	if err != nil {
		return nil, err
	}

	return &AvailableJob{
		job:       job,
		companyID: companyID,
	}, nil
}

var (
	ErrJobNotFound      = errors.New("job not found")
	ErrJobAlreadyExists = errors.New("job already exists")
)

type Repository interface {
	Get(ctx context.Context, uid uuid.UUID) (*AvailableJob, error)
	Create(ctx context.Context, job *AvailableJob) error
	Update(ctx context.Context, job *AvailableJob) error
	Delete(ctx context.Context, uid uuid.UUID) error
}
