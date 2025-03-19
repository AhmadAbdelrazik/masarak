package httpport

import (
	"net/http"
)

func (h *HttpServer) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/google_login", h.auth.GoogleLogin)
	mux.HandleFunc("POST /v1/google_callback", h.auth.GoogleCallback)
	mux.HandleFunc("GET /v1/google_callback", h.auth.GoogleCallback)

	mux.HandleFunc("POST /v1/login", h.login)
	mux.HandleFunc("POST /v1/signup", h.Signup)

	mux.HandleFunc("GET /v1/health", h.HealthCheck)
	mux.Handle("GET /v1/auth_health", h.IsAuthenticated(h.HealthCheck))

	// Owner
	mux.Handle("POST /v1/owner", h.IsAuthenticated(h.postOwner))

	// Company
	mux.Handle("POST /v1/companies", h.IsAuthenticated(h.postCompany))

	// Job
	mux.Handle("POST /v1/jobs", h.IsAuthenticated(h.postJob))

	return mux
}

func (h *HttpServer) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthy\n"))
}
