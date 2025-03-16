package entity

import "github.com/google/uuid"

type Person struct {
	ID    uuid.UUID
	Name  string
	Email string
}

func NewPerson(name, email string) (*Person, error) {
	return &Person{
		ID:    uuid.New(),
		Name:  name,
		Email: email,
	}, nil
}
