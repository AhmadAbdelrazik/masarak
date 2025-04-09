package authservice

import (
	"github.com/ahmadabdelrazik/masarak/config"
	"github.com/ahmadabdelrazik/masarak/internal/app"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
)

// AuthService - HTTP Port service that implement basic authentication and
// OAuth for Google
type AuthService struct {
	tokenRepo authuser.TokenRepository
	userRepo  authuser.UserRepository
	Google    *GoogleAuthService
	app       *app.Application
	cfg       *config.Config
}

func New(
	app *app.Application,
	cfg *config.Config,
	userRepo authuser.UserRepository,
	tokenRepo authuser.TokenRepository,
) *AuthService {
	if tokenRepo == nil {
		panic("token repo not found")
	}
	if userRepo == nil {
		panic("user repo not found")
	}
	if app == nil {
		panic("application not found")
	}

	google := newGoogleOAuthService(tokenRepo, cfg, app)

	return &AuthService{
		tokenRepo: tokenRepo,
		userRepo:  userRepo,
		app:       app,
		cfg:       cfg,
		Google:    google,
	}
}
