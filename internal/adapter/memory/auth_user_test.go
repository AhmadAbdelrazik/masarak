package memory

import (
	"context"
	"testing"

	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/valueobject"
	"github.com/ahmadabdelrazik/masarak/pkg/assert"
)

func TestMemory_AuthUser(t *testing.T) {
	mem := NewMemory()
	userRepo := NewInMemoryAuthUserRepository(mem)
	tokenRepo := NewInMemoryTokenRepository(mem, userRepo)
	ctx := context.Background()

	role, err := valueobject.NewRole("user")
	assert.Nil(t, err)

	user, err := authuser.New("1", "John Doe", "johndoe@gmail.com", "johndoe1234", role)
	assert.Nil(t, err)

	err = userRepo.Add(ctx, user)
	assert.Nil(t, err)

	token, err := tokenRepo.GenerateToken(ctx, user.Email)
	assert.Nil(t, err)

	gottenUser, err := tokenRepo.GetFromToken(ctx, token)
	assert.Nil(t, err)

	assert.Equal(t, gottenUser.Email, user.Email)
}
