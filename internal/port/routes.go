package port

import (
	"net/http"
)

func (h *HttpServer) Routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./uploads"))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	mux.HandleFunc("GET /v1/google_login", h.auth.Google.GoogleLogin)
	mux.HandleFunc("POST /v1/google_callback", h.auth.Google.GoogleCallback)
	mux.HandleFunc("GET /v1/google_callback", h.auth.Google.GoogleCallback)

	mux.HandleFunc("POST /v1/login", h.auth.Login)
	mux.HandleFunc("POST /v1/signup", h.auth.Signup)
	mux.Handle("GET /v1/logout", h.auth.IsAuthenticated(h.auth.Logout))
	mux.Handle("POST /v1/select_role", h.auth.IsAuthenticated(h.selectRole))

	mux.HandleFunc("GET /v1/health", h.healthCheck)
	mux.Handle("GET /v1/auth_health", h.auth.IsAuthenticated(h.healthCheck))

	mux.Handle("POST /v1/freelancer_profiles", h.auth.IsAuthenticated(h.createFreelancerProfileHandler))
	mux.HandleFunc("GET /v1/freelancer_profiles/{username}", h.getFreelancerProfile)
	mux.HandleFunc("GET /v1/freelancer_profiles", h.searchFreelancerProfiles)
	mux.Handle("PATCH /v1/freelancer_profiles/{username}", h.auth.IsAuthenticated(h.updateFreelancerProfile))

	return mux
}

func (h *HttpServer) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthy\n"))
}
