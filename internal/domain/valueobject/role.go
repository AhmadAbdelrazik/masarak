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
		role: "user",
		permissions: []string{
			"role.select",
		},
	}
	roleFreelancer = Role{
		role: "freelancer",
		permissions: []string{
			"freelancer_profile.create",
			"freelancer_profile.read",
			"freelancer_profile.update",
			"freelancer_profile.delete",

			"resume.create",
			"resume.read",
			"resume.update",
			"resume.delete",

			"application.create",
			"application.read",
			"application.update",
			"application.delete",

			"job.read",

			"application_history.read",
		},
	}
	roleBusinessOwner = Role{
		role: "owner",
		permissions: []string{
			"business.create",
			"business.read",
			"business.update",
			"business.delete",

			"job.create",
			"job.read",
			"job.update",
			"job.delete",

			"application.read",
			"application.update",
		},
	}
	roleBusinessEmployee = Role{
		role: "employee",
		permissions: []string{
			"read.business",
			"update.business",

			"job.create",
			"job.read",
			"job.update",
			"job.delete",

			"application.read",
			"application.update",
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
