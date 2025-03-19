package main

import (
	"github.com/ahmadabdelrazik/masarak/config"
	"github.com/ahmadabdelrazik/masarak/internal/adapter/memory"
	"github.com/ahmadabdelrazik/masarak/internal/core/app"
	"github.com/ahmadabdelrazik/masarak/internal/core/httpport"
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
	application := app.NewApplication(repos)
	tokens := memory.NewInMemoryTokenRepository(mem, repos.AuthUsers)
	googleAuthService := httpport.NewGoogleOAuthService(repos.AuthUsers, tokens, cfg)
	server := httpport.NewHttpServer(application, cfg, googleAuthService, tokens, repos.AuthUsers)

	httpserver.Serve(server.Routes(), cfg)
}
