package postgres

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/ahmadabdelrazik/masarak/internal/app"
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
		switch {
		case strings.Contains(err.Error(), "duplicate key"):
			return authuser.ErrUserAlreadyExists
		default:
			return err
		}
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
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, authuser.ErrUserNotFound
		default:
			return nil, err
		}
	}

	user := authuser.Instantiate(name, email, passwordHash, role)

	return user, nil
}

func (r *AuthUserRepository) GetByToken(ctx context.Context, token authuser.Token) (*authuser.User, error) {
	query := `
	SELECT name, users.email, password, role
	FROM users
	JOIN tokens ON tokens.email = users.email
	WHERE tokens.token = $1`

	var name, email, role string
	var passwordHash []byte

	tokenHash := hashToken(token)

	err := r.db.QueryRowContext(ctx, query, tokenHash).Scan(
		&name,
		&email,
		&passwordHash,
		&role,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, authuser.ErrUserNotFound
		default:
			return nil, err
		}
	}

	user := authuser.Instantiate(name, email, passwordHash, role)

	return user, nil
}

// Update - Gets the user by email, and pass it to the updateFn for updating
// user using it's method. After updating the user object, it's saved in the
// database with the condition that there was no updates since getting the
// user in the beginning. Returns ErrEditConflict in case of collision or
// ErrUserNotFound
func (r *AuthUserRepository) Update(ctx context.Context, email string, updateFn func(ctx context.Context, user *authuser.User) error) error {
	query := `
	SELECT name, password, role, version
	FROM users
	WHERE email = $1`

	var name, role string
	var passwordHash []byte

	// version ensures that there would be no update collisions
	var version int

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&name,
		&passwordHash,
		&role,
		&version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return authuser.ErrUserNotFound
		default:
			return err
		}
	}

	user := authuser.Instantiate(name, email, passwordHash, role)

	if err := updateFn(ctx, user); err != nil {
		return err
	}

	query = `
	UPDATE users
	SET name=$1, password=$2, role=$3, version = version + 1
	WHERE email = $4 AND version = $5`

	if _, err := r.db.ExecContext(
		ctx,
		query,
		user.Name(),
		user.Password.Hash(),
		user.Role(),
		user.Email(),
		version,
	); err != nil {
		return app.ErrEditConflict
	}

	return nil
}
