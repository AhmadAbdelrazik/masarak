package memory

import (
	"sync"

	"github.com/ahmadabdelrazik/masarak/internal/core/app"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/company"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/job"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/owner"
)

type Memory struct {
	owners    []owner.Owner
	jobs      []job.Job
	companies []company.Company
	authUsers []authuser.AuthUser
	tokens    map[[32]byte]string // hash -> email

	sync.Mutex
}

func NewMemory() *Memory {
	return &Memory{
		owners:    make([]owner.Owner, 0),
		jobs:      make([]job.Job, 0),
		companies: make([]company.Company, 0),
		authUsers: make([]authuser.AuthUser, 0),
		tokens:    make(map[[32]byte]string),
	}
}

func NewInMemoryRepositories(mem *Memory) *app.Repositories {
	return &app.Repositories{
		Companies: NewInMemoryCompanyRepository(mem),
		Jobs:      NewInMemoryJobRepository(mem),
		Owner:     NewInMemoryOwnerRepository(mem),
		AuthUsers: NewInMemoryAuthUserRepository(mem),
	}
}
