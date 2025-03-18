package httpport

import (
	"github.com/ahmadabdelrazik/linkedout/config"
	"github.com/ahmadabdelrazik/linkedout/internal/core/app"
)

type HttpServer struct {
	app  *app.Application
	cfg  *config.Config
	auth *GoogleAuthService
}

func NewHttpServer(app *app.Application, cfg *config.Config, google *GoogleAuthService) *HttpServer {
	if cfg == nil {
		panic("config not found")
	}
	if app == nil {
		panic("config not found")
	}
	if google == nil {
		panic("google auth service not found")
	}

	return &HttpServer{
		app:  app,
		cfg:  cfg,
		auth: google,
	}
}
