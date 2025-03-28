package memory

import (
	"sync"

	"github.com/ahmadabdelrazik/masarak/internal/core/app"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/business"
)

type Memory struct {
	authUsers  []*authuser.AuthUser
	tokens     map[[32]byte]string // hash -> email
	businesses []*business.Business

	sync.Mutex
}

func NewMemory() *Memory {
	return &Memory{
		authUsers:  make([]*authuser.AuthUser, 0),
		tokens:     make(map[[32]byte]string),
		businesses: make([]*business.Business, 0),
	}
}

func NewInMemoryRepositories(mem *Memory) *app.Repositories {
	return &app.Repositories{}
}
