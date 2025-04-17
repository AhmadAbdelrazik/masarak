// types are the DTOs returned by the application layer queries.
package app

import (
	"time"

	"github.com/ahmadabdelrazik/masarak/internal/domain/business"
	"github.com/ahmadabdelrazik/masarak/internal/domain/freelancerprofile"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func toUserDTO(user *authuser.User) User {
	if user == nil {
		return User{}
	}

	return User{
		ID:       user.ID(),
		Username: user.Username(),
		Name:     user.Name(),
		Email:    user.Email(),
		Role:     user.Role(),
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
	if profile == nil {
		return FreelancerProfile{}
	}

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

type Business struct {
	Name          string `json:"name"`
	BusinessEmail string `json:"businessEmail"`
	Description   string `json:"description"`
	ImageURL      string `json:"imageURL"`
}

func toBusiness(business *business.Business) Business {
	if business == nil {
		return Business{}
	}
	return Business{
		Name:          business.Name(),
		BusinessEmail: business.BusinessEmail(),
		Description:   business.Description(),
		ImageURL:      business.ImageURL(),
	}
}

type Job struct {
	BusinessID       int    `json:"business_id"`
	BusinessName     string `json:"business_name"`
	BusinessImageURL string `json:"business_image_url"`

	ID           int      `json:"iD"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	WorkLocation string   `json:"workLocation"`
	WorkTime     string   `json:"workTime"`
	Skills       []string `json:"skills"`

	YearsOfExperience struct {
		From int `json:"from"`
		To   int `json:"to"`
	} `json:"years_of_experience"`

	ExpectedSalary struct {
		From     int    `json:"from"`
		To       int    `json:"to"`
		Currency string `json:"currency"`
	} `json:"expected_salary"`

	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func toJob(job *business.Job) Job {
	if job == nil {
		return Job{}
	}

	jobDTO := Job{
		ID:               0,
		Title:            job.Title(),
		BusinessID:       job.BusinessID(),
		BusinessName:     job.BusinessName(),
		BusinessImageURL: job.BusinessImageURL(),
		Description:      job.Description(),
		WorkLocation:     job.WorkLocation(),
		WorkTime:         job.WorkTime(),
		Skills:           job.Skills(),
		YearsOfExperience: struct {
			From int "json:\"from\""
			To   int "json:\"to\""
		}{},
		ExpectedSalary: struct {
			From     int    "json:\"from\""
			To       int    "json:\"to\""
			Currency string "json:\"currency\""
		}{},
		Status:    job.Status(),
		CreatedAt: job.CreatedAt(),
		UpdatedAt: job.UpdatedAt(),
	}

	jobDTO.YearsOfExperience.From = job.YearsOfExperience().From
	jobDTO.YearsOfExperience.To = job.YearsOfExperience().To

	jobDTO.ExpectedSalary.From = int(job.ExpectedSalary().From.Amount())
	jobDTO.ExpectedSalary.To = int(job.ExpectedSalary().To.Amount())
	jobDTO.ExpectedSalary.Currency = job.ExpectedSalary().From.Currency().Code

	return jobDTO
}

type JobApplication struct {
	BusinessID        int    `json:"business_id"`
	BusinessName      string `json:"business_name"`
	JobID             int    `json:"job_id"`
	JobTitle          string `json:"job_title"`
	Id                int    `json:"id"`
	Status            string `json:"status"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Title             string `json:"title"`
	YearsOfExperience int    `json:"years_of_experience"`

	HourlyRate struct {
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
	} `json:"hourly_rate"`

	FreelancerProfile string `json:"freelancer_profile"`
	ResumeURL         string `json:"resume_url"`
}

func toApplication(application *business.Application) JobApplication {
	if application == nil {
		return JobApplication{}
	}

	hourlyRateAmount, hourlyRateCurrency := application.HourlyRate().Amount(), application.HourlyRate().Currency().Code

	return JobApplication{
		BusinessID:        application.BusinessID(),
		BusinessName:      application.BusinessName(),
		JobID:             application.JobID(),
		JobTitle:          application.JobTitle(),
		Id:                application.ID(),
		Status:            application.Status(),
		Name:              application.Name(),
		Email:             application.Email(),
		Title:             application.Title(),
		YearsOfExperience: application.YearsOfExperience(),
		HourlyRate: struct {
			Amount   int    "json:\"amount\""
			Currency string "json:\"currency\""
		}{
			Amount:   int(hourlyRateAmount),
			Currency: hourlyRateCurrency,
		},
		FreelancerProfile: application.FreelancerProfile,
		ResumeURL:         application.ResumeURL,
	}
}
