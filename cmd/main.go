package main

import (
	"github.com/ahmadabdelrazik/linkedout/config"
	"github.com/ahmadabdelrazik/linkedout/internal/adapter/memory"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/company"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/job"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/owner"
	"github.com/ahmadabdelrazik/linkedout/internal/core/port"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// initialize zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	_, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	mem := memory.NewMemory()
	owner := memory.NewInMemoryOwnerRepository(mem)
	company := memory.NewInMemoryCompanyRepository(mem)
	job := memory.NewInMemoryJobRepository(mem)
	authUser := memory.NewInMemoryAuthUserRepository(mem)
	token := memory.NewInMemoryTokenRepository(mem, authUser)

	repos := Repos{
		owner:    owner,
		company:  company,
		job:      job,
		authUser: authUser,
		token:    token,
	}

	_ = repos
}

type Repos struct {
	owner    owner.Repository
	company  company.Repository
	job      job.Repository
	authUser entity.AuthUserRepository
	token    port.TokenRepository
}
