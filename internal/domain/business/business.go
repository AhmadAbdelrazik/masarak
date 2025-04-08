package business

import (
	"github.com/ahmadabdelrazik/masarak/internal/domain/entity/job"
)

type Business struct {
	id             int
	name           string
	businessEmail  string
	description    string
	imageURL       string
	jobs           []*job.Job
	ownerEmail     string
	employeeEmails []string
}

func New(name, businessEmail, ownerEmail, description, imageURL string) (*Business, error) {
	business := &Business{
		name:          name,
		businessEmail: businessEmail,
		ownerEmail:    ownerEmail,
		description:   description,
		imageURL:      imageURL,
	}

	return business, nil
}

func ReconstituteBusiness(
	id int,
	name, businessEmail, ownerEmail, description, imageURL string,
	employeeEmails []string,
	jobs []*job.Job,
) (*Business, error) {
	business := &Business{
		id:             id,
		name:           name,
		businessEmail:  businessEmail,
		ownerEmail:     ownerEmail,
		description:    description,
		imageURL:       imageURL,
		employeeEmails: employeeEmails,
		jobs:           jobs,
	}

	return business, nil
}

func (b *Business) ID() int {
	return b.id
}

func (b *Business) Name() string {
	return b.name
}

// Email - The official email of the business itself, for the email of the Owner use OwnerEmail
func (b *Business) Email() string {
	return b.businessEmail
}

func (b *Business) ImageURL() string {
	return b.imageURL
}

// OwnerEmail - The email of the owner
func (b *Business) OwnerEmail() string {
	return b.ownerEmail
}
