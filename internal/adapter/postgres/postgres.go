package postgres

import (
	"database/sql"

	"github.com/ahmadabdelrazik/masarak/internal/app"
)

func New(dsn string) (*app.Repositories, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &app.Repositories{
		Users:             &AuthUserRepository{db},
		Tokens:            &TokensRepository{db},
		FreelancerProfile: &FreelancerProfileRepository{db},
	}, nil
}
