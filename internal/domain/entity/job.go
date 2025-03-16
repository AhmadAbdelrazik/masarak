package entity

import "github.com/ahmadabdelrazik/linkedout/internal/domain/valueobjects"

type Job struct {
	Title       string
	Category    string
	SalaryRange *valueobjects.SalaryRange
	Details     string
}

type JobOption func(*Job) error

func NewJob(title, category, details string, opts ...JobOption) (*Job, error) {
	job := &Job{
		Title:    title,
		Category: category,
		Details:  details,
	}

	for _, opt := range opts {
		if err := opt(job); err != nil {
			return nil, err
		}
	}

	return job, nil
}

func WithSalaryRange(salaryRange *valueobjects.SalaryRange) JobOption {
	return func(j *Job) error {
		j.SalaryRange = salaryRange
		return nil
	}
}
