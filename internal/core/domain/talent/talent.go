package talent

import (
	"context"
	"errors"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
	"github.com/google/uuid"
)

type Talent struct {
	person *authuser.AuthUser
	id     uuid.UUID
}

func New(user *authuser.AuthUser) (*Talent, error) {
	return &Talent{
		person: user,
		id:     uuid.New(),
	}, nil
}

func (t *Talent) ID() uuid.UUID {
	return t.id
}

func (t *Talent) Name() string {
	return t.person.Name
}

func (t *Talent) Email() string {
	return t.person.Email
}

var (
	ErrTalentNotFound = errors.New("talent not found")
	ErrAlreadyExists  = errors.New("talent already exists")
)

type Repository interface {
	Get(ctx context.Context, uid uuid.UUID) (*Talent, error)
	GetByEmail(ctx context.Context, email string) (*Talent, error)
	Create(ctx context.Context, talent *Talent) error
	Update(ctx context.Context, talent *Talent) error
	Delete(ctx context.Context, uid uuid.UUID) error
}
