package postgres

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"

	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

type TokensRepository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) authuser.TokenRepository {
	return &TokensRepository{
		db: db,
	}
}

func (r *TokensRepository) GenerateToken(ctx context.Context, email string) (authuser.Token, error) {
	token, err := generateToken()
	if err != nil {
		return authuser.Token(""), err
	}
	hash := hashToken(token)

	query := `
	INSERT INTO tokens(token, email)
	VALUES ($1, $2)`

	if _, err := r.db.ExecContext(ctx, query, hash, email); err != nil {
		return authuser.Token(""), err
	}

	return token, nil
}

// generateToken - generate a 26 byte random token.
func generateToken() (authuser.Token, error) {
	randomBytes := make([]byte, 16)

	if _, err := rand.Read(randomBytes); err != nil {
		return authuser.Token(""), err
	}

	plainText := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	return authuser.Token(plainText), nil
}

// hashToken - return a 32 byte hash of the token to be stored in the database
func hashToken(token authuser.Token) [32]byte {
	return sha256.Sum256([]byte(token))
}
