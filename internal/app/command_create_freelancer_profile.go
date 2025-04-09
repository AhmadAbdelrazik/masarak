package app

import "context"

type CreateFreelancerProfile struct {
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
