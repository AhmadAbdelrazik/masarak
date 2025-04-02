package business

import (
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity/job"
	"github.com/google/uuid"
)

type Business struct {
	business       *entity.Business
	jobs           []*job.Job
	ownerEmail     string
	employeeEmails []string
}

func NewBusiness(name, businessEmail, ownerEmail, description, imageURL string) (*Business, error) {
	business, err := entity.NewBusiness(name, businessEmail, description, imageURL)

	if err != nil {
		return nil, err
	}

	return &Business{
		business:       business,
		ownerEmail:     ownerEmail,
		jobs:           make([]*job.Job, 0),
		employeeEmails: make([]string, 0),
	}, nil
}

func (b *Business) ID() uuid.UUID {
	return b.business.ID
}

func (b *Business) Name() string {
	return b.business.Name
}

// Email - The official email of the business itself, for the email of the Owner use OwnerEmail
func (b *Business) Email() string {
	return b.business.Email
}

func (b *Business) ImageURL() string {
	return b.business.ImageURL
}

// OwnerEmail - The email of the owner
func (b *Business) OwnerEmail() string {
	return b.ownerEmail
}
