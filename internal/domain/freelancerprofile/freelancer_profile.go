package freelancerprofile

import (
	"errors"

	"github.com/Rhymond/go-money"
)

type FreelancerProfile struct {
	email             string
	name              string
	title             string
	pictureURL        string
	skills            []string
	yearsOfExperience int
	hourlyRate        *money.Money
	resumeURL         string
}

var (
	ErrSkillLimitReached        = errors.New("skill number must not be more than 10")
	ErrInvalidYearsOfExperience = errors.New("invalid years of experience")
	ErrInvalidHourlyRate        = errors.New("invalid hourly rate")
)

// New - Creates new Freelancer Profile. for creating freelancer profile from
// database, use Instantiate instead. returns ErrSkillLimitReached if it
// exceeds the limit. returns ErrInvalidYearsOfExperience for invalid YoE
func New(
	email, name, title, pictureURL string,
	skills []string,
	yearsOfExperience int,
	hourlyRateAmount int,
	hourlyRateCurrency string,
	resumeURL string,
) (*FreelancerProfile, error) {
	if len(skills) > 10 {
		return nil, ErrSkillLimitReached
	}

	if yearsOfExperience < 0 || yearsOfExperience > 40 {
		return nil, ErrInvalidYearsOfExperience
	}

	if hourlyRateAmount < 0 {
		return nil, ErrInvalidHourlyRate
	}

	if hourlyRateCurrency != "EGP" && hourlyRateCurrency != "USD" {
		return nil, ErrInvalidHourlyRate
	}

	hourlyRate := money.New(int64(hourlyRateAmount), hourlyRateCurrency)

	return &FreelancerProfile{
		email:             email,
		name:              name,
		title:             title,
		pictureURL:        pictureURL,
		skills:            skills,
		yearsOfExperience: yearsOfExperience,
		hourlyRate:        hourlyRate,
		resumeURL:         resumeURL,
	}, nil
}

// Instantiate - Create a freelancer profile from database.
func Instantiate(
	email, name, title, pictureURL string,
	skills []string,
	yearsOfExperience int,
	hourlyRateAmount int,
	hourlyRateCurrency string,
	resumeURL string,
) *FreelancerProfile {
	hourlyRate := money.New(int64(hourlyRateAmount), hourlyRateCurrency)

	return &FreelancerProfile{
		email:             email,
		name:              name,
		title:             title,
		pictureURL:        pictureURL,
		skills:            skills,
		yearsOfExperience: yearsOfExperience,
		hourlyRate:        hourlyRate,
		resumeURL:         resumeURL,
	}
}

func (f *FreelancerProfile) Email() string {
	return f.email
}

func (f *FreelancerProfile) Name() string {
	return f.name
}

func (f *FreelancerProfile) Title() string {
	return f.title
}

func (f *FreelancerProfile) PictureURL() string {
	return f.pictureURL
}

func (f *FreelancerProfile) Skills() []string {
	return f.skills
}

func (f *FreelancerProfile) YearsOfExperience() int {
	return f.yearsOfExperience
}

func (f *FreelancerProfile) HourlyRate() *money.Money {
	return f.hourlyRate
}

func (f *FreelancerProfile) ResumeURL() string {
	return f.resumeURL
}

func (f *FreelancerProfile) UpdateName(name string) error {
	f.name = name
	return nil
}
func (f *FreelancerProfile) UpdateTitle(title string) error {
	f.title = title
	return nil
}

func (f *FreelancerProfile) UpdatePictureURL(pictureURL string) error {
	f.pictureURL = pictureURL
	return nil
}

func (f *FreelancerProfile) UpdateSkills(skills []string) error {
	if len(skills) > 10 {
		return ErrSkillLimitReached
	}

	f.skills = skills
	return nil
}

func (f *FreelancerProfile) UpdateYearsOfExperience(yearsOfExperience int) error {
	if yearsOfExperience < 0 || yearsOfExperience > 40 {
		return ErrInvalidYearsOfExperience
	}

	f.yearsOfExperience = yearsOfExperience
	return nil
}

func (f *FreelancerProfile) UpdateHourlyRate(hourlyRate *money.Money) error {
	f.hourlyRate = hourlyRate

	return nil
}

func (f *FreelancerProfile) UpdateResumeURL(resumeURL string) error {
	f.resumeURL = resumeURL
	return nil
}
