package job

import (
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/business"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity"
	"github.com/google/uuid"
)

type AvailableJob struct {
	job          *entity.Job
	applications []*entity.Application
	business     *business.Business
}

func (j *AvailableJob) ID() uuid.UUID {
	return j.job.ID
}

func New(
	business *business.Business,
	title, description, yearsOfExperience, workLocation, workTime, expectedSalary string,
) (*AvailableJob, error) {
	job, err := entity.NewJob(title, description, yearsOfExperience, workLocation, workTime, expectedSalary)
	if err != nil {
		return nil, err
	}

	return &AvailableJob{
		job:          job,
		applications: make([]*entity.Application, 0),
		business:     business,
	}, nil
}
