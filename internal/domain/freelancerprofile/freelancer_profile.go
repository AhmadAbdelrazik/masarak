package freelancerprofile

import (
	"fmt"

	"github.com/Rhymond/go-money"
)

type FreelancerProfile struct {
	id                int
	username          string
	email             string
	name              string
	title             string
	pictureURL        string
	skills            []string
	yearsOfExperience int
	hourlyRate        *money.Money
	resumeURL         string
}

// New - Creates new Freelancer Profile. for creating freelancer profile from
// database, use Instantiate instead. returns ErrSkillLimitReached if it
// exceeds the limit. returns ErrInvalidYearsOfExperience for invalid YoE
func New(
	username, email, name, title, pictureURL string,
	skills []string,
	yearsOfExperience int,
	hourlyRateAmount int,
	hourlyRateCurrency string,
	resumeURL string,
) (*FreelancerProfile, error) {
	if len(skills) > 10 {
		return nil, fmt.Errorf("%w: skill number must not be more than 10", ErrInvalidProperties)
	}
	if len(skills) == 0 {
		return nil, fmt.Errorf("%w: must specify at least one skill", ErrInvalidProperties)
	}

	if yearsOfExperience < 0 || yearsOfExperience > 40 {
		return nil, fmt.Errorf("%w: invalid years of experience", ErrInvalidProperties)
	}

	if hourlyRateAmount <= 0 {
		return nil, fmt.Errorf("%w: hourly rate amount must be higher than 0", ErrInvalidProperties)
	}

	if hourlyRateCurrency != "EGP" && hourlyRateCurrency != "USD" {
		return nil, fmt.Errorf("%w: hourly rate amount must be higher than 0", ErrInvalidProperties)
	}

	hourlyRate := money.New(int64(hourlyRateAmount), hourlyRateCurrency)

	return &FreelancerProfile{
		username:          username,
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
	id int,
	username, email, name, title, pictureURL string,
	skills []string,
	yearsOfExperience int,
	hourlyRateAmount int,
	hourlyRateCurrency string,
	resumeURL string,
) *FreelancerProfile {
	hourlyRate := money.New(int64(hourlyRateAmount), hourlyRateCurrency)

	return &FreelancerProfile{
		id:                id,
		username:          username,
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

func (f *FreelancerProfile) ID() int {
	return f.id
}

func (f *FreelancerProfile) Username() string {
	return f.username
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
		return fmt.Errorf("%w: skill number must not be more than 10", ErrInvalidProperties)
	}

	f.skills = skills
	return nil
}

func (f *FreelancerProfile) UpdateYearsOfExperience(yearsOfExperience int) error {
	if yearsOfExperience < 0 || yearsOfExperience > 40 {
		return fmt.Errorf("%w: invalid years of experience", ErrInvalidProperties)
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
