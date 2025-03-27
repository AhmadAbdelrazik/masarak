package business

import (
	"errors"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity/job"
	"github.com/google/uuid"
)

var (
	ErrDuplicateJob = errors.New("job with the same title already exists")
	ErrJobNotFound  = errors.New("job not found")
)

// CreateNewJob - Creates a new job in the business. The job must not have the same title as the other jobs.
func (b *Business) CreateNewJob(
	title, description, yearsOfExperience, workLocation, workTime, expectedSalary string,
	skills []string,
) error {
	job, err := job.NewJob(
		title,
		description,
		yearsOfExperience,
		workLocation,
		workTime,
		expectedSalary,
		skills,
	)
	if err != nil {
		return err
	}

	for _, jj := range b.availableJobs {
		if jj.Title == job.Title {
			return ErrDuplicateJob
		}
	}

	b.availableJobs = append(b.availableJobs, job)

	return nil
}

// GetAllAvailableJobs - Returns jobs that are still available
func (b *Business) GetAllAvailableJobs() []job.Job {
	jobs := make([]job.Job, 0, len(b.availableJobs))

	for _, jj := range b.availableJobs {
		if jj.IsAvailable() {
			jobs = append(jobs, *jj)
		}
	}

	return jobs
}

// GetJobByName - Query the job in available jobs
func (b *Business) GetJobByName(title string) (job.Job, error) {
	for _, j := range b.availableJobs {
		if j.Title == title {
			return *j, nil
		}
	}

	return job.Job{}, ErrJobNotFound
}

// GetJobByID - Query the job in available jobs
func (b *Business) GetJobByID(id uuid.UUID) (job.Job, error) {
	for _, j := range b.availableJobs {
		if j.ID == id {
			return *j, nil
		}
	}

	return job.Job{}, ErrJobNotFound
}

// UpdateJob - Updates the details of a job
func (b *Business) UpdateJob(
	jobID uuid.UUID,
	title, description, yearsOfExperience, workLocation, workTime, expectedSalary string,
	skills []string,
) error {
	for _, job := range b.availableJobs {
		if job.ID == jobID {
			return job.Update(
				title,
				description,
				yearsOfExperience,
				workLocation,
				workTime,
				expectedSalary,
				skills,
			)
		}
	}

	return ErrJobNotFound
}
