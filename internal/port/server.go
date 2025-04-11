package port

import (
	"context"

	"github.com/ahmadabdelrazik/masarak/config"
	"github.com/ahmadabdelrazik/masarak/internal/app"
	"github.com/ahmadabdelrazik/masarak/internal/port/authservice"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

type HttpServer struct {
	app       *app.Application
	cfg       *config.Config
	auth      *authservice.AuthService
	userRepo  authuser.UserRepository
	tokenRepo authuser.TokenRepository
}

func NewHttpServer(
	app *app.Application,
	cfg *config.Config,
	userRepo authuser.UserRepository,
	tokenRepo authuser.TokenRepository,
) *HttpServer {
	authService := authservice.New(app, cfg, userRepo, tokenRepo)

	return &HttpServer{
		app:      app,
		cfg:      cfg,
		auth:     authService,
		userRepo: userRepo,
	}
}

func userFromCtx(ctx context.Context) (*authuser.User, error) {
	return authservice.UserFromCtx(ctx)
}
