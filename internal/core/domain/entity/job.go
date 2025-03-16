package entity

import "github.com/google/uuid"

type Job struct {
	ID          uuid.UUID
	Title       string
	Description string
}

func NewJob(title, description string) (*Job, error) {
	return &Job{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
	}, nil
}
