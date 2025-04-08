package authuser

import (
	"errors"

	"github.com/ahmadabdelrazik/masarak/internal/domain/valueobject"
)

type AuthUser struct {
	name     string
	email    string
	Password *Password
	role     *valueobject.Role
}

// New - Creates a new user and validate the input data. For reconstructing
// users from database, use Instantiate instead.
func New(name, email, passwordText, role string) (*AuthUser, error) {
	password, err := createPassword(passwordText)
	if err != nil {
		return nil, err
	}

	r, err := valueobject.NewRole(role)
	if err != nil {
		return nil, err
	}

	return &AuthUser{
		name:     name,
		email:    email,
		role:     r,
		Password: password,
	}, nil
}

// Instantiate - Construct user from database.
func Instantiate(name, email string, passwordHash []byte, role string) *AuthUser {
	r, _ := valueobject.NewRole(role)

	return &AuthUser{
		name:     name,
		email:    email,
		Password: &Password{hash: passwordHash},
		role:     r,
	}
}

func (a *AuthUser) UpdateName(name string) error {
	if len(name) <= 2 {
		return errors.New("name must be longer than 2 bytes")
	} else if len(name) > 32 {
		return errors.New("name must be less than 33 bytes")
	}

	a.name = name

	return nil
}

func (a *AuthUser) UpdateRole(role string) error {
	r, err := valueobject.NewRole(role)
	if err != nil {
		return err
	}

	a.role = r

	return nil
}

func (a *AuthUser) UpdatePassword(oldPassword, newPassword string) error {
	if oldPassword == newPassword {
		return errors.New("new password must be different than old password")
	}

	if match, err := a.Password.Matches(oldPassword); err != nil {
		return err
	} else if !match {
		return errors.New("old password doesn't match")
	}

	password, err := createPassword(newPassword)
	if err != nil {
		return err
	}

	a.Password = password

	return nil
}

func (a *AuthUser) Email() string {
	return a.email
}

func (a *AuthUser) Role() string {
	return a.role.Role()
}

func (a *AuthUser) Name() string {
	return a.name
}

func (a *AuthUser) HasPermission(permission string) bool {
	return a.role.HasPermission(permission)
}
