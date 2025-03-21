package owner

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/company"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Owner struct {
	id       uuid.UUID
	authuser *authuser.AuthUser
}

func (o *Owner) Email() string {
	return o.authuser.Email
}

func (o *Owner) Name() string {
	return o.authuser.Name
}

func (o *Owner) ID() uuid.UUID {
	return o.id
}

func (o *Owner) CreateCompany(name, details, lineOfBusiness string) (*company.Company, error) {
	return company.New(name, details, lineOfBusiness, o.ID())
}

func New(authuser *authuser.AuthUser) (*Owner, error) {
	return &Owner{
		authuser: authuser,
		id:       uuid.New(),
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
