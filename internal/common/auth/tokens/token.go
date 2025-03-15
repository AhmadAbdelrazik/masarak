package tokens

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
)

type TokenManager interface {
	GenerateToken(user User) (string, error)
	GetFromToken(string) (User, error)
}

type User struct {
	UUID        string
	Email       string
	Role        string
	DisplayName string
}

type ctxKey int

const (
	UserContextKey ctxKey = iota
)

var (
	// if we expect that the user of the function may be interested with concrete error,
	// it's a good idea to provide variable with this error
	NoUserInContextError = errors.New("No User in Context")
)

func UserFromCtx(ctx context.Context) (User, error) {
	u, ok := ctx.Value(UserContextKey).(User)
	if ok {
		return u, nil
	}

	return User{}, NoUserInContextError
}

func generateToken() string {
	bytes := make([]byte, 16)

	rand.Read(bytes)

	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(bytes)
}

func hashToken(token string) [32]byte {
	return sha256.Sum256([]byte(token))
}
