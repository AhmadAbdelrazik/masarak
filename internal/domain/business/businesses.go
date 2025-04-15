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
	ImageURL       string
	jobs           []*Job
	ownerEmail     string
	employeeEmails []string
}

var (
	ErrInvalidBusinessProperty = fmt.Errorf("%w: business", domain.ErrInvalidProperty)
	ErrInvalidBusinessUpdate   = fmt.Errorf("%w: business", domain.ErrInvalidUpdate)
)

// New creates a new business and validate it's name, description, and emails.
// returns ErrInvalidBusinessProperty in case of any invalid properties.
func New(name, businessEmail, description, imageURL, ownerEmail string) (*Business, error) {
	if len(name) < 0 && len(name) > 30 {
		return nil, fmt.Errorf("%w: name must be between 0 and 30 bytes", ErrInvalidBusinessProperty)
	}

	if len(description) < 0 && len(description) > 1000 {
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
		ImageURL:       imageURL,
		jobs:           make([]*Job, 0),
		ownerEmail:     ownerEmail,
		employeeEmails: []string{},
	}, nil
}

// Instantiate Creates a new business from the database.
func Instantiate(
	id int,
	name, businessEmail, description, imageURL, ownerEmail string,
	jobs []*Job,
	employeeEmails []string,
) *Business {
	return &Business{
		id:             id,
		name:           name,
		businessEmail:  businessEmail,
		description:    description,
		ImageURL:       imageURL,
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
	if len(name) < 0 && len(name) > 30 {
		return fmt.Errorf("%w: name must be between 0 and 30 bytes", ErrInvalidBusinessProperty)
	}

	b.name = name
	return nil
}

func (b *Business) Description() string {
	return b.description
}

func (b *Business) UpdateDescription(description string) error {
	if len(description) < 0 && len(description) > 1000 {
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
	for _, e := range b.employeeEmails {
		if e == email {
			return true
		}
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
