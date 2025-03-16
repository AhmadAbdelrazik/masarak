package owner

import (
	"context"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Owner struct {
	person            *entity.Person
	ownedCompaniesIDs uuid.UUIDs
}

func (o *Owner) ID() uuid.UUID {
	return o.person.ID
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
	Create(ctx context.Context, owner *Owner) error
	Update(ctx context.Context, owner *Owner) error
	Delete(ctx context.Context, uid uuid.UUID) error
}
