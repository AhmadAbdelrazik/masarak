package postgres

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"sync"

	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

type TokensRepository struct {
	db     *sql.DB
	memory map[[32]byte]int

	sync.Mutex
}

func (r *TokensRepository) GenerateToken(ctx context.Context, id int) (authuser.Token, error) {
	token, err := generateToken()
	if err != nil {
		return authuser.Token(""), err
	}
	hash := hashToken(token)

	r.Lock()
	defer r.Unlock()

	r.memory[[32]byte(hash)] = id

	return token, nil
}

func (r *TokensRepository) DeleteTokensByID(ctx context.Context, id int) error {
	r.Lock()
	defer r.Unlock()

	for hash, userID := range r.memory {
		if id == userID {
			delete(r.memory, hash)
		}
	}

	return nil
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
func hashToken(token authuser.Token) []byte {
	hash := sha256.Sum256([]byte(token))
	return hash[:]
}
