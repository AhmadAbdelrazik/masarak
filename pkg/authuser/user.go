package authuser

import (
	"errors"

	"github.com/ahmadabdelrazik/masarak/internal/domain/valueobject"
)

type User struct {
	id       int
	email    string
	username string
	name     string
	Password *Password
	role     *valueobject.Role
}

// New - Creates a new user and validate the input data. For reconstructing
// users from database, use Instantiate instead. the user instance will not
// have an id, the id would be given by the database
func New(username, email, name, passwordText, role string) (*User, error) {
	password, err := createPassword(passwordText)
	if err != nil {
		return nil, err
	}

	r, err := valueobject.NewRole(role)
	if err != nil {
		return nil, err
	}

	if !isValidUsername(username) {
		return nil, errors.New("invalid username")
	}

	return &User{
		username: username,
		email:    email,
		name:     name,
		role:     r,
		Password: password,
	}, nil
}

// Instantiate - Construct user from database.
func Instantiate(id int, username, email, name string, passwordHash []byte, role string) *User {
	r, err := valueobject.NewRole(role)
	if err != nil {
		panic(err)
	}

	return &User{
		id:       id,
		username: username,
		email:    email,
		name:     name,
		Password: &Password{hash: passwordHash},
		role:     r,
	}
}

func (a *User) UpdateName(name string) error {
	if len(name) <= 2 {
		return errors.New("name must be longer than 2 bytes")
	} else if len(name) > 32 {
		return errors.New("name must be less than 33 bytes")
	}

	a.name = name

	return nil
}

func (a *User) UpdateRole(role string) error {
	r, err := valueobject.NewRole(role)
	if err != nil {
		return err
	}

	a.role = r

	return nil
}

func (a *User) UpdatePassword(oldPassword, newPassword string) error {
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

func (a *User) ID() int {
	return a.id
}

func (a *User) Username() string {
	return a.username
}

func (a *User) Email() string {
	return a.email
}

func (a *User) Role() string {
	return a.role.Role()
}

func (a *User) Name() string {
	return a.name
}

func (a *User) HasPermission(permission string) bool {
	return a.role.HasPermission(permission)
}

// isValidUsername checks if the username characters all are either
// alphanumeric or underscore
func isValidUsername(username string) bool {
	for _, c := range username {
		if c >= 'a' && c <= 'z' {
			continue
		} else if c >= '0' && c <= '9' {
			continue
		} else {
			return false
		}
	}

	return true
}
