package query

import "context"

// Place the inputs needed for the use case.
type UseCaseQuery struct {
}

// Place all the needed repositories and services for the use case
type UseCaseQueryHandler struct {
}

// Factory pattern for making handlers, we can pass the needed services
// here. we can also use the config pattern to make it easier.
func NewUseQueryHandler() *UseCaseQueryHandler {
	return &UseCaseQueryHandler{}
}

// Implement the logic for the handle here, most of the time it would be
// orchestrating different functions provided by the repositories or the
// services.
func (c *UseCaseQueryHandler) Handle(ctx context.Context, cmd UseCaseQuery) error {
	return nil
}
