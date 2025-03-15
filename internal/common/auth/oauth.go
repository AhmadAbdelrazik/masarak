package auth

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/ahmadabdelrazik/layout/config"
	"github.com/ahmadabdelrazik/layout/internal/common/auth/tokens"
	"github.com/ahmadabdelrazik/layout/internal/common/server"
	"github.com/ahmadabdelrazik/layout/internal/common/server/httperr"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type authConfig struct {
	GoogleLoginConfig oauth2.Config
}

type AuthService struct {
	cfg    *config.Config
	auth   *authConfig
	tokens tokens.TokenManager
}

type AuthServiceConfigs func(*AuthService) error

func WithTokenManager(t tokens.TokenManager, as *AuthService) error {
	as.tokens = t
	return nil
}

func WithInMemoryTokenManager(as *AuthService) error {
	memory := tokens.NewInMemoryTokenManager()
	return WithTokenManager(memory, as)
}

func NewAuthService(cfg *config.Config, cfgs ...AuthServiceConfigs) (*AuthService, error) {
	authCfg := &authConfig{}
	authCfg.GoogleLoginConfig = oauth2.Config{
		RedirectURL:  "http://localhost:8080/google_callback",
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	authService := &AuthService{
		cfg:  cfg,
		auth: authCfg,
	}

	for _, c := range cfgs {
		err := c(authService)
		if err != nil {
			return &AuthService{}, err
		}
	}

	return authService, nil
}

func (a AuthService) Middleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("session_id")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				httperr.AuthenticationErrorResponse(w, r)
			default:
				httperr.BadRequestResponse(w, r, err)
			}
			return
		}

		user, err := a.tokens.GetFromToken(cookie.Value)

		ctx := r.Context()

		ctx = context.WithValue(ctx, tokens.UserContextKey, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (a AuthService) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := a.auth.GoogleLoginConfig.AuthCodeURL(a.cfg.RandomState)

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (a AuthService) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	if state != a.cfg.RandomState {
		httperr.BadRequestResponse(w, r, errors.New("States don't match"))
		return
	}

	code := r.URL.Query().Get("code")

	token, err := a.auth.GoogleLoginConfig.Exchange(context.Background(), code)
	if err != nil {
		httperr.BadRequestResponse(w, r, errors.New("Code-Token Exchange Failed"))
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		httperr.BadRequestResponse(w, r, errors.New("User Data Fetch Failed"))
		return
	}

	userData, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		httperr.BadRequestResponse(w, r, errors.New("JSON parsing failed"))
		return
	}

	var input struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
	}

	err = json.Unmarshal(userData, &input)
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	user := tokens.User{
		UUID:        input.ID,
		Email:       input.Email,
		DisplayName: input.Name,
		Role:        "user",
	}

	userToken, err := a.tokens.GenerateToken(user)
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    userToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)
	if err := server.WriteJSON(w, http.StatusOK, server.Envelope{"message": "logged in successfully"}, nil); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}
}
