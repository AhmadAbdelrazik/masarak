package httpport

import (
	"github.com/ahmadabdelrazik/linkedout/config"
	"github.com/ahmadabdelrazik/linkedout/internal/core/app"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
)

type HttpServer struct {
	app       *app.Application
	cfg       *config.Config
	auth      *GoogleAuthService
	tokenRepo TokenRepository
	userRepo  entity.AuthUserRepository
}

func NewHttpServer(app *app.Application, cfg *config.Config, google *GoogleAuthService, tokenRepo TokenRepository, userRepo entity.AuthUserRepository) *HttpServer {
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
