package httpport

import (
	"github.com/ahmadabdelrazik/masarak/config"
	"github.com/ahmadabdelrazik/masarak/internal/core/app"
	"github.com/ahmadabdelrazik/masarak/internal/core/domain/authuser"
)

type HttpServer struct {
	app       *app.Application
	cfg       *config.Config
	auth      *GoogleAuthService
	tokenRepo TokenRepository
	userRepo  authuser.Repository
}

func NewHttpServer(app *app.Application, cfg *config.Config, google *GoogleAuthService, tokenRepo TokenRepository, userRepo authuser.Repository) *HttpServer {
	if cfg == nil {
		panic("config not found")
	}
	if app == nil {
		panic("config not found")
	}
	if google == nil {
		panic("google auth service not found")
	}
	if tokenRepo == nil {
		panic("token repo not found")
	}
	if userRepo == nil {
		panic("auth user repo not found")
	}

	return &HttpServer{
		app:       app,
		cfg:       cfg,
		auth:      google,
		tokenRepo: tokenRepo,
		userRepo:  userRepo,
	}
}
