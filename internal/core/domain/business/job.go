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

	for _, jj := range b.jobs {
		if jj.Title == job.Title {
			return ErrDuplicateJob
		}
	}

	b.jobs = append(b.jobs, job)

	return nil
}

// UpdateJob - Updates the details of a job
func (b *Business) UpdateJob(
	jobID uuid.UUID,
	title, description, yearsOfExperience, workLocation, workTime, expectedSalary string,
	skills []string,
) error {
	for _, job := range b.jobs {
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

func (b *Business) MarkJobClosed(jobID uuid.UUID) error {
	for _, j := range b.jobs {
		if j.ID == jobID {
			j.SetStatusToClosed()
			return nil
		}
	}

	return ErrJobNotFound
}

func (b *Business) MarkJobOpen(jobID uuid.UUID) error {
	for _, j := range b.jobs {
		if j.ID == jobID {
			j.SetStatusToOpen()
			return nil
		}
	}

	return ErrJobNotFound
}
func (b *Business) MarkJobArchived(jobID uuid.UUID) error {
	for _, j := range b.jobs {
		if j.ID == jobID {
			j.SetStatusToArchived()
			return nil
		}
	}

	return ErrJobNotFound
}

func (b *Business) GetJobByName(title string) (job.Job, error) {
	for _, j := range b.jobs {
		if j.Title == title {
			return *j, nil
		}
	}

	return job.Job{}, ErrJobNotFound
}

func (b *Business) GetJobByID(id uuid.UUID) (job.Job, error) {
	for _, j := range b.jobs {
		if j.ID == id {
			return *j, nil
		}
	}

	return job.Job{}, ErrJobNotFound
}

// GetAllJobs - Returns all jobs posted by the business. For only
// the open jobs use GetAllOpenJobs instead.
func (b *Business) GetAllJobs() []job.Job {
	jobs := make([]job.Job, 0, len(b.jobs))

	for _, jj := range b.jobs {
		jobs = append(jobs, *jj)
	}

	return jobs
}

func (b *Business) GetAllOpenJobs() []job.Job {
	return b.getAllJobsByStatus("open")
}

func (b *Business) GetAllClosedJobs() []job.Job {
	return b.getAllJobsByStatus("closed")
}

func (b *Business) GetAllArchivedJobs() []job.Job {
	return b.getAllJobsByStatus("archived")
}

// getAllJobsByStatus - returns list of jobs by status ("open", "closed", "archived")
func (b *Business) getAllJobsByStatus(status string) []job.Job {
	jobs := make([]job.Job, 0, len(b.jobs))

	for _, jj := range b.jobs {
		if jj.Status() == status {
			jobs = append(jobs, *jj)
		}
	}

	return jobs
}
