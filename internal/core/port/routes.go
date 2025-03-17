package port

import (
	"net/http"
)

func (h *HttpServer) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /google_login", h.auth.GoogleLogin)
	mux.HandleFunc("POST /google_callback", h.auth.GoogleCallback)
	mux.HandleFunc("GET /google_callback", h.auth.GoogleCallback)

	return mux
}
