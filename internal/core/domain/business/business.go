package business

import (
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity"
	"github.com/google/uuid"
)

type Business struct {
	business *entity.Business
}

func NewBusiness(name, email, description, imageURL string) (*Business, error) {
	business, err := entity.NewBusiness(name, email, description, imageURL)

	if err != nil {
		return nil, err
	}

	return &Business{
		business: business,
	}, nil
}

func (b *Business) ID() uuid.UUID {
	return b.business.ID
}

func (b *Business) Name() string {
	return b.business.Name
}

func (b *Business) Email() string {
	return b.business.Email
}

func (b *Business) ImageURL() string {
	return b.business.ImageURL
}
