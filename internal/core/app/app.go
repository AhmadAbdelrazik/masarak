package app

import (
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/company"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/job"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/owner"
)

type Application struct {
	Commands *Commands
	Queries  *Queries
}

func NewApplication(repos *Repositories) *Application {
	commands := &Commands{
		CreateCompany: NewCreateCompanyHandler(repos.Companies, repos.Owner),
		CreateJob:     NewCreateJobHandler(repos.Owner, repos.Companies, repos.Jobs),
		CreateOwner:   NewCreateOwnerHandler(repos.Owner, repos.AuthUsers),
	}
	queries := &Queries{}

	return &Application{
		Commands: commands,
		Queries:  queries,
	}
}

type Commands struct {
	CreateCompany *CreateCompanyHandler
	CreateJob     *CreateJobHandler
	CreateOwner   *CreateOwnerHandler
}

type Queries struct {
}

type Repositories struct {
	Companies company.Repository
	Jobs      job.Repository
	Owner     owner.Repository
	AuthUsers entity.AuthUserRepository
}
