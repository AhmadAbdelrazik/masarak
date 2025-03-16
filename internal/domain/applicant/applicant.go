package applicant

import (
	"context"
	"errors"

	"github.com/ahmadabdelrazik/linkedout/internal/domain/entity"
)

type Applicant struct {
	person *entity.Person
}

func (a *Applicant) GetEmail() string {
	return a.person.Email
}

func (a *Applicant) GetFullName() string {
	return a.person.FirstName + " " + a.person.LastName
}

func NewApplicant(id, firstName, lastName, email string) (*Applicant, error) {
	person, err := entity.NewPerson(id, firstName, lastName, email)
	if err != nil {
		return nil, err
	}

	return &Applicant{
		person: person,
	}, nil
}

var (
	ErrApplicantExists   = errors.New("applicant already exists")
	ErrApplicantNotFound = errors.New("applicant not found")
)

type Repository interface {
	Add(ctx context.Context, applicant *Applicant) error
	GetByEmail(ctx context.Context, email string) (*Applicant, error)
	GetApplicantNumbers(ctx context.Context) (int, error)
	GetAllApplicants(ctx context.Context) ([]*Applicant, error)
}
