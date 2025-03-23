package httpport

import (
	"net/http"
)

func (h *HttpServer) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/google_login", h.auth.Google.GoogleLogin)
	mux.HandleFunc("POST /v1/google_callback", h.auth.Google.GoogleCallback)
	mux.HandleFunc("GET /v1/google_callback", h.auth.Google.GoogleCallback)

	mux.HandleFunc("POST /v1/login", h.auth.Login)
	mux.HandleFunc("POST /v1/signup", h.auth.Signup)

	mux.HandleFunc("GET /v1/health", h.HealthCheck)
	mux.Handle("GET /v1/auth_health", h.auth.IsAuthenticated(h.HealthCheck))

	// Owner
	mux.Handle("POST /v1/select_role", h.auth.IsAuthenticated(h.registerUser))
	mux.HandleFunc("GET /v1/owners/owner", h.getOwner)
	mux.HandleFunc("GET /v1/owners", h.getOwners)

	// Company
	mux.Handle("POST /v1/companies", h.auth.IsAuthenticated(h.postCompany))

	// Job
	mux.Handle("POST /v1/jobs", h.auth.IsAuthenticated(h.postJob))

	return mux
}

func (h *HttpServer) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthy\n"))
}
