package postgres

import (
	"database/sql"

	"github.com/ahmadabdelrazik/masarak/internal/app"
)

type PostgresDB struct {
	db *sql.DB
}

// NewPostgresDB - establishes a new postgres DB that implements the
// repositories in the domain layer. the data source name (dsn) is used to
// connect to the postgres db
func NewPostgresDB(dsn string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &PostgresDB{
		db: db,
	}, nil
}

func NewRepository(authUser *AuthUserRepository) *app.Repositories {
	return &app.Repositories{
		AuthUsers: authUser,
	}
}
