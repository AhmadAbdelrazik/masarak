package postgres

import (
	"database/sql"
	"fmt"

	"github.com/ahmadabdelrazik/masarak/internal/app"

	_ "github.com/lib/pq"
)

var ErrDatabaseError = fmt.Errorf("%w: postgres", app.ErrDatabaseError)

func New(dsn string) (*app.Repositories, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	token := &TokensRepository{
		memory: make(map[[32]byte]int),
		db:     db,
	}

	return &app.Repositories{
		Users:             &AuthUserRepository{db, token},
		Tokens:            token,
		FreelancerProfile: &FreelancerProfileRepository{db},
	}, nil
}
