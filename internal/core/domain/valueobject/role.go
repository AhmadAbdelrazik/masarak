package valueobject

import "github.com/pkg/errors"

type Role struct {
	role string
}

var (
	roleUser  Role = Role{role: "User"}
	roleAdmin      = Role{role: "Admin"}
	roleOwner      = Role{role: "Owner"}

	ErrInvalidRole = errors.New("invalid role")
)

func NewRole(role string) (*Role, error) {
	switch role {
	case "admin", "Admin":
		return &roleAdmin, nil
	case "owner", "Owner":
		return &roleOwner, nil
	case "user", "User":
		return &roleUser, nil
	default:
		return nil, ErrInvalidRole
	}
}
