package app

import (
	"context"

	"github.com/Rhymond/go-money"
	"github.com/ahmadabdelrazik/masarak/internal/domain/freelancerprofile"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

type CreateFreelancerProfile struct {
	User               *authuser.User
	Username           string
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
		cmd.Username,
		cmd.Email,
		cmd.Name,
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

type UpdateFreelancerProfile struct {
	User               *authuser.User
	Username           string
	Name               *string
	Title              *string
	PictureURL         *string
	Skills             []string
	YearsOfExperience  *int
	HourlyRateAmount   *int
	HourlyRateCurrency *string
	ResumeURL          *string
}

// UpdateFreelancerProfileHandler - Updates freelancer Profile Info. Returns
// ErrProfileNotFound. can also return ErrSkillLimitReached or
// ErrInvalidYearsOfExperience or ErrInvalidHourlyRate for validation errors.
// or returns ErrEditConflict
func (c *Commands) UpdateFreelancerProfileHandler(ctx context.Context, cmd UpdateFreelancerProfile) error {
	if !cmd.User.HasPermission("freelancer_profile.update") || cmd.User.Username() != cmd.Username {
		return ErrUnauthorized
	}

	err := c.repo.FreelancerProfile.Update(
		ctx,
		cmd.Username,
		func(ctx context.Context, profile *freelancerprofile.FreelancerProfile) error {
			if cmd.Name != nil {
				if err := profile.UpdateName(*cmd.Name); err != nil {
					return err
				}
			}
			if cmd.Title != nil {
				if err := profile.UpdateTitle(*cmd.Title); err != nil {
					return err
				}
			}
			if cmd.PictureURL != nil {
				if err := profile.UpdatePictureURL(*cmd.PictureURL); err != nil {
					return err
				}
			}
			if cmd.Skills != nil {
				if err := profile.UpdateSkills(cmd.Skills); err != nil {
					return err
				}
			}
			if cmd.YearsOfExperience != nil {
				if err := profile.UpdateYearsOfExperience(*cmd.YearsOfExperience); err != nil {
					return err
				}
			}
			if cmd.HourlyRateAmount != nil {
				hourlyRate := money.New(int64(*cmd.HourlyRateAmount), profile.HourlyRate().Currency().Code)
				if err := profile.UpdateHourlyRate(hourlyRate); err != nil {
					return err
				}
			}
			if cmd.HourlyRateCurrency != nil {
				hourlyRate := money.New(profile.HourlyRate().Amount(), *cmd.HourlyRateCurrency)
				if err := profile.UpdateHourlyRate(hourlyRate); err != nil {
					return err
				}
			}
			if cmd.ResumeURL != nil {
				if err := profile.UpdateResumeURL(*cmd.ResumeURL); err != nil {
					return err
				}
			}

			return nil
		},
	)
	if err != nil {
		return err
	}

	return nil
}
