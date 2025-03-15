package service

import (
	"context"

	"github.com/ahmadabdelrazik/layout/internal/app"
	"github.com/ahmadabdelrazik/layout/internal/app/command"
	"github.com/ahmadabdelrazik/layout/internal/app/query"
)

// NewApplication - Used to initialize an application and place the
// commands and queries and initialize them too. It may also be used
// to initialize clients to other services such as gRPC clients.
//
// It returns a function also that will be deferred in the main.go to
// close all the clients that we opened.
func NewApplication(ctx context.Context) (app.Application, func()) {
	return app.Application{
			Commands: app.Commands{
				UseCase: command.NewUseCommandHandler(),
			},
			Queries: app.Queries{
				UseCase: query.NewUseQueryHandler(),
			},
		}, func() {

		}
}
