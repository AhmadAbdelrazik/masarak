package httpport

import (
	"github.com/ahmadabdelrazik/masarak/config"
	"github.com/ahmadabdelrazik/masarak/internal/app"
	"github.com/ahmadabdelrazik/masarak/internal/httpport/authservice"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

type HttpServer struct {
	app      *app.Application
	cfg      *config.Config
	auth     *authservice.AuthService
	userRepo authuser.UserRepository
}

func NewHttpServer(app *app.Application, cfg *config.Config, authService *authservice.AuthService, userRepo authuser.UserRepository) *HttpServer {
	if cfg == nil {
		panic("config not found")
	}
	if app == nil {
		panic("config not found")
	}
	if authService == nil {
		panic("token repo not found")
	}
	if userRepo == nil {
		panic("auth user repo not found")
	}

	return &HttpServer{
		app:      app,
		cfg:      cfg,
		auth:     authService,
		userRepo: userRepo,
	}
}
