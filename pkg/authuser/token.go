package authuser

import "context"

type Token string

type TokenRepository interface {
	// GenerateToken - Generate token for the user and save it in the
	// database.
	GenerateToken(ctx context.Context, email string) (Token, error)
}
