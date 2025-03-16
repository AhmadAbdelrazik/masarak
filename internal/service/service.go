package service

import (
	"context"

	"github.com/ahmadabdelrazik/linkedout/internal/adapter"
	"github.com/ahmadabdelrazik/linkedout/internal/app"
	"github.com/ahmadabdelrazik/linkedout/internal/app/command"
	"github.com/ahmadabdelrazik/linkedout/internal/app/query"
)

// NewApplication - Used to initialize an application and place the
// commands and queries and initialize them too. It may also be used
// to initialize clients to other services such as gRPC clients.
//
// It returns a function also that will be deferred in the main.go to
// close all the clients that we opened.
func NewApplication(ctx context.Context) (app.Application, func()) {
	applicantRepo := adapter.NewInMemoryApplicantRepo()

	return app.Application{
			Commands: app.Commands{
				SelectPersonRole: command.NewSelectPersonRoleHandler(applicantRepo),
			},
			Queries: app.Queries{
				GetApplicant:       query.NewGetApplicantHandler(applicantRepo),
				GetApplicantNumber: query.NewGetApplicantNumberHandler(applicantRepo),
			},
		}, func() {

		}
}
