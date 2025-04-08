package main

import (
	"github.com/ahmadabdelrazik/masarak/config"
	"github.com/ahmadabdelrazik/masarak/internal/adapter/postgres"
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

	postgres.NewPostgresDB(cfg.DSN)
}
