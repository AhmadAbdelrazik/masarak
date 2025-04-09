package main

import (
	"github.com/ahmadabdelrazik/masarak/config"
	"github.com/ahmadabdelrazik/masarak/internal/adapter/postgres"
	"github.com/ahmadabdelrazik/masarak/internal/app"
	"github.com/ahmadabdelrazik/masarak/internal/port"
	"github.com/ahmadabdelrazik/masarak/pkg/httpserver"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
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

	repo, err := postgres.New(cfg.DSN)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	application := app.New(repo, repo)

	httpApp := port.NewHttpServer(
		application,
		cfg,
		repo.Users,
		repo.Tokens,
	)

	if err := httpserver.Serve(httpApp.Routes(), cfg); err != nil {
		log.Fatal().Err(err).Msg("")
	}

}
