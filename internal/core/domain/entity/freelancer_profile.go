package entity

import (
	"errors"

	"github.com/Rhymond/go-money"
)

type FreelancerProfile struct {
	Name              string
	Email             string
	PictureURL        string
	Title             string
	Skills            []string
	YearsOfExperience int
	HourlyRate        *money.Money
	ResumeURL         string
}

var ErrSkillLimitReached = errors.New("skill number must not be more than 10")

func NewFreelancerProfile(
	name, email, pictureURL, title string,
	skills []string,
	yearsOfExperience int,
	hourlyRate *money.Money,
) (*FreelancerProfile, error) {
	if len(skills) > 10 {
		return nil, ErrSkillLimitReached
	}

	return &FreelancerProfile{
		Name:              name,
		Email:             email,
		PictureURL:        pictureURL,
		Title:             title,
		YearsOfExperience: yearsOfExperience,
		HourlyRate:        hourlyRate,
		Skills:            skills,
	}, nil
}
