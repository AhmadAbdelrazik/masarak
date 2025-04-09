package app

import "context"

type CreateFreelancerProfile struct {
	Name              string
	Email             string
	PictureURL        string
	Title             string
	Skills            []string
	YearsOfExperience int
}

func (c *Commands) CreateFreelancerProfileHandler(ctx context.Context, cmd CreateFreelancerProfile) error {

}
