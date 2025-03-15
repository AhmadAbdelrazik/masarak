package httpport

import (
	"net/http"

	"github.com/ahmadabdelrazik/layout/config"
	"github.com/ahmadabdelrazik/layout/internal/app"
	"github.com/ahmadabdelrazik/layout/internal/app/command"
)

type HttpServer struct {
	app app.Application
	cfg *config.Config
}

func NewHttpServer(app app.Application, cfg *config.Config) *HttpServer {
	return &HttpServer{app, cfg}
}

func (h HttpServer) CommandUseCase(w http.ResponseWriter, r *http.Request) {
	// 1. Parse the input
	// 2. Validate the input
	// 3. Place the input in the CommandUseCase

	cmd := command.UseCaseCommand{}

	err := h.app.Commands.UseCase.Handle(r.Context(), cmd)
	if err != nil {
		// Handle error
		return
	}

	// 4. Handle Happy Path Output
}
