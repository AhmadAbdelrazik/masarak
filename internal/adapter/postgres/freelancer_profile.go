package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/ahmadabdelrazik/masarak/internal/app"
	"github.com/ahmadabdelrazik/masarak/internal/domain/freelancerprofile"
	"github.com/lib/pq"
)

type FreelancerProfileRepository struct {
	db *sql.DB
}

func (r *FreelancerProfileRepository) Create(

	ctx context.Context,
	email, name, title, pictureURL string,
	skills []string,
	yearsOfExperience int,
	hourlyRateAmount int,
	hourlyRateCurrency string,
	resumeURL string,

) (*freelancerprofile.FreelancerProfile, error) {
	profile, err := freelancerprofile.New(
		email,
		name,
		title,
		pictureURL,
		skills,
		yearsOfExperience,
		hourlyRateAmount,
		hourlyRateCurrency,
		resumeURL,
	)

	if err != nil {
		return nil, err
	}

	query := `
	INSERT INTO freelancer_profiles(email, name, title, picture_url,
	skills, years_of_experience, hourly_rate_currency, hourly_rate_amount,
	resume_url)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	if _, err := r.db.ExecContext(
		ctx,
		query,
		email,
		name,
		title,
		pictureURL,
		pq.Array(skills),
		yearsOfExperience,
		hourlyRateCurrency,
		hourlyRateAmount,
		resumeURL,
	); err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key"):
			return nil, freelancerprofile.ErrDuplicateProfile
		default:
			return nil, err
		}
	}

	return profile, nil
}

func (r *FreelancerProfileRepository) GetByEmail(ctx context.Context, email string) (*freelancerprofile.FreelancerProfile, error) {
	query := `
	SELECT name, title, picture_url, skills, years_of_experience,
	hourly_rate_currency, hourly_rate_amount, resume_url
	FROM freelancer_profiles WHERE email = $1`

	var name, title, pictureURL string
	var skills []string
	var yearsOfExperience, hourlyRateAmount int
	var hourlyRateCurrency, resumeURL string

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&name,
		&title,
		&pictureURL,
		pq.Array(&skills),
		&yearsOfExperience,
		&hourlyRateCurrency,
		&hourlyRateAmount,
		&resumeURL,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, freelancerprofile.ErrProfileNotFound
		default:
			return nil, err
		}
	}

	profile := freelancerprofile.Instantiate(
		email,
		name,
		title,
		pictureURL,
		skills,
		yearsOfExperience,
		hourlyRateAmount,
		hourlyRateCurrency,
		resumeURL,
	)

	return profile, nil
}

func (r *FreelancerProfileRepository) Update(
	ctx context.Context,
	email string,
	updateFn func(ctx context.Context, profile *freelancerprofile.FreelancerProfile) error,
) error {
	query := `
	SELECT name, title, picture_url, skills, years_of_experience,
	hourly_rate_currency, hourly_rate_amount, resume_url, version
	FROM freelancer_profiles WHERE email = $1`

	var name, title, pictureURL string
	var skills []string
	var yearsOfExperience, hourlyRateAmount, version int
	var hourlyRateCurrency, resumeURL string

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&name,
		&title,
		&pictureURL,
		pq.Array(&skills),
		&yearsOfExperience,
		&hourlyRateCurrency,
		&hourlyRateAmount,
		&resumeURL,
		&version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return freelancerprofile.ErrProfileNotFound
		default:
			return err
		}
	}

	profile := freelancerprofile.Instantiate(
		email,
		name,
		title,
		pictureURL,
		skills,
		yearsOfExperience,
		hourlyRateAmount,
		hourlyRateCurrency,
		resumeURL,
	)

	if err := updateFn(ctx, profile); err != nil {
		return err
	}

	query = `
	UPDATE freelancer_profiles
	SET name=$1, title=$2, picture_url=$3, skills=$4,
	years_of_experience=$5, hourly_rate_currency=$6,
	hourly_rate_amount=$7, resume_url=$8, version=version + 1
	WHERE email = $9`

	if _, err := r.db.ExecContext(
		ctx,
		query,
		profile.Name(),
		profile.Title(),
		profile.PictureURL(),
		pq.Array(profile.Skills()),
		profile.YearsOfExperience(),
		profile.HourlyRate().Currency().Code,
		int(profile.HourlyRate().Amount()),
		profile.ResumeURL(),
		email,
	); err != nil {
		fmt.Printf("err: %v\n", err)
		return app.ErrEditConflict
	}

	return nil
}
