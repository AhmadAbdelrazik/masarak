package httpport

import (
	"github.com/ahmadabdelrazik/linkedout/config"
	"github.com/ahmadabdelrazik/linkedout/internal/app"
)

type HttpServer struct {
	app  app.Application
	cfg  *config.Config
	auth OAuthService
}

type ServerOption func(*HttpServer) error

func NewHttpServer(app app.Application, cfg *config.Config, opts ...ServerOption) (*HttpServer, error) {
	server := &HttpServer{
		app: app,
		cfg: cfg,
	}

	for _, opt := range opts {
		err := opt(server)
		if err != nil {
			return nil, err
		}
	}

	return server, nil
}

func WithOAuthService(AuthService OAuthService) ServerOption {
	return func(hs *HttpServer) error {
		hs.auth = AuthService
		return nil
	}
}
