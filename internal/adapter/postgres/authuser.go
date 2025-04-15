package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/ahmadabdelrazik/masarak/internal/app"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

// AuthUserRepository - Postgres Implemntation for user repository
type AuthUserRepository struct {
	db     *sql.DB
	tokens *TokensRepository
}

func (r *AuthUserRepository) Create(ctx context.Context, username, email, name, passwordText, role string) (*authuser.User, error) {
	user, err := authuser.New(username, email, name, passwordText, role)
	if err != nil {
		return nil, authuser.ErrInvalidProperty
	}

	query := `
	INSERT INTO users(username, email, name, password, role)
	VALUES($1,$2,$3,$4,$5)
	RETURNING id`

	args := []interface{}{username, email, name, user.Password.Hash(), role}

	var id int
	if err := r.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		switch {
		case strings.Contains(err.Error(), "duplicate key"):
			return nil, authuser.ErrUserAlreadyExists
		default:
			return nil, fmt.Errorf("%w: %w", ErrDatabaseError, err)
		}
	}

	newUser := authuser.Instantiate(id, username, email, name, []byte(passwordText), role)

	return newUser, nil
}

func (r *AuthUserRepository) GetByEmail(ctx context.Context, email string) (*authuser.User, error) {
	query := `
	SELECT id, username, name, password, role
	FROM users
	WHERE email = $1`

	var id int
	var username, name, role string
	var passwordHash []byte

	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&id,
		&username,
		&name,
		&passwordHash,
		&role,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, authuser.ErrUserNotFound
		default:
			return nil, fmt.Errorf("%w: %w", ErrDatabaseError, err)
		}
	}

	user := authuser.Instantiate(id, username, email, name, passwordHash, role)

	return user, nil
}

func (r *AuthUserRepository) GetByUsername(ctx context.Context, username string) (*authuser.User, error) {
	query := `
	SELECT id, email, name, password, role
	FROM users
	WHERE email = $1`

	var id int
	var email, name, role string
	var passwordHash []byte

	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&id,
		&email,
		&name,
		&passwordHash,
		&role,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, authuser.ErrUserNotFound
		default:
			return nil, fmt.Errorf("%w: %w", ErrDatabaseError, err)
		}
	}

	user := authuser.Instantiate(id, username, email, name, passwordHash, role)

	return user, nil
}
func (r *AuthUserRepository) GetByID(ctx context.Context, id int) (*authuser.User, error) {
	query := `
	SELECT username, email, name, password, role
	FROM users
	WHERE email = $1`

	var username, email, name, role string
	var passwordHash []byte

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&username,
		&email,
		&name,
		&passwordHash,
		&role,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, authuser.ErrUserNotFound
		default:
			return nil, fmt.Errorf("%w: %w", ErrDatabaseError, err)
		}
	}

	user := authuser.Instantiate(id, username, email, name, passwordHash, role)

	return user, nil
}

func (r *AuthUserRepository) GetByToken(ctx context.Context, token authuser.Token) (*authuser.User, error) {
	r.tokens.Lock()
	defer r.tokens.Unlock()

	hash := string(hashToken(token))

	var id int

	for tokenHash, userID := range r.tokens.memory {
		if string(tokenHash[:32]) == hash {
			id = userID
			break
		}
	}

	if id == 0 {
		return nil, authuser.ErrUserNotFound
	}

	query := `
	SELECT username, email, name, password, role
	FROM users
	WHERE id = $1`

	var username, email, name, role string
	var passwordHash []byte

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&username,
		&email,
		&name,
		&passwordHash,
		&role,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, authuser.ErrUserNotFound
		default:
			return nil, fmt.Errorf("%w: %w", ErrDatabaseError, err)
		}
	}

	user := authuser.Instantiate(id, username, email, name, passwordHash, role)

	return user, nil
}

// Update - Gets the user by id, and pass it to the updateFn for updating
// user using it's method. After updating the user object, it's saved in the
// database with the condition that there was no updates since getting the
// user in the beginning. Returns ErrEditConflict in case of collision or
// ErrUserNotFound
func (r *AuthUserRepository) Update(ctx context.Context, id int, updateFn func(ctx context.Context, user *authuser.User) error) error {
	query := `
	SELECT username, email, name, password, role, version
	FROM users
	WHERE id = $1`

	var username, email, name, role string
	var passwordHash []byte

	// version ensures that there would be no update collisions
	var version int

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&username,
		&email,
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
			return fmt.Errorf("%w: %w", ErrDatabaseError, err)
		}
	}

	user := authuser.Instantiate(id, username, email, name, passwordHash, role)

	if err := updateFn(ctx, user); err != nil {
		return err
	}

	query = `
	UPDATE users
	SET name=$1, password=$2, role=$3, version = version + 1
	WHERE id = $4 AND version = $5`

	args := []interface{}{
		user.Name(),
		user.Password.Hash(),
		user.Role(),
		user.ID(),
		version,
	}

	if _, err := r.db.ExecContext(ctx, query, args...); err != nil {
		return app.ErrEditConflict
	}

	return nil
}
