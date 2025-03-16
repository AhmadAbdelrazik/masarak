package app

import (
	"github.com/ahmadabdelrazik/linkedout/internal/app/command"
	"github.com/ahmadabdelrazik/linkedout/internal/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SelectPersonRole *command.SelectPersonRoleHandler
}

type Queries struct {
	GetApplicant       *query.GetApplicantHandler
	GetApplicantNumber *query.GetApplicantNumberHandler
}
