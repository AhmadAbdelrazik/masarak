package freelancerprofile

import (
	"errors"

	"github.com/Rhymond/go-money"
)

type FreelancerProfile struct {
	name              string
	email             string
	pictureURL        string
	title             string
	skills            []string
	yearsOfExperience int
	hourlyRate        *money.Money
	resumeURL         string
}

var ErrSkillLimitReached = errors.New("skill number must not be more than 10")

func New(
	name, email, pictureURL, title string,
	skills []string,
	yearsOfExperience int,
	hourlyRate *money.Money,
) (*FreelancerProfile, error) {
	if len(skills) > 10 {
		return nil, ErrSkillLimitReached
	}

	return &FreelancerProfile{
		name:              name,
		email:             email,
		pictureURL:        pictureURL,
		title:             title,
		yearsOfExperience: yearsOfExperience,
		hourlyRate:        hourlyRate,
		skills:            skills,
	}, nil
}

func (f *FreelancerProfile) UpdateMainInfo(name, pictureURL, title string, yearsOfExperience int) error {
	f.name = name
	f.pictureURL = pictureURL
	f.title = title
	f.yearsOfExperience = yearsOfExperience

	return nil
}

func (f *FreelancerProfile) UpdateResumeURL(resumeURL string) error {
	f.resumeURL = resumeURL

	return nil
}

func (f *FreelancerProfile) UpdateHourlyRate(hourlyRate *money.Money) error {
	f.hourlyRate = hourlyRate

	return nil
}

func (f *FreelancerProfile) GetResumeURL() string {
	return f.resumeURL
}
