package authservice

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/ahmadabdelrazik/masarak/config"
	"github.com/ahmadabdelrazik/masarak/internal/app"
	"github.com/ahmadabdelrazik/masarak/pkg/authuser"
	"github.com/ahmadabdelrazik/masarak/pkg/httperr"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleAuthService struct {
	cfg              *config.Config
	GoogleAuthConfig oauth2.Config
	app              *app.Application
	tokenRepo        authuser.TokenRepository
}

func newGoogleOAuthService(
	tokenRepo authuser.TokenRepository,
	cfg *config.Config,
	app *app.Application,
) *GoogleAuthService {
	if cfg == nil {
		panic("config not found")
	}
	if app == nil {
		panic("application not found")
	}
	if tokenRepo == nil {
		panic("token not found")
	}

	googleAuthConfig := oauth2.Config{
		RedirectURL:  "http://localhost:8080/google_callback",
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}

	return &GoogleAuthService{
		cfg:              cfg,
		tokenRepo:        tokenRepo,
		GoogleAuthConfig: googleAuthConfig,
	}

}

func (a *GoogleAuthService) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := a.GoogleAuthConfig.AuthCodeURL(a.cfg.RandomState)

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (a *GoogleAuthService) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	if state != a.cfg.RandomState {
		httperr.BadRequestResponse(w, r, errors.New("States don't match"))
		return
	}

	code := r.URL.Query().Get("code")

	token, err := a.GoogleAuthConfig.Exchange(context.Background(), code)
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

	if _, err := a.app.Queries.GetUser(context.Background(), input.Email); err != nil {
		switch {
		case errors.Is(err, authuser.ErrUserNotFound):
			cmd := app.CreateUser{
				Name:     input.Name,
				Email:    input.Email,
				Password: input.ID,
				Role:     "user",
			}

			if err := a.app.Commands.CreateUserHandler(context.Background(), cmd); err != nil {
				httperr.ServerErrorResponse(w, r, err)
				return
			}
		default:
			httperr.ServerErrorResponse(w, r, err)
			return
		}
	}

	cookie, err := getTokenCookie(r, input.Email, a.tokenRepo)
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}
	var output struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	output.Name = input.Name
	output.Email = input.Email

	http.SetCookie(w, cookie)
	if err := writeJSON(w, http.StatusCreated, envelope{"message": "logged in successfully", "user": output}, nil); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}
}
