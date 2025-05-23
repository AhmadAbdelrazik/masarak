package business

import (
	"fmt"
	"slices"

	"github.com/ahmadabdelrazik/masarak/internal/domain"
	"github.com/ahmadabdelrazik/masarak/pkg/validator"
)

type Business struct {
	id             int
	name           string
	businessEmail  string
	description    string
	imageURL       string
	jobs           []Job
	ownerEmail     string
	employeeEmails []string
}

var (
	ErrInvalidBusinessProperty = fmt.Errorf("%w: business", domain.ErrInvalidProperty)
	ErrInvalidBusinessUpdate   = fmt.Errorf("%w: business", domain.ErrInvalidUpdate)
	ErrBusinessConflict        = fmt.Errorf("%w: business conflict", domain.ErrInvalidProperty)
)

// New creates a new business and validate it's name, description, and emails.
// returns ErrInvalidBusinessProperty in case of any invalid properties.
func New(name, businessEmail, description, imageURL, ownerEmail string) (*Business, error) {
	if len(name) == 0 || len(name) > 30 {
		return nil, fmt.Errorf("%w: name must be between 0 and 30 bytes", ErrInvalidBusinessProperty)
	}

	if len(description) == 0 || len(description) > 1000 {
		return nil, fmt.Errorf("%w: description must be between 0 and 1000 bytes", ErrInvalidBusinessProperty)
	}

	if !validator.Matches(businessEmail, validator.EmailRX) {
		return nil, fmt.Errorf("%w: invalid business email", ErrInvalidBusinessProperty)
	}

	if !validator.Matches(ownerEmail, validator.EmailRX) {
		return nil, fmt.Errorf("%w: invalid owner email", ErrInvalidBusinessProperty)
	}

	return &Business{
		name:           name,
		businessEmail:  businessEmail,
		description:    description,
		imageURL:       imageURL,
		jobs:           make([]Job, 0),
		ownerEmail:     ownerEmail,
		employeeEmails: []string{},
	}, nil
}

// Instantiate Creates a new business from the database.
func Instantiate(
	id int,
	name, businessEmail, description, imageURL, ownerEmail string,
	jobs []Job,
	employeeEmails []string,
) *Business {
	return &Business{
		id:             id,
		name:           name,
		businessEmail:  businessEmail,
		description:    description,
		imageURL:       imageURL,
		ownerEmail:     ownerEmail,
		jobs:           jobs,
		employeeEmails: employeeEmails,
	}
}

func (b *Business) ID() int {
	return b.id
}

func (b *Business) Name() string {
	return b.name
}

func (b *Business) UpdateName(name string) error {
	if len(name) == 0 || len(name) > 30 {
		return fmt.Errorf("%w: name must be between 0 and 30 bytes", ErrInvalidBusinessProperty)
	}

	b.name = name

	for i := range b.jobs {
		b.jobs[i].businessName = name
	}

	return nil
}

func (b *Business) Description() string {
	return b.description
}

func (b *Business) UpdateDescription(description string) error {
	if len(description) == 0 || len(description) > 1000 {
		return fmt.Errorf("%w: description must be between 0 and 1000 bytes", ErrInvalidBusinessProperty)
	}

	b.description = description
	return nil
}

func (b *Business) OwnerEmail() string {
	return b.ownerEmail
}

func (b *Business) UpdateOwnerEmail(ownerEmail string) error {
	if !validator.Matches(ownerEmail, validator.EmailRX) {
		return fmt.Errorf("%w: invalid owner email", ErrInvalidBusinessProperty)
	}

	b.ownerEmail = ownerEmail

	return nil
}

func (b *Business) BusinessEmail() string {
	return b.businessEmail
}

func (b *Business) UpdateBusinessEmail(businessEmail string) error {
	if !validator.Matches(businessEmail, validator.EmailRX) {
		return fmt.Errorf("%w: invalid business email", ErrInvalidBusinessProperty)
	}

	b.businessEmail = businessEmail

	return nil
}

func (b *Business) IsEmployee(email string) bool {
	if email == b.ownerEmail {
		return true
	} else if slices.Contains(b.employeeEmails, email) {
		return true
	}

	return false
}

// AddEmployee adds a new employee to the business. if employee already exists
// it returns ErrInvalidBusinessUpdate
func (b *Business) AddEmployee(email string) error {
	if !validator.Matches(email, validator.EmailRX) {
		return fmt.Errorf("%w: invalid employee email", ErrInvalidBusinessProperty)
	}

	for _, e := range b.employeeEmails {
		if e == email {
			return fmt.Errorf("%w: employee already exists", ErrInvalidBusinessUpdate)
		}
	}

	b.employeeEmails = append(b.employeeEmails, email)

	return nil
}

// RemoveEmployee removes an existing employee from the business. returns
// ErrInvalidBusinessUpdate if not exists
func (b *Business) RemoveEmployee(email string) error {
	for i, e := range b.employeeEmails {
		if e == email {
			b.employeeEmails = slices.Delete(b.employeeEmails, i, i+1)
			return nil
		}
	}

	return fmt.Errorf("%w: employee doesn't exist", ErrInvalidBusinessUpdate)
}

func (b *Business) NewJob(title, description, workLocation, workTime string, skills []string) (*Job, error) {
	for _, j := range b.jobs {
		if j.title == title {
			return nil, fmt.Errorf("%w: job with the same name already exists", ErrBusinessConflict)
		}
	}

	job, err := newJob(
		b.id,
		b.name,
		b.imageURL,
		title,
		description,
		workLocation,
		workLocation,
		skills,
	)
	if err != nil {
		return nil, err
	}

	b.jobs = append(b.jobs, *job)

	return job, nil
}

// Job Gets a job from business by ID. returns ErrNotFound if not exists
func (b *Business) Job(jobID int) (*Job, error) {
	for _, j := range b.jobs {
		if j.ID() == jobID {
			return &j, nil
		}
	}

	return nil, fmt.Errorf("%w: job not found", domain.ErrNotFound)
}

// DeleteJob deletes a job from the business. It's advised not to use this
// method directly and rather use the use case related to deletion instead.
func (b *Business) DeleteJob(jobID int) error {
	for i, j := range b.jobs {
		if j.id == jobID {
			b.jobs = slices.Delete(b.jobs, i, i+1)
			return nil
		}
	}

	return fmt.Errorf("%w: job not found", domain.ErrNotFound)
}

func (b *Business) UpdateJobTitle(jobID int, title string) error {
	var job *Job
	for _, j := range b.jobs {
		if j.title == title {
			return fmt.Errorf("%w: job with the same title already exists", ErrInvalidBusinessUpdate)
		}
		if j.id == jobID {
			job = &j
		}
	}

	if job == nil {
		return fmt.Errorf("%w: job not found", domain.ErrNotFound)
	}

	return job.updateTitle(title)
}

func (b *Business) ImageURL() string {
	return b.imageURL
}

func (b *Business) UpdateImageURL(url string) error {
	b.imageURL = url

	for i := range b.jobs {
		b.jobs[i].businessImageURL = url
	}

	return nil
}
