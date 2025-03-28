package freelancerprofile

import (
	"errors"

	"github.com/Rhymond/go-money"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/entity"
)

type FreelancerProfile struct {
	profile *entity.FreelancerProfile
}

var ErrSkillLimitReached = errors.New("skill number must not be more than 10")

func New(
	name, email, pictureURL, title string,
	skills []string,
	yearsOfExperience int,
	hourlyRate *money.Money,
) (*FreelancerProfile, error) {
	profile, err := entity.NewFreelancerProfile(
		name,
		email,
		pictureURL,
		title,
		skills,
		yearsOfExperience,
		hourlyRate,
	)

	if err != nil {
		return nil, err
	}

	return &FreelancerProfile{
		profile: profile,
	}, nil
}

func (f *FreelancerProfile) UpdateMainInfo(name, pictureURL, title string, yearsOfExperience int) error {
	f.profile.Name = name
	f.profile.PictureURL = pictureURL
	f.profile.Title = title
	f.profile.YearsOfExperience = yearsOfExperience

	return nil
}

func (f *FreelancerProfile) UpdateResumeURL(resumeURL string) error {
	f.profile.ResumeURL = resumeURL

	return nil
}

func (f *FreelancerProfile) UpdateHourlyRate(hourlyRate *money.Money) error {
	f.profile.HourlyRate = hourlyRate

	return nil
}

func (f *FreelancerProfile) Data() entity.FreelancerProfile {
	return *f.profile
}

func (f *FreelancerProfile) Email() string {
	return f.profile.Email
}
