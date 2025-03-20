package owner

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/company"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Owner struct {
	person *entity.Person
}

func (o *Owner) Email() string {
	return o.person.Email
}

func (o *Owner) Name() string {
	return o.person.Name
}

func (o *Owner) ID() uuid.UUID {
	return o.person.ID
}

func (o *Owner) CreateCompany(name, details, lineOfBusiness string) (*company.Company, error) {
	return company.New(name, details, lineOfBusiness, o.ID())
}

func New(name, email string) (*Owner, error) {
	person, err := entity.NewPerson(name, email)
	if err != nil {
		return nil, err
	}
	return &Owner{
		person: person,
	}, nil
}

var (
	ErrOwnerNotFound = errors.New("owner not found")
	ErrAlreadyExists = errors.New("owner already exists")
)

type Repository interface {
	Get(ctx context.Context, uid uuid.UUID) (*Owner, error)
	GetByEmail(ctx context.Context, email string) (*Owner, error)
	Create(ctx context.Context, owner *Owner) error
	Update(ctx context.Context, owner *Owner) error
	Delete(ctx context.Context, uid uuid.UUID) error
}
