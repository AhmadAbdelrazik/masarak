package valueobject

import (
	"strings"

	"github.com/pkg/errors"
)

type Role struct {
	role string
}

func (r *Role) Is(role string) bool {
	return r.role == strings.ToLower(role)
}

func (r *Role) String() string {
	return r.role
}

var (
	roleUser   Role = Role{role: "user"}
	roleAdmin       = Role{role: "admin"}
	roleOwner       = Role{role: "owner"}
	roleTalent      = Role{role: "talent"}

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
	case "talent", "Talent":
		return &roleUser, nil
	default:
		return nil, ErrInvalidRole
	}
}
