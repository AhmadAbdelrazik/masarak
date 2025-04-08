package authuser

import "github.com/ahmadabdelrazik/masarak/internal/domain/valueobject"

type AuthUser struct {
	name     string
	email    string
	Password *Password
	role     *valueobject.Role
}

func New(name, email, passwordText, role string) (*AuthUser, error) {
	password, err := newPassword(passwordText)
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

func Instantiate(name, email string, passwordHash []byte, role string) *AuthUser {
	r, _ := valueobject.NewRole(role)

	return &AuthUser{
		name:     name,
		email:    email,
		Password: &Password{hash: passwordHash},
		role:     r,
	}
}

func (a *AuthUser) Update(name string, role *valueobject.Role) error {
	a.name = name
	a.role = role

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
