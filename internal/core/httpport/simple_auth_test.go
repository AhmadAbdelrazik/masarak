package httpport_test

import (
	"net/http"
	"testing"

	"github.com/ahmadabdelrazik/masarak/pkg/assert"
)

func TestAuthUser(t *testing.T) {
	t.Parallel()
	tc := NewTestClient(t)
	defer tc.Close()

	t.Run("Register", func(t *testing.T) {
		testRegisterUser(t, tc)
	})
	t.Run("Login", func(t *testing.T) {
		testLoggedInUser(t, tc)
	})
}

func testRegisterUser(t *testing.T, tc *TestClient) {
	t.Run("not registered user", func(t *testing.T) {
		var input struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		input.Name = "John Doe"
		input.Email = "johndoe@example.com"
		input.Password = "john1234"

		res, err := tc.Post("/v1/signup", input)
		assert.Nil(t, err)

		assert.Equal(t, res.StatusCode, http.StatusCreated)

		cookie := tc.GetCookie(res, "session_id")
		assert.True(t, cookie.Value != "")

		var output struct {
			Message string `json:"message"`
		}
		tc.ParseResponseBody(res, &output)

		assert.Equal(t, output.Message, "registered successfully")

	})
	t.Run("already registered user", func(t *testing.T) {
		var input struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		input.Name = "John Doe"
		input.Email = "johndoe@example.com"
		input.Password = "john1234"

		res, err := tc.Post("/v1/signup", input)
		assert.Nil(t, err)

		assert.Equal(t, res.StatusCode, http.StatusForbidden)

		cookie := tc.GetCookie(res, "session_id")
		assert.True(t, cookie == nil)

		var output struct {
			Error string `json:"error"`
		}
		tc.ParseResponseBody(res, &output)

		assert.Equal(t, output.Error, "user already exists")

	})
}

func testLoggedInUser(t *testing.T, tc *TestClient) {
	t.Run("valid login", func(t *testing.T) {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		input.Email = "johndoe@example.com"
		input.Password = "john1234"

		res, err := tc.Post("/v1/login", input)
		assert.Nil(t, err)

		assert.Equal(t, res.StatusCode, http.StatusOK)

		cookie := tc.GetCookie(res, "session_id")
		assert.True(t, cookie.Value != "")

		var output struct {
			Message string `json:"message"`
		}
		tc.ParseResponseBody(res, &output)

		assert.Equal(t, output.Message, "logged in successfully")
	})
	t.Run("incorrect email", func(t *testing.T) {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		input.Email = "janesmith@example.com"
		input.Password = "john1234"

		res, err := tc.Post("/v1/login", input)
		assert.Nil(t, err)

		assert.Equal(t, res.StatusCode, http.StatusUnauthorized)

		cookie := tc.GetCookie(res, "session_id")
		assert.True(t, cookie == nil)

		var output struct {
			Error string `json:"error"`
		}
		tc.ParseResponseBody(res, &output)

		assert.Equal(t, output.Error, "Invalid Authentication Credentials")
	})
	t.Run("incorrect password", func(t *testing.T) {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		input.Email = "johndoe@example.com"
		input.Password = "john12345"

		res, err := tc.Post("/v1/login", input)
		assert.Nil(t, err)

		assert.Equal(t, res.StatusCode, http.StatusUnauthorized)

		cookie := tc.GetCookie(res, "session_id")
		assert.True(t, cookie == nil)

		var output struct {
			Error string `json:"error"`
		}
		tc.ParseResponseBody(res, &output)

		assert.Equal(t, output.Error, "Invalid Authentication Credentials")
	})

}
