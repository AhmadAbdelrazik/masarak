package entity

import (
	"github.com/google/uuid"
)

type Business struct {
	ID          uuid.UUID
	Name        string
	Email       string
	Description string
	ImageURL    string
}

func NewBusiness(name, email, description, imageURL string) (*Business, error) {
	business := &Business{
		ID:          uuid.New(),
		Name:        name,
		Email:       email,
		Description: description,
		ImageURL:    imageURL,
	}

	return business, nil
}
