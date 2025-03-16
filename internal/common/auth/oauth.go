package auth

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/ahmadabdelrazik/linkedout/config"
	"github.com/ahmadabdelrazik/linkedout/internal/adapter"
	"github.com/ahmadabdelrazik/linkedout/internal/common/auth/tokens"
	"github.com/ahmadabdelrazik/linkedout/internal/common/server/httperr"
	command "github.com/ahmadabdelrazik/linkedout/internal/domain/user"
	users "github.com/ahmadabdelrazik/linkedout/internal/domain/user"
	httpport "github.com/ahmadabdelrazik/linkedout/internal/port/http"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type authConfig struct {
	GoogleLoginConfig oauth2.Config
}

type AuthService struct {
	cfg    *config.Config
	auth   *authConfig
	tokens httpport.TokenRepository
	users  users.Repository
}

type AuthServiceOption func(*AuthService) error

func NewAuthService(cfg *config.Config, opts ...AuthServiceOption) (*AuthService, error) {
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

	for _, opt := range opts {
		err := opt(authService)
		if err != nil {
			return &AuthService{}, err
		}
	}

	return authService, nil
}

func WithInMemoryUserRepository() AuthServiceOption {
	r := adapter.NewInMemoryUserRepository()
	return WithUserRepository(r)
}

func WithUserRepository(r command.Repository) AuthServiceOption {
	return func(as *AuthService) error {
		as.users = r
		return nil
	}
}

func WithTokenManager(t httpport.TokenRepository) AuthServiceOption {
	return func(as *AuthService) error {
		as.tokens = t
		return nil
	}
}

func WithInMemoryTokenManager(userRepo users.Repository) AuthServiceOption {
	memory := tokens.NewInMemoryTokenManager(userRepo)
	return WithTokenManager(memory)
}

func (a AuthService) AuthMiddleware(next http.HandlerFunc) http.Handler {
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

		user, err := a.tokens.GetFromToken(r.Context(), cookie.Value)

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

	_, err = a.users.Get(r.Context(), input.Email)
	if errors.Is(err, users.ErrUserNotFound) {
		user := &users.User{
			Email: input.Email,
			Name:  input.Name,
			Role:  "user",
		}

		err := a.users.Add(r.Context(), user)
		if err != nil {
			httperr.ServerErrorResponse(w, r, err)
			return
		}
	} else {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	userToken, err := a.tokens.GenerateToken(r.Context(), input.Email)
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
	// if err := server.WriteJSON(w, http.StatusOK, server.Envelope{"message": "logged in successfully"}, nil); err != nil {
	// 	httperr.ServerErrorResponse(w, r, err)
	// }

	http.Redirect(w, r, a.cfg.HostURL+"/create_account", http.StatusSeeOther)
}
