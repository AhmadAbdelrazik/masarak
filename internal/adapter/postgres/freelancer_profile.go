package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/ahmadabdelrazik/masarak/internal/app"
	"github.com/ahmadabdelrazik/masarak/internal/domain/freelancerprofile"
	"github.com/ahmadabdelrazik/masarak/pkg/filters"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type FreelancerProfileRepository struct {
	db *sqlx.DB
}

func (r *FreelancerProfileRepository) Create(
	ctx context.Context,
	username, email, name, title, pictureURL string,
	skills []string,
	yearsOfExperience int,
	hourlyRateAmount int,
	hourlyRateCurrency string,
	resumeURL string,
) (*freelancerprofile.FreelancerProfile, error) {
	_, err := freelancerprofile.New(
		username,
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
		return nil, fmt.Errorf("%w: %w", ErrDatabaseError, err)
	}

	query := `
	INSERT INTO freelancer_profiles(username, email, name, title, picture_url,
	skills, years_of_experience, hourly_rate_currency, hourly_rate_amount,
	resume_url)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id`

	args := []interface{}{
		username,
		email,
		name,
		title,
		pictureURL,
		pq.Array(skills),
		yearsOfExperience,
		hourlyRateCurrency,
		hourlyRateAmount,
		resumeURL,
	}

	var id int
	if err := r.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key"):
			return nil, freelancerprofile.ErrDuplicateProfile
		default:
			return nil, fmt.Errorf("%w: %w", ErrDatabaseError, err)
		}
	}

	profile := freelancerprofile.Instantiate(
		id,
		username,
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

func (r *FreelancerProfileRepository) GetByUsername(ctx context.Context, username string) (*freelancerprofile.FreelancerProfile, error) {
	query := `
	SELECT id, email, name, title, picture_url, skills, years_of_experience,
	hourly_rate_currency, hourly_rate_amount, resume_url
	FROM freelancer_profiles WHERE username = $1`

	var id int
	var email, name, title, pictureURL string
	var skills []string
	var yearsOfExperience, hourlyRateAmount int
	var hourlyRateCurrency, resumeURL string

	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&id,
		&email,
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
			return nil, fmt.Errorf("%w: %w", ErrDatabaseError, err)
		}
	}

	profile := freelancerprofile.Instantiate(
		id,
		username,
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

func (r *FreelancerProfileRepository) GetByEmail(ctx context.Context, email string) (*freelancerprofile.FreelancerProfile, error) {
	query := `
	SELECT id, username, name, title, picture_url, skills, years_of_experience,
	hourly_rate_currency, hourly_rate_amount, resume_url
	FROM freelancer_profiles WHERE email = $1`

	var id int
	var username, name, title, pictureURL string
	var skills []string
	var yearsOfExperience, hourlyRateAmount int
	var hourlyRateCurrency, resumeURL string

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&id,
		&username,
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
			return nil, fmt.Errorf("%w: %w", ErrDatabaseError, err)
		}
	}

	profile := freelancerprofile.Instantiate(
		id,
		username,
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

func (r *FreelancerProfileRepository) GetByID(ctx context.Context, id int) (*freelancerprofile.FreelancerProfile, error) {
	query := `
	SELECT username, email, name, title, picture_url, skills, years_of_experience,
	hourly_rate_currency, hourly_rate_amount, resume_url
	FROM freelancer_profiles WHERE id = $1`

	var username, email, name, title, pictureURL string
	var skills []string
	var yearsOfExperience, hourlyRateAmount int
	var hourlyRateCurrency, resumeURL string

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&username,
		&email,
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
			return nil, fmt.Errorf("%w: %w", ErrDatabaseError, err)
		}
	}

	profile := freelancerprofile.Instantiate(
		id,
		username,
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
	username string,
	updateFn func(ctx context.Context, profile *freelancerprofile.FreelancerProfile) error,
) error {
	query := `
	SELECT id, email, name, title, picture_url, skills, years_of_experience,
	hourly_rate_currency, hourly_rate_amount, resume_url, version
	FROM freelancer_profiles WHERE username = $1`

	var id int
	var email, name, title, pictureURL string
	var skills []string
	var yearsOfExperience, hourlyRateAmount, version int
	var hourlyRateCurrency, resumeURL string

	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&id,
		&email,
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
			return fmt.Errorf("%w: %w", ErrDatabaseError, err)
		}
	}

	profile := freelancerprofile.Instantiate(
		id,
		username,
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
	WHERE email = $9 AND version = $10`

	args := []interface{}{
		profile.Name(),
		profile.Title(),
		profile.PictureURL(),
		pq.Array(profile.Skills()),
		profile.YearsOfExperience(),
		profile.HourlyRate().Currency().Code,
		int(profile.HourlyRate().Amount()),
		profile.ResumeURL(),
		email,
		version,
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("%w: %w", app.ErrEditConflict, err)
	}

	return nil
}

func (r *FreelancerProfileRepository) Search(
	ctx context.Context,
	username, name, title string,
	skills []string,
	yearsOfExperience int,
	hourlyRateAmount int,
	hourlyRateCurrency string,
	filter filters.Filter,
) ([]freelancerprofile.FreelancerProfile, filters.Metadata, error) {
	query := fmt.Sprintf(`
	SELECT COUNT(*) OVER(), id, username, email, name, title, picture_url, skills, years_of_experience,
	hourly_rate_currency, hourly_rate_amount, resume_url
	FROM freelancer_profiles
	WHERE (to_tsvector('simple', username) @@ plainto_tsquery('simple', $1) OR $1 = '')
	AND (to_tsvector('simple', name) @@ plainto_tsquery('simple', $2) OR $2 = '')
	AND (to_tsvector('simple', title) @@ plainto_tsquery('simple', $3) OR $3 = '')
	AND (skills @> $4 OR $4 = '{}')
	AND (to_tsvector('simple', hourly_rate_currency) @@ plainto_tsquery('simple', $5) OR $5 = '')
	AND (years_of_experience = $6 OR $6 = -1)
	AND (hourly_rate_amount = $7 OR $7 = -1)
	ORDER BY %s %s, email ASC
	LIMIT $8 OFFSET $9`, filter.SortColumn(), filter.SortDirection())

	args := []interface{}{
		username,
		name,
		title,
		pq.Array(skills),
		hourlyRateCurrency,
		yearsOfExperience,
		hourlyRateAmount,
		filter.Limit(),
		filter.Offset(),
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, filters.Metadata{}, fmt.Errorf("%w: %w", ErrDatabaseError, err)
	}
	defer rows.Close()

	var profiles []freelancerprofile.FreelancerProfile
	totalRecords := 0

	for rows.Next() {
		var id int
		var username, email, name, title, pictureURL string
		var skills []string
		var yearsOfExperience, hourlyRateAmount int
		var hourlyRateCurrency, resumeURL string

		if err := rows.Scan(
			&totalRecords,
			&id,
			&username,
			&email,
			&name,
			&title,
			&pictureURL,
			pq.Array(&skills),
			&yearsOfExperience,
			&hourlyRateCurrency,
			&hourlyRateAmount,
			&resumeURL,
		); err != nil {
			return nil, filters.Metadata{}, fmt.Errorf("%w: %w", ErrDatabaseError, err)
		}

		profile := freelancerprofile.Instantiate(
			id,
			username,
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

		profiles = append(profiles, *profile)
	}

	if err := rows.Err(); err != nil {
		return nil, filters.Metadata{}, fmt.Errorf("%w: %w", ErrDatabaseError, err)
	}

	meta := filters.CalculateMetaData(totalRecords, filter.Page(), filter.PageSize())

	return profiles, meta, nil
}
