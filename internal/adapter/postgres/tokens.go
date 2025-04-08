package postgres

import "database/sql"

type TokensRepository struct {
	db *sql.DB
}
