package postgres

import (
	"context"
	"database/sql"

	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

// AuthUserRepository - Postgres Implemntation for user repository
type AuthUserRepository struct {
	db *sql.DB
}

func (r *AuthUserRepository) Create(ctx context.Context, name, email, passwordText, role string) error {
	if _, err := authuser.New(name, email, passwordText, role); err != nil {
		return authuser.ErrInvalidProperty
	}

	query := `
	INSERT INTO users(email, name, password, role)
	VALUES($1,$2,$3,$4)`

	_, err := r.db.ExecContext(ctx, query, email, name, passwordText, role)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthUserRepository) GetByEmail(ctx context.Context, email string) (*authuser.User, error) {
	query := `
	SELECT name, password, role
	FROM users
	WHERE email = $1`

	var name, role string
	var passwordHash []byte

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&name,
		&passwordHash,
		&role,
	)
	if err != nil {
		return nil, err
	}

	user := authuser.Instantiate(name, email, passwordHash, role)

	return user, nil
}

func (r *AuthUserRepository) Save(ctx context.Context, email string, updateFn func(ctx context.Context, user *authuser.User) error) error {
	u, err := r.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	if err := updateFn(ctx, u); err != nil {
		return err
	}

	return nil
}
