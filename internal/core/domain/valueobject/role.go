package valueobject

import "github.com/pkg/errors"

type Role struct {
	role string
}

var (
	RoleUser  Role = Role{role: "User"}
	RoleAdmin      = Role{role: "Admin"}
	RoleOwner      = Role{role: "Owner"}

	ErrInvalidRole = errors.New("invalid role")
)

func NewRole(role string) (*Role, error) {
	switch role {
	case "admin", "Admin":
		return &RoleAdmin, nil
	case "owner", "Owner":
		return &RoleOwner, nil
	case "user", "User":
		return &RoleUser, nil
	default:
		return nil, ErrInvalidRole
	}
}
