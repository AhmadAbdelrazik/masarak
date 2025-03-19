package authuser

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hash []byte
}

func (p *Password) Matches(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func newPassword(password string) (*Password, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Password{
		hash: hash,
	}, nil
}
