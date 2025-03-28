package applicationshistory

import "context"

type Repository interface {
	// GetByEmail - gets the application history for a freelancer
	// if not exists, create new one
	GetByEmail(ctx context.Context, email string) *ApplicationHistory
	Save(ctx context.Context, applicationHistory *ApplicationHistory) error
}
