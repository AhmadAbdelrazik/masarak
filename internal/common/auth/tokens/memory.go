package tokens

import (
	"context"
	"errors"
	"sync"

	users "github.com/ahmadabdelrazik/linkedout/internal/domain/user"
)

type InMemoryTokenManager struct {
	userRepo users.Repository
	tokens   map[[32]byte]string // hash -> email

	sync.Mutex
}

func NewInMemoryTokenManager(userRepo users.Repository) *InMemoryTokenManager {
	return &InMemoryTokenManager{
		userRepo: userRepo,
	}
}

func (t *InMemoryTokenManager) GenerateToken(ctx context.Context, email string) (string, error) {
	_, err := t.userRepo.Get(ctx, email)
	if err != nil {
		return "", err
	}

	token := generateToken()

	hash := hashToken(token)

	t.Lock()
	defer t.Unlock()
	// think about multiple tokens
	t.tokens[hash] = email

	return token, nil
}

var (
	ErrInvalidToken = errors.New("invalid token")
)

func (t *InMemoryTokenManager) GetFromToken(ctx context.Context, token string) (*users.User, error) {
	t.Lock()
	defer t.Unlock()

	hash := hashToken(token)
	email, ok := t.tokens[hash]
	if !ok {
		return nil, ErrInvalidToken
	}

	return t.userRepo.Get(ctx, email)
}
