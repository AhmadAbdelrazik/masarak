package app

import (
	"github.com/ahmadabdelrazik/layout/internal/app/command"
	"github.com/ahmadabdelrazik/layout/internal/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	UseCase *command.UseCaseCommandHandler
}

type Queries struct {
	UseCase *query.UseCaseQueryHandler
}
