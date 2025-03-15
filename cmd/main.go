package main

import (
	"context"

	"github.com/ahmadabdelrazik/layout/config"
	"github.com/ahmadabdelrazik/layout/internal/common/auth"
	"github.com/ahmadabdelrazik/layout/internal/common/server"
	httpport "github.com/ahmadabdelrazik/layout/internal/port/http"
	"github.com/ahmadabdelrazik/layout/internal/service"
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

	ctx := context.Background()
	app, f := service.NewApplication(ctx)
	defer f()

	srv := httpport.NewHttpServer(app, cfg)
	authSrv, err := auth.NewAuthService(cfg, auth.WithInMemoryTokenManager)

	routes := httpport.Routes(srv, authSrv)

	if err := server.RunHTTPServer(routes); err != nil {
		log.Error().Stack().Err(err).Msg("")
	}

}
