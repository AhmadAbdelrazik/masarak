package main

import (
	"github.com/ahmadabdelrazik/masarak/config"
	"github.com/ahmadabdelrazik/masarak/internal/adapter/memory"
	"github.com/ahmadabdelrazik/masarak/internal/core/app"
	"github.com/ahmadabdelrazik/masarak/internal/core/httpport"
	"github.com/ahmadabdelrazik/masarak/internal/core/httpport/auth"
	"github.com/ahmadabdelrazik/masarak/pkg/httpserver"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// initialize zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	mem := memory.NewMemory()
	repos := memory.NewInMemoryRepositories(mem)
	application := app.NewApplication(repos, repos)
	tokens := memory.NewInMemoryTokenRepository(mem, repos.AuthUsers)
	authService := auth.New(tokens, repos.AuthUsers)
	server := httpport.NewHttpServer(application, cfg, authService, repos.AuthUsers)

	httpserver.Serve(server.Routes(), cfg)
}
