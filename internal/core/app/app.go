package app

import (
	"errors"

	applicationshistory "github.com/ahmadabdelrazik/masarak/internal/core/domain/applicationsHistory"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/business"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/freelancerprofile"
)

var (
	// Standard error for unauthorized action in the application layer
	ErrUnauthorized = errors.New("unauthorized")
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
	AuthUsers          authuser.Repository
	Businesses         business.Repository
	FreelancerProfile  freelancerprofile.Repository
	ApplicationHistory applicationshistory.Repository
}
