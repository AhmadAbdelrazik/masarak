package memory

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
	"github.com/ahmadabdelrazik/masarak/internal/core/httpport/auth"
)

type InMemoryTokenRepository struct {
	memory *Memory
	users  authuser.Repository
}

func NewInMemoryTokenRepository(memory *Memory, user authuser.Repository) *InMemoryTokenRepository {
	return &InMemoryTokenRepository{
		memory: memory,
		users:  user,
	}
}

func (r *InMemoryTokenRepository) GenerateToken(ctx context.Context, email string) (auth.Token, error) {
	token := generateToken()
	hash := hashToken(token)

	r.memory.Lock()
	defer r.memory.Unlock()

	r.memory.tokens[hash] = email

	return token, nil
}

func (r *InMemoryTokenRepository) GetFromToken(ctx context.Context, token auth.Token) (*authuser.AuthUser, error) {
	hash := hashToken(token)

	r.memory.Lock()
	email, ok := r.memory.tokens[hash]
	if !ok {
		return nil, authuser.ErrUserNotFound
	}
	r.memory.Unlock()

	return r.users.GetByEmail(ctx, email)
}

func generateToken() auth.Token {
	bytes := make([]byte, 16)

	rand.Read(bytes)

	return auth.Token(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(bytes))

}

func hashToken(token auth.Token) [32]byte {
	return sha256.Sum256([]byte(token))
}
