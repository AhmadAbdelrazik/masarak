package db

import (
	"context"
	"database/sql"
	"errors"
)

// DB - a wrapper that hides if using a db or tx from the repository
type DB struct {
	db        *sql.DB
	tx        *sql.Tx
	committed bool
}

func NewDB(driver, dsn string) (*DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	return &DB{
		db: db,
	}, nil
}

func (db *DB) Begin() error {
	tx, err := db.db.Begin()
	if err != nil {
		return err
	}

	db.committed = false
	db.tx = tx

	return nil
}

func (db *DB) Commit() error {
	if db.tx == nil {
		return errors.New("no transaction")
	}

	if err := db.tx.Commit(); err != nil {
		db.tx.Rollback()
		return err
	}

	db.tx = nil
	db.committed = false

	return nil
}

func (db *DB) Rollback() error {
	if db.tx == nil {
		return nil
	}

	if err := db.tx.Rollback(); err != nil {
		return err
	}

	db.tx = nil
	db.committed = false

	return nil
}

func (db *DB) Query(query string, args ...any) (*sql.Rows, error) {
	if db.tx == nil {
		return db.db.Query(query, args...)
	} else {
		return db.tx.Query(query, args...)
	}
}

func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	if db.tx == nil {
		return db.db.QueryContext(ctx, query, args...)
	} else {
		return db.tx.QueryContext(ctx, query, args...)
	}
}

func (db *DB) Exec(query string, args ...any) (sql.Result, error) {
	if db.tx == nil {
		return db.db.Exec(query, args...)
	} else {
		return db.tx.Exec(query, args...)
	}
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	if db.tx == nil {
		return db.db.ExecContext(ctx, query, args...)
	} else {
		return db.tx.ExecContext(ctx, query, args...)
	}
}

func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	if db.tx == nil {
		return db.db.QueryRowContext(ctx, query, args...)
	} else {
		return db.tx.QueryRowContext(ctx, query, args...)
	}
}
