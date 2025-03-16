package memory

import (
	"sync"

	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/company"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/job"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/owner"
)

type Memory struct {
	owners    []*owner.Owner
	jobs      []*job.Job
	companies []*company.Company
	authUsers []*entity.AuthUser
	tokens    map[[32]byte]string // hash -> email

	sync.Mutex
}

func NewMemory() *Memory {
	return &Memory{
		owners:    make([]*owner.Owner, 0),
		jobs:      make([]*job.Job, 0),
		companies: make([]*company.Company, 0),
	}
}
