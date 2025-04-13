package authuser

import "context"

type Token string

type TokenRepository interface {
	// GenerateToken - Generate token for the user and save it in the
	// database.
	GenerateToken(ctx context.Context, id int) (Token, error)

	// DeleteTokensByID - Delete all tokens related to an id, used for
	// logging out
	DeleteTokensByID(ctx context.Context, id int) error
}
