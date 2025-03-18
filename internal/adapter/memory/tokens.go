package memory

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
	"github.com/ahmadabdelrazik/linkedout/internal/core/httpport"
)

type InMemoryTokenRepository struct {
	memory *Memory
	users  entity.AuthUserRepository
}

func NewInMemoryTokenRepository(memory *Memory, user entity.AuthUserRepository) *InMemoryTokenRepository {
	return &InMemoryTokenRepository{
		memory: memory,
		users:  user,
	}
}

func (r *InMemoryTokenRepository) GenerateToken(ctx context.Context, email string) (httpport.Token, error) {
	token := generateToken()
	hash := hashToken(token)

	r.memory.Lock()
	defer r.memory.Unlock()

	r.memory.tokens[hash] = email

	return token, nil
}

func (r *InMemoryTokenRepository) GetFromToken(ctx context.Context, token httpport.Token) (*entity.AuthUser, error) {
	hash := hashToken(token)

	r.memory.Lock()
	email := r.memory.tokens[hash]
	r.memory.Unlock()

	return r.users.GetByEmail(ctx, email)
}

func generateToken() httpport.Token {
	bytes := make([]byte, 16)

	rand.Read(bytes)

	return httpport.Token(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(bytes))

}

func hashToken(token httpport.Token) [32]byte {
	return sha256.Sum256([]byte(token))
}
