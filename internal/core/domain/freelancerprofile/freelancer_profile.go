package freelancerprofile

import (
	"github.com/Rhymond/go-money"
)

type FreelancerProfile struct {
	name              string
	email             string
	pictureURL        string
	title             string
	yearsOfExperience int
	hourlyRate        *money.Money
	resumeURL         string
}

func New(name, email, pictureURL, title string, yearsOfExperiencce int, hourlyRate *money.Money) *FreelancerProfile {
	return &FreelancerProfile{
		name:              name,
		email:             email,
		pictureURL:        pictureURL,
		title:             title,
		yearsOfExperience: yearsOfExperiencce,
		hourlyRate:        hourlyRate,
	}
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
