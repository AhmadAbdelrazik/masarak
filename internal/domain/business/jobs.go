package business

import (
	"fmt"
	"slices"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/ahmadabdelrazik/masarak/internal/domain"
	"github.com/ahmadabdelrazik/masarak/internal/domain/valueobject"
)

type Range[T comparable] struct {
	From T
	To   T
}

type SalaryRange Range[*money.Money]
type YearsOfExperienceRange Range[int]

type Job struct {
	businessID       int
	businessName     string
	businessImageURL string

	id           int
	title        string
	description  string
	workLocation string // office, hybrid, remote
	workTime     string // full time or part time
	skills       []string

	yearsOfExperience YearsOfExperienceRange
	expectedSalary    SalaryRange

	status       *valueobject.JobStatus
	applications []Application
	createdAt    time.Time
	updatedAt    time.Time
}

var (
	ErrInvalidJobProperty = fmt.Errorf("%w: Job", domain.ErrInvalidProperty)
	ErrInvalidJobUpdate   = fmt.Errorf("%w: Job", domain.ErrInvalidUpdate)
	ErrJobConflict        = fmt.Errorf("%w: Job conflict", domain.ErrInvalidUpdate)
)

func newJob(
	businessID int,
	businessName, businessImage string,
	title, description, workLocation, workTime string,
	skills []string,
) (*Job, error) {
	if len(title) == 0 || len(title) > 30 {
		return nil, fmt.Errorf("%w: job title must be between 0 and 30 bytes", ErrInvalidJobProperty)

	}

	if len(description) == 0 || len(description) > 1000 {
		return nil, fmt.Errorf("%w: description must be between 0 and 1000 bytes", ErrInvalidJobProperty)
	}

	if !slices.Contains([]string{"remote", "office", "hybrid"}, workLocation) {
		return nil, fmt.Errorf("%w: incorrect work location. (remote - office - hybrid)", ErrInvalidJobProperty)
	}
	if !slices.Contains([]string{"full time", "part time"}, workTime) {
		return nil, fmt.Errorf("%w: incorrect work time. (full time - part time)", ErrInvalidJobProperty)
	}

	if len(skills) == 0 || len(skills) > 10 {
		return nil, fmt.Errorf("%w: invalid number of skills (must be between 1 and 10)", ErrInvalidJobProperty)
	}

	status, err := valueobject.NewJobStatus("open")
	if err != nil {
		panic(err)
	}

	from, to := money.New(0, "EGP"), money.New(999_999_999, "EGP")

	return &Job{
		businessID:        businessID,
		businessName:      businessName,
		businessImageURL:  businessImage,
		title:             title,
		description:       description,
		workLocation:      workLocation,
		workTime:          workTime,
		skills:            skills,
		yearsOfExperience: YearsOfExperienceRange{From: 0, To: 50},
		expectedSalary:    SalaryRange{From: from, To: to},
		status:            status,
		applications:      []Application{},
		createdAt:         time.Now(),
		updatedAt:         time.Now(),
	}, nil
}

// InstantiateJob instantiate a job from the database.
func InstantiateJob(
	businessID int,
	businessName, businessImage string,
	id int,
	title, description, workLocation, workTime string,
	skills []string,
	yearsOfExperience YearsOfExperienceRange,
	salary SalaryRange,
	status string,
	applications []Application,
	createdAt, updatedAt time.Time,
) *Job {
	jobStatus, err := valueobject.NewJobStatus(status)
	if err != nil {
		panic(err)
	}

	return &Job{
		businessID:        businessID,
		businessName:      businessName,
		businessImageURL:  businessImage,
		id:                id,
		title:             title,
		description:       description,
		workLocation:      workLocation,
		workTime:          workTime,
		skills:            skills,
		yearsOfExperience: yearsOfExperience,
		expectedSalary:    salary,
		status:            jobStatus,
		applications:      applications,
		createdAt:         createdAt,
		updatedAt:         updatedAt,
	}
}

func (j *Job) BusinessID() int {
	return j.id
}

func (j *Job) BusinessName() string {
	return j.businessName
}

func (j *Job) BusinessImageURL() string {
	return j.businessImageURL
}

func (j *Job) ID() int {
	return j.id
}

func (j *Job) Title() string {
	return j.title
}

func (j *Job) updateTitle(title string) error {
	if len(title) == 0 || len(title) > 30 {
		return fmt.Errorf("%w: job title must be between 0 and 30 bytes", ErrInvalidJobUpdate)
	}

	j.title = title

	j.updatedAt = time.Now()

	return nil
}

func (j *Job) Description() string {
	return j.description
}

func (j *Job) UpdateDescription(description string) error {
	if len(description) == 0 || len(description) > 1000 {
		return fmt.Errorf("%w: description must be between 0 and 1000 bytes", ErrInvalidJobUpdate)
	}

	j.description = description

	j.updatedAt = time.Now()
	return nil
}

func (j *Job) WorkLocation() string {
	return j.workLocation
}

func (j *Job) UpdateWorkLocation(workLocation string) error {
	if !slices.Contains([]string{"remote", "office", "hybrid"}, workLocation) {
		return fmt.Errorf("%w: incorrect work location. (remote - office - hybrid)", ErrInvalidJobUpdate)
	}

	j.workLocation = workLocation

	j.updatedAt = time.Now()
	return nil
}

