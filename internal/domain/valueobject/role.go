package valueobject

import (
	"github.com/pkg/errors"
)

type Role struct {
	role        string
	permissions []string
}

func (r *Role) HasPermission(permission string) bool {
	for _, p := range r.permissions {
		if p == permission {
			return true
		}
	}

	return false
}

func (r *Role) Role() string {
	return r.role
}

var (
	roleUser Role = Role{
		role:        "user",
		permissions: []string{},
	}
	roleFreelancer = Role{
		role: "freelancer",
		permissions: []string{
			"create:freelancer_profile",
			"read:freelancer_profile",
			"update:freelancer_profile",
			"delete:freelancer_profile",

			"create:resume",
			"read:resume",
			"update:resume",
			"delete:resume",

			"create:application",
			"read:application",
			"update:application",
			"delete:application",

			"read:job",

			"read:application_history",
		},
	}
	roleBusinessOwner = Role{
		role: "business_owner",
		permissions: []string{
			"create:business",
			"read:business",
			"update:business",
			"delete:business",

			"create:job",
			"read:job",
			"update:job",
			"delete:job",

			"read:application",
			"update:application",
		},
	}
	roleBusinessEmployee = Role{
		role: "business_employee",
		permissions: []string{
			"read:business",
			"update:business",

			"create:job",
			"read:job",
			"update:job",
			"delete:job",

			"read:application",
			"update:application",
		},
	}

	ErrInvalidRole = errors.New("invalid role")
)

func NewRole(role string) (*Role, error) {
	switch role {
	case "owner":
		return &roleBusinessOwner, nil
	case "employee":
		return &roleBusinessEmployee, nil
	case "user":
		return &roleUser, nil
	case "freelancer":
		return &roleFreelancer, nil
	default:
		return nil, ErrInvalidRole
	}
}
