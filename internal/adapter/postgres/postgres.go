package postgres

import "github.com/ahmadabdelrazik/masarak/pkg/db"

type PostgresDB struct {
	db *db.DB
}

// NewPostgresDB - establishes a new postgres DB that implements the
// repositories in the domain layer. the data source name (dsn) is used to
// connect to the postgres db
func NewPostgresDB(dsn string) (*PostgresDB, error) {
	db, err := db.NewDB("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &PostgresDB{
		db: db,
	}, nil
}

func (p *PostgresDB) Begin() error {
	return p.db.Begin()
}

func (p *PostgresDB) Commit() error {
	return p.db.Commit()
}

func (p *PostgresDB) Rollback() error {
	return p.db.Rollback()
}
