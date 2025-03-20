package memory

import (
	"context"
	"slices"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/job"
	"github.com/google/uuid"
)

type InMemoryJobRepository struct {
	memory *Memory
}

func NewInMemoryJobRepository(memory *Memory) *InMemoryJobRepository {
	return &InMemoryJobRepository{
		memory: memory,
	}
}

func (r *InMemoryJobRepository) Get(ctx context.Context, uid uuid.UUID) (*job.Job, error) {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, j := range r.memory.jobs {
		if j.ID() == uid {
			return &j, nil
		}
	}

	return nil, job.ErrJobNotFound
}

func (r *InMemoryJobRepository) Create(ctx context.Context, j *job.Job) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for _, jj := range r.memory.jobs {
		if jj.ID() == j.ID() {
			return job.ErrJobAlreadyExists
		}
	}

	r.memory.jobs = append(r.memory.jobs, *j)
	return nil
}
func (r *InMemoryJobRepository) Update(ctx context.Context, j *job.Job) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for i, jj := range r.memory.jobs {
		if jj.ID() == j.ID() {
			r.memory.jobs[i] = *j
			return nil
		}
	}

	return job.ErrJobNotFound
}
func (r *InMemoryJobRepository) Delete(ctx context.Context, uid uuid.UUID) error {
	r.memory.Lock()
	defer r.memory.Unlock()

	for i, jj := range r.memory.jobs {
		if jj.ID() == uid {
			r.memory.jobs = slices.Delete(r.memory.jobs, i, i+1)
			return nil
		}
	}

	return job.ErrJobNotFound
}
