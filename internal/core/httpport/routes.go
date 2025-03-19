package httpport

import (
	"net/http"
)

func (h *HttpServer) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /google_login", h.auth.GoogleLogin)
	mux.HandleFunc("POST /google_callback", h.auth.GoogleCallback)
	mux.HandleFunc("GET /google_callback", h.auth.GoogleCallback)

	mux.HandleFunc("POST /login", h.login)
	mux.HandleFunc("POST /signup", h.Signup)

	mux.HandleFunc("GET /health", h.HealthCheck)
	mux.Handle("GET /auth_health", h.IsAuthenticated(h.HealthCheck))

	return mux
}

func (h *HttpServer) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthy\n"))
}
