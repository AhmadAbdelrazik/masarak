package app

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
	"github.com/ahmadabdelrazik/masarak/pkg/filters"
)

type GetJob struct {
	User       *authuser.User
	BusinessID int
	JobID      int
}

func (q *Queries) GetJobHandler(ctx context.Context, cmd GetJob) (Job, error) {
	business, err := q.repo.Businesses.GetByID(ctx, cmd.BusinessID)
	if err != nil {
		return Job{}, err
	}

	job, err := business.Job(cmd.JobID)
	if err != nil {
		return Job{}, err
	}

	return toJob(job), nil
}

type SearchJobs struct {
	User              *authuser.User
	Title             string
	WorkLocation      string
	WorkTime          string
	skills            []string
	YearsOfExperience int // default value is -1 and not 0
	Filters           filters.Filter
}

// SearchJobs Handler fetches the jobs that match the filters, use the nil
// value for each category you don't want to filter by it, except
// yearsOfExeprience use -1 instead of 0
func (q *Queries) SearchJobsHandler(ctx context.Context, cmd SearchJobs) ([]Job, filters.Metadata, error) {
	jobs, meta, err := q.repo.Businesses.SearchJobs(
		ctx,
		cmd.Title,
		cmd.WorkLocation,
		cmd.WorkTime,
		cmd.skills,
		cmd.YearsOfExperience,
		cmd.Filters,
	)
	if err != nil {
		return nil, filters.Metadata{}, err
	}

	jobDTOs := make([]Job, 0, len(jobs))

	for _, job := range jobs {
		jobDTOs = append(jobDTOs, toJob(job))
	}

	return jobDTOs, meta, nil
}
