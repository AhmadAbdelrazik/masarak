package command

import (
	"context"

	"github.com/ahmadabdelrazik/linkedout/internal/domain/applicant"
	users "github.com/ahmadabdelrazik/linkedout/internal/domain/user"
	"github.com/google/uuid"
)

type SelectPersonRole struct {
	User      users.User
	FirstName string
	LastName  string
	Role      string
}

type SelectPersonRoleHandler struct {
	applicantRepo applicant.Repository
	UserRepo      users.Repository
}

func NewSelectPersonRoleHandler(applicantRepo applicant.Repository, userRepo users.Repository) *SelectPersonRoleHandler {
	return &SelectPersonRoleHandler{
		applicantRepo: applicantRepo,
		UserRepo:      userRepo,
	}
}

func (h *SelectPersonRoleHandler) Handle(ctx context.Context, cmd SelectPersonRole) error {
	switch cmd.Role {

	case "applicant":
		a, err := applicant.NewApplicant(
			uuid.NewString(),
			cmd.FirstName,
			cmd.LastName,
			cmd.User.Email,
		)
		if err != nil {
			return err
		}

		if err := h.applicantRepo.Add(ctx, a); err != nil {
			return err
		}
	}

	return nil
}
