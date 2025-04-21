package postgres

import (
	"database/sql"

	"github.com/Rhymond/go-money"
	"github.com/ahmadabdelrazik/masarak/internal/domain/business"
	"github.com/ahmadabdelrazik/masarak/internal/domain/freelancerprofile"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
	"github.com/lib/pq"
)

type User struct {
	ID       int    `db:"id" `
	Username string `db:"username"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Role     string `db:"role"`
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
	Email             string   `db:"email"`
	Name              string   `db:"name"`
	Title             string   `db:"title"`
	PictureURL        string   `db:"picture_url"`
	Skills            []string `db:"skills"`
	YearsOfExperience int      `db:"years_of_experience"`
	HourlyRate        struct {
		Amount   int    `db:"amount"`
		Currency string `db:"currency"`
	} `db:"hourly_rate"`
	ResumeURL string `db:"resume_url"`
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
			Amount   int    `db:"amount"`
			Currency string `db:"currency"`
		}{
			Amount:   int(profile.HourlyRate().Amount()),
			Currency: profile.HourlyRate().Currency().Code,
		},
		ResumeURL: profile.ResumeURL(),
	}
}

type Business struct {
	Name          string `db:"name"`
	BusinessEmail string `db:"businessEmail"`
	Description   string `db:"description"`
	ImageURL      string `db:"imageURL"`
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
	BusinessID       sql.NullInt32  `db:"business_id"`
	BusinessName     sql.NullString `db:"business_name"`
	BusinessImageURL sql.NullString `db:"business_image_url"`

	ID           sql.NullInt32  `db:"iD"`
	Title        sql.NullString `db:"title"`
	Description  sql.NullString `db:"description"`
	WorkLocation sql.NullString `db:"workLocation"`
	WorkTime     sql.NullString `db:"workTime"`
	Skills       pq.StringArray `db:"skills"`

	YearsOfExperience struct {
		From sql.NullInt32 `db:"from"`
		To   sql.NullInt32 `db:"to"`
	}

	ExpectedSalary struct {
		From     sql.NullInt32  `db:"from"`
		To       sql.NullInt32  `db:"to"`
		Currency sql.NullString `db:"currency"`
	}

	Applications []Application
	Status       sql.NullString `db:"status"`
	CreatedAt    sql.NullTime   `db:"created_at"`
	UpdatedAt    sql.NullTime   `db:"updated_at"`
}

func toJob(job *Job) *business.Job {
	if job == nil {
		return nil
	}

	salaryRange := business.SalaryRange{
		From: money.New(int64(job.ExpectedSalary.From.Int32), money.ErrCurrencyMismatch.Error()),
		To:   money.New(int64(job.ExpectedSalary.To.Int32), money.ErrCurrencyMismatch.Error()),
	}

	var applications []business.Application

	for _, app := range job.Applications {
		applications = append(applications, toApplication(app))
	}

	return business.InstantiateJob(
		int(job.BusinessID.Int32),
		job.BusinessName.String,
		job.BusinessImageURL.String,
		int(job.ID.Int32),
		job.Title.String,
		job.Description.String,
		job.WorkLocation.String,
		job.WorkTime.String,
		[]string(job.Skills),
		business.YearsOfExperienceRange{
			From: int(job.YearsOfExperience.From.Int32),
			To:   int(job.YearsOfExperience.To.Int32),
		},
		salaryRange,
		job.Status.String,
		applications,
		job.CreatedAt.Time,
		job.UpdatedAt.Time,
	)
}

type Application struct {
	BusinessID        sql.NullInt32  `db:"business_id"`
	BusinessName      sql.NullString `db:"business_name"`
	JobID             sql.NullInt32  `db:"job_id"`
	JobTitle          sql.NullString `db:"job_title"`
	Id                sql.NullInt32  `db:"id"`
	Status            sql.NullString `db:"status"`
	Name              sql.NullString `db:"name"`
	Email             sql.NullString `db:"email"`
	Title             sql.NullString `db:"title"`
	YearsOfExperience sql.NullInt32  `db:"years_of_experience"`

	HourlyRate struct {
		Amount   sql.NullInt32  `db:"amount"`
		Currency sql.NullString `db:"currency"`
	} `db:"hourly_rate"`

	FreelancerProfile sql.NullString `db:"freelancer_profile"`
	ResumeURL         sql.NullString `db:"resume_url"`
	CreatedAt         sql.NullTime
	UpdatedAt         sql.NullTime
}

func toApplication(application Application) business.Application {
	return *business.InstantiateApplication(
		int(application.BusinessID.Int32),
		int(application.JobID.Int32),
		application.BusinessName.String,
		application.JobTitle.String,
		int(application.Id.Int32),
		application.Status.String,
		application.Name.String,
		application.Email.String,
		application.Title.String,
		int(application.YearsOfExperience.Int32),
		int(application.HourlyRate.Amount.Int32),
		application.HourlyRate.Currency.String,
		application.FreelancerProfile.String,
		application.ResumeURL.String,
		application.CreatedAt.Time,
		application.UpdatedAt.Time,
	)
}
