package app

import (
	"errors"
)

var (
	// Standard error for unauthorized action in the application layer
	ErrUnauthorized = errors.New("unauthorized")
	ErrEditConflict = errors.New("edit conflict")
)

type Application struct {
	Commands *Commands
	Queries  *Queries
}

func New(commandRepos, queryRepos *Repositories) *Application {
	commands := &Commands{
		repo: commandRepos,
	}
	queries := &Queries{
		repo: queryRepos,
	}

	return &Application{
		Commands: commands,
		Queries:  queries,
	}
}

type Commands struct {
	repo *Repositories
}

type Queries struct {
	repo *Repositories
}
