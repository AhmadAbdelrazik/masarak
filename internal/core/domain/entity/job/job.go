package job

import (
	"errors"
	"time"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/valueobject"
	"github.com/google/uuid"
)

type Job struct {
	ID                uuid.UUID
	status            *valueobject.JobStatus
	Title             string
	Skills            []string
	Description       string
	YearsOfExperience string
	WorkLocation      string
	WorkTime          string
	ExpectedSalary    string
	PostDate          time.Time
	applications      []*entity.Application
}

var ErrSkillLimitReached = errors.New("skill number must not be more than 10")

func NewJob(title, description, yearsOfExperience, workLocation, workTime, expectedSalary string, skills []string) (*Job, error) {
	if len(skills) > 10 {
		return nil, ErrSkillLimitReached
	}

	return &Job{
		ID:                uuid.New(),
		status:            valueobject.NewJobStatus("available"),
		Title:             title,
		Description:       description,
		YearsOfExperience: yearsOfExperience,
		WorkLocation:      workLocation,
		WorkTime:          workTime,
		ExpectedSalary:    expectedSalary,
		PostDate:          time.Now(),
		Skills:            skills,
		applications:      make([]*entity.Application, 0),
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

func (j *Job) SetStatusToOpen() {
	j.setStatus("available")
}

func (j *Job) SetStatusToClosed() {
	j.setStatus("closed")
}

func (j *Job) SetStatusToArchived() {
	j.setStatus("archived")
}

func (j *Job) setStatus(status string) {
	if j.status.Status() == status {
		return
	}

	newStatus, err := valueobject.NewJobStatus(status)
	if err != nil {
		panic(err)
	}

	j.status = newStatus
}

func (j *Job) Status() string {
	return j.status.Status()
}

func (j *Job) IsAvailable() bool {
	return j.status.IsOpen()
}
