package entity

type Business struct {
	ID          int
	Name        string
	Email       string
	Description string
	ImageURL    string
}

func NewBusiness(name, email, description, imageURL string) (*Business, error) {
	business := &Business{
		Name:        name,
		Email:       email,
		Description: description,
		ImageURL:    imageURL,
	}

	return business, nil
}
