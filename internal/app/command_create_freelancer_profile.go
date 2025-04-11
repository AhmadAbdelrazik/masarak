package app

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

type CreateFreelancerProfile struct {
	User               *authuser.User
	Email              string
	Name               string
	Title              string
	PictureURL         string
	Skills             []string
	YearsOfExperience  int
	HourlyRateAmount   int
	HourlyRateCurrency string
	ResumeURL          string
}

func (c *Commands) CreateFreelancerProfileHandler(ctx context.Context, cmd CreateFreelancerProfile) error {
	if !cmd.User.HasPermission("freelancer_profile.create") {
		return ErrUnauthorized
	} else if cmd.User.Email() != cmd.Email {
		return ErrUnauthorized
	}

	_, err := c.repo.FreelancerProfile.Create(
		ctx,
		cmd.Name,
		cmd.Email,
		cmd.Title,
		cmd.PictureURL,
		cmd.Skills,
		cmd.YearsOfExperience,
		cmd.HourlyRateAmount,
		cmd.HourlyRateCurrency,
		cmd.ResumeURL,
	)

	return err
}
