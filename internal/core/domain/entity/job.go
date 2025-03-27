package entity

import (
	"errors"
	"time"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/valueobject"
	"github.com/google/uuid"
)

type Job struct {
	ID                uuid.UUID
	Status            *valueobject.JobStatus
	Title             string
	Skills            []string
	Description       string
	YearsOfExperience string
	WorkLocation      string
	WorkTime          string
	ExpectedSalary    string
	PostDate          time.Time
}

var ErrSkillLimitReached = errors.New("skill number must not be more than 10")

func NewJob(title, description, yearsOfExperience, workLocation, workTime, expectedSalary string, skills []string) (*Job, error) {
	if len(skills) > 10 {
		return nil, ErrSkillLimitReached
	}

	return &Job{
		ID:                uuid.New(),
		Status:            valueobject.NewJobStatus("available"),
		Title:             title,
		Description:       description,
		YearsOfExperience: yearsOfExperience,
		WorkLocation:      workLocation,
		WorkTime:          workTime,
		ExpectedSalary:    expectedSalary,
		PostDate:          time.Now(),
		Skills:            skills,
	}, nil
}

func (j *Job) Update(
	title, description, yearsOfExperience, workLocation, workTime, expectedSalary string,
	skills []string,
) error {
	if len(skills) > 10 {
		return ErrSkillLimitReached
	}

	j.Skills = skills
	j.Title = title
	j.Description = description
	j.YearsOfExperience = yearsOfExperience
	j.WorkLocation = workLocation
	j.WorkTime = workTime
	j.ExpectedSalary = expectedSalary
	j.Skills = skills

	return nil
}

// SetAvailable - Set the status of the job to available
func (j *Job) SetAvailable() {
	j.setStatus("available")
}

func (j *Job) SetCompleted() {
	j.setStatus("completed")
}

func (j *Job) SetPending() {
	j.setStatus("pending")
}

func (j *Job) setStatus(status string) {
	if j.Status.Status() == status {
		return
	}

	j.Status = valueobject.NewJobStatus(status)
}
