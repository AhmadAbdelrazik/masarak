package app

import (
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/company"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/job"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/owner"
)

type Application struct {
	Commands *Commands
	Queries  *Queries
}

func NewApplication(repos *Repositories) *Application {
	commands := &Commands{
		CreateCompany: NewCreateCompanyHandler(repos.Companies, repos.Owner),
		CreateJob:     NewCreateJobHandler(repos.Owner, repos.Companies, repos.Jobs),
		RegisterOwner: NewRegisterOwnerHandler(repos.Owner, repos.AuthUsers),
	}
	queries := &Queries{
		GetOwner: NewGetOwnerHandler(repos.Owner, repos.Companies),
	}

	return &Application{
		Commands: commands,
		Queries:  queries,
	}
}

type Commands struct {
	CreateCompany *CreateCompanyHandler
	CreateJob     *CreateJobHandler
	RegisterOwner *RegisterOwnerHandler
}

type Queries struct {
	GetOwner *GetOwnerHandler
}

type Repositories struct {
	Companies company.Repository
	Jobs      job.Repository
	Owner     owner.Repository
	AuthUsers authuser.Repository
}
