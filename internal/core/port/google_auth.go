package port

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/ahmadabdelrazik/linkedout/config"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/entity"
	"github.com/ahmadabdelrazik/linkedout/internal/core/domain/valueobject"
	"github.com/ahmadabdelrazik/linkedout/pkg/httperr"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleAuthService struct {
	cfg              *config.Config
	GoogleAuthConfig oauth2.Config
	userRepo         entity.AuthUserRepository
	tokenRepo        TokenRepository
}

func NewGoogleOAuthService(
	userRepo entity.AuthUserRepository,
	tokenRepo TokenRepository,
	cfg *config.Config,
) *GoogleAuthService {
	if cfg == nil {
		panic("config not found")
	}
	if userRepo == nil {
		panic("user not found")
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
		userRepo:         userRepo,
		tokenRepo:        tokenRepo,
		GoogleAuthConfig: googleAuthConfig,
	}

}

func (a GoogleAuthService) Middleware(next http.HandlerFunc) http.Handler {
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

		user, err := a.tokenRepo.GetFromToken(r.Context(), Token(cookie.Value))

		ctx := r.Context()

		ctx = context.WithValue(ctx, UserContextKey, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func (a GoogleAuthService) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := a.GoogleAuthConfig.AuthCodeURL(a.cfg.RandomState)

	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (a GoogleAuthService) GoogleCallback(w http.ResponseWriter, r *http.Request) {
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

	userRole, err := valueobject.NewRole("user")
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	user := &entity.AuthUser{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
		Role:  userRole,
	}

	err = a.userRepo.Add(r.Context(), user)
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	userToken, err := a.tokenRepo.GenerateToken(r.Context(), user.Email)
	if err != nil {
		httperr.ServerErrorResponse(w, r, err)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    string(userToken),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)
	if err := writeJSON(w, http.StatusOK, envelope{"message": "logged in successfully"}, nil); err != nil {
		httperr.ServerErrorResponse(w, r, err)
	}
}
