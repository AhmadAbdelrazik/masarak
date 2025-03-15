package httpport

import (
	"net/http"

	"github.com/ahmadabdelrazik/layout/internal/common/auth"
)

func Routes(h *HttpServer, a *auth.AuthService) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /google_login", a.GoogleLogin)
	mux.HandleFunc("POST /google_callback", a.GoogleCallback)
	mux.HandleFunc("GET /google_callback", a.GoogleCallback)

	mux.Handle("/usecase", a.Middleware(h.CommandUseCase))

	return mux
}
