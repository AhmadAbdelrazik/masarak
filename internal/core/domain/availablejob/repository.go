package job

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

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
