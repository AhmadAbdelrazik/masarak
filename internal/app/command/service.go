package command

import "context"

// we can also define service interfaces that will be implemented in
// infrastructure layer here to be used by other commands it maybe an
// external service for example such as authentication

type ExternalService interface {
	Do(ctx context.Context) error
}
