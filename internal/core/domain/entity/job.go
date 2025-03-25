package entity

import (
	"time"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/valueobject"
	"github.com/google/uuid"
)

type Job struct {
	ID                uuid.UUID
	JobStatus         *valueobject.JobStatus
	Title             string
	Description       string
	YearsOfExperience string
	WorkLocation      string
	WorkTime          string
	ExpectedSalary    string
	PostDate          time.Time
}

func NewJob(title, description, yearsOfExperience, workLocation, workTime, expectedSalary string) (*Job, error) {
	return &Job{
		ID:                uuid.New(),
		JobStatus:         valueobject.NewJobStatus("available"),
		Title:             title,
		Description:       description,
		YearsOfExperience: yearsOfExperience,
		WorkLocation:      workLocation,
		WorkTime:          workTime,
		ExpectedSalary:    expectedSalary,
		PostDate:          time.Now(),
	}, nil
}
