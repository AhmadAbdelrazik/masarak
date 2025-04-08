package postgres

import (
	"context"

	"github.com/Rhymond/go-money"
	"github.com/ahmadabdelrazik/masarak/internal/domain/freelancerprofile"
	"github.com/ahmadabdelrazik/masarak/pkg/db"
	"github.com/lib/pq"
)

type FreelancerProfileRepository struct {
	db *db.DB
}

func newFreelancerProfileRepository(db *db.DB) *FreelancerProfileRepository {
	return &FreelancerProfileRepository{
		db: db,
	}
}

func (r *FreelancerProfileRepository) Create(ctx context.Context, name, email,
	pictureURL, title string, skills []string, yearsOfExperience int,
	hourlyRate *money.Money) (*freelancerprofile.FreelancerProfile,
	error) {

	profile, err := freelancerprofile.New(name, email, pictureURL, title, skills, yearsOfExperience, hourlyRate)
	if err != nil {
		return nil, err
	}

	query := `
	INSERT INTO freelancer_profile(email, name, picture_url, title, years_of_experience, skills, hourly_rate_amount, hourly_rate_currency)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	currency, amount := hourlyRate.Currency().Code, hourlyRate.Amount()

	if _, err := r.db.ExecContext(ctx, query, email, name, pictureURL, title, yearsOfExperience, pq.Array(skills), amount, currency); err != nil {
		return nil, err
	}

	return profile, nil
}

func (r *FreelancerProfileRepository) GetByEmail(ctx context.Context, email string) (*freelancerprofile.FreelancerProfile, error) {
	query := `
	SELECT name, picture_url, title, years_of_experience, hourly_rate_amount, hourly_rate_currency
	FROM freelancer_profile
	WHERE email = $1
	`

	var name, pictureURL, title, currency string
	var yearsOfExperience, hourlyRateAmount int

	if err := r.db.QueryRowContext(ctx, query, email).Scan(
		&name,
		&pictureURL,
		&title,
		&yearsOfExperience,
		&hourlyRateAmount,
		&currency,
	); err != nil {
		return nil, err
	}

	return nil, nil
}
