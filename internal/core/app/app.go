package app

import (
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/company"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/job"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/owner"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/talent"
)

type Application struct {
	Commands *Commands
	Queries  *Queries
}

func NewApplication(commandRepos, queryRepos *Repositories) *Application {
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

type Repositories struct {
	Companies company.Repository
	Jobs      job.Repository
	Owner     owner.Repository
	AuthUsers authuser.Repository
	Talents   talent.Repository
}
