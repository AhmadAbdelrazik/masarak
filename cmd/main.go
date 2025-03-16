package main

import (
	"context"

	"github.com/ahmadabdelrazik/linkedout/config"
	"github.com/ahmadabdelrazik/linkedout/internal/common/auth"
	"github.com/ahmadabdelrazik/linkedout/internal/common/server"
	httpport "github.com/ahmadabdelrazik/linkedout/internal/port/http"
	"github.com/ahmadabdelrazik/linkedout/internal/service"
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

	authSrv, err := auth.NewAuthService(
		cfg,
		auth.WithInMemoryTokenManager(),
		auth.WithInMemoryUserRepository(),
	)

	srv, err := httpport.NewHttpServer(
		app,
		cfg,
		httpport.WithOAuthService(authSrv),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	server.RunHTTPServer(srv.Routes())
}
