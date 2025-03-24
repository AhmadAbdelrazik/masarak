package entity

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ID                uuid.UUID
	Title             string
	Description       string
	YearsOfExperience string
	WorkLocation      string
	ExpectedSalary    string
	PostDate          time.Time
}

func NewJob(title, description, yearsOfExperience, workLocation, expectedSalary string) (*Job, error) {
	return &Job{
		ID:                uuid.New(),
		Title:             title,
		Description:       description,
		YearsOfExperience: yearsOfExperience,
		WorkLocation:      workLocation,
		ExpectedSalary:    expectedSalary,
		PostDate:          time.Now(),
	}, nil
}
