package job

import (
	"context"
	"errors"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
	"github.com/google/uuid"
)

type Job struct {
	job       *entity.Job
	companyID uuid.UUID
}

func (j *Job) ID() uuid.UUID {
	return j.job.ID
}

func New(title, description string, companyID uuid.UUID) (*Job, error) {
	job, err := entity.NewJob(title, description)
	if err != nil {
		return nil, err
	}

	return &Job{
		job:       job,
		companyID: companyID,
	}, nil
}

var (
	ErrJobNotFound   = errors.New("job not found")
	ErrAlreadyExists = errors.New("job already exists")
)

type Repository interface {
	Get(ctx context.Context, uid uuid.UUID) (*Job, error)
	Create(ctx context.Context, job *Job) error
	Update(ctx context.Context, job *Job) error
	Delete(ctx context.Context, uid uuid.UUID) error
}
