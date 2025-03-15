package tokens

import "sync"

type InMemoryTokenManager struct {
	users map[User][32]byte

	sync.Mutex
}

func NewInMemoryTokenManager() *InMemoryTokenManager {
	return &InMemoryTokenManager{
		users: make(map[User][32]byte),
	}
}

func (t *InMemoryTokenManager) GenerateToken(user User) (string, error) {
	token := generateToken()
	hash := hashToken(token)

	t.Lock()
	defer t.Unlock()
	// think about multiple tokens
	t.users[user] = hash

	return token, nil
}

func (t *InMemoryTokenManager) GetFromToken(token string) (User, error) {
	t.Lock()
	defer t.Unlock()

	return User{}, nil
}
