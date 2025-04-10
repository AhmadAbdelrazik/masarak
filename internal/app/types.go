// types are the DTOs returned by the application layer queries.
package app

import (
	"github.com/ahmadabdelrazik/masarak/internal/domain/freelancerprofile"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func toUserDTO(user *authuser.User) User {
	return User{
		Name:  user.Name(),
		Email: user.Email(),
		Role:  user.Role(),
	}
}

type FreelancerProfile struct {
	Email             string   `json:"email"`
	Name              string   `json:"name"`
	Title             string   `json:"title"`
	PictureURL        string   `json:"picture_url"`
	Skills            []string `json:"skills"`
	YearsOfExperience int      `json:"years_of_experience"`
	HourlyRate        struct {
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	} `json:"hourly_rate"`
	ResumeURL string `json:"resume_url"`
}

func toFreelancerProfile(profile *freelancerprofile.FreelancerProfile) FreelancerProfile {
	return FreelancerProfile{
		Email:             profile.Email(),
		Name:              profile.Name(),
		Title:             profile.Title(),
		PictureURL:        profile.PictureURL(),
		Skills:            profile.Skills(),
		YearsOfExperience: profile.YearsOfExperience(),
		HourlyRate: struct {
			Amount   int    `json:"amount"`
			Currency string `json:"currency"`
		}{
			Amount:   int(profile.HourlyRate().Amount()),
			Currency: profile.HourlyRate().Currency().Code,
		},
		ResumeURL: profile.ResumeURL(),
	}
}
