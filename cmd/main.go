package main

import (
	"github.com/ahmadabdelrazik/linkedout/config"
	"github.com/ahmadabdelrazik/linkedout/internal/adapter/memory"
	"github.com/ahmadabdelrazik/linkedout/internal/core/app"
	"github.com/ahmadabdelrazik/linkedout/internal/core/httpport"
	"github.com/ahmadabdelrazik/linkedout/pkg/httpserver"
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
	application := app.NewApplication(repos)
	tokens := memory.NewInMemoryTokenRepository(mem, repos.AuthUsers)
	googleAuthService := httpport.NewGoogleOAuthService(repos.AuthUsers, tokens, cfg)
	server := httpport.NewHttpServer(application, cfg, googleAuthService)

	httpserver.Serve(server.Routes(), cfg)
}