func (j *Job) WorkTime() string {
	return j.workTime
}

func (j *Job) UpdateWorkTime(workTime string) error {
	if !slices.Contains([]string{"full time", "part time"}, workTime) {
		return fmt.Errorf("%w: incorrect work time. (full time - part time)", ErrInvalidJobUpdate)
	}

	j.workTime = workTime

	j.updatedAt = time.Now()

	return nil
}

func (j *Job) Skills() []string {
	return j.skills
}

func (j *Job) UpdateSkills(skills []string) error {
	if len(skills) == 0 || len(skills) > 10 {
		return fmt.Errorf("%w: invalid number of skills (must be between 1 and 10)", ErrInvalidJobUpdate)
	}

	j.skills = skills

	j.updatedAt = time.Now()
	return nil
}

// YearsOfExperience returns the range of years of experiences, default value is 0 to 50
func (j *Job) YearsOfExperience() YearsOfExperienceRange {
	return j.yearsOfExperience
}

func (j *Job) UpdateYearsOfExperience(from, to int) error {
	if from < 0 {
		return fmt.Errorf("%w: invalid years of experience, can't be less than 0", ErrInvalidJobUpdate)
	}
	if from > to {
		return fmt.Errorf("%w: from can't be more than to", ErrInvalidJobUpdate)
	}

	j.yearsOfExperience.From = from
	j.yearsOfExperience.To = to

	j.updatedAt = time.Now()

	return nil
}

func (j *Job) ExpectedSalary() SalaryRange {
	return j.expectedSalary
}

// UpdateExpectedSalary updates the expected salary. the values are with
// currency lowest unit. for 100$ specify 10000 for 10000 cents.
func (j *Job) UpdateExpectedSalary(from, to int, currency string) error {
	if !slices.Contains([]string{"EGP", "USD"}, currency) {
		return fmt.Errorf("%w: invalid currency (EGP - USD)", ErrInvalidJobUpdate)
	}

	startingSalary, endingSalary := money.New(int64(from), currency), money.New(int64(to), currency)

	j.expectedSalary.From = startingSalary
	j.expectedSalary.To = endingSalary

	j.updatedAt = time.Now()

	return nil
}

// Status can be open, closed, or archived
func (j *Job) Status() string {
	return j.status.Status()
}

func (j *Job) UpdateStatus(status string) error {
	jobStatus, err := valueobject.NewJobStatus(status)
	if err != nil {
		return fmt.Errorf("%w: invalid job status (open - closed - archived)", valueobject.ErrInvalidJobStatus)
	}

	j.status = jobStatus

	j.updatedAt = time.Now()

	return nil
}

func (j *Job) CreatedAt() time.Time {
	return j.createdAt
}

func (j *Job) UpdatedAt() time.Time {
	return j.updatedAt
}

func (j *Job) NewApplication(
	name, email, title string,
	yearsOfExperience, hourlyRateAmount int,
	hourlyRateCurrency string,
	freelancerProfile, resumeURL string,
) (*Application, error) {
	for _, e := range j.applications {
		if e.email == email {
			return nil, fmt.Errorf("%w: applicant already applied", ErrJobConflict)
		}
	}

	application, err := newApplication(
		j.businessID,
		j.id,
		j.businessName,
		j.title,
		name,
		email,
		title,
		yearsOfExperience,
		hourlyRateAmount,
		hourlyRateCurrency,
		freelancerProfile,
		resumeURL,
	)
	if err != nil {
		return nil, err
	}

	j.applications = append(j.applications, *application)

	return application, nil
}

func (j *Job) Application(applicationID int) (*Application, error) {
	for _, a := range j.applications {
		if a.id == applicationID {
			return &a, nil
		}
	}

	return nil, fmt.Errorf("%w: application not found", domain.ErrNotFound)
}

// UpdateApplicationStatus Change the status of a job application. returns
// ErrJobConflict in case the job was not open, and returns ErrNotFound if the
// application is not found, and returns ErrInvalidApplicationUpdate in case of
// error with the provided application status
func (j *Job) UpdateApplicationStatus(applicationID int, status string) error {
	if j.status.Status() != "open" {
		return fmt.Errorf("%w: job is not open", ErrJobConflict)
	}

	var application *Application

	for _, a := range j.applications {
		if a.id == applicationID {
			application = &a
			break
		}
	}

	if application == nil {
		return fmt.Errorf("%w: application not found", domain.ErrNotFound)
	}

	return application.updateStatus(status)
}

// DeleteApplication deletes an application, it's advised to not use this
// method directly and use the application layer use case instead for
// protection from database not knowing about the deletion.
func (j *Job) DeleteApplication(applicationID int) error {
	if j.status.Status() != "open" {
		return fmt.Errorf("%w: job is not open", ErrJobConflict)
	}

	for i, a := range j.applications {
		if a.id == applicationID {
			j.applications = slices.Delete(j.applications, i, i+1)
			return nil
		}
	}

	return fmt.Errorf("%w: application not found", domain.ErrNotFound)
}
