package httpport_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ahmadabdelrazik/masarak/config"
	"github.com/ahmadabdelrazik/masarak/internal/adapter/memory"
	"github.com/ahmadabdelrazik/masarak/internal/core/app"
	"github.com/ahmadabdelrazik/masarak/internal/core/httpport"
	"github.com/ahmadabdelrazik/masarak/internal/core/httpport/auth"
	"github.com/ahmadabdelrazik/masarak/pkg/assert"
)

type TestClient struct {
	server *httptest.Server
}

func NewTestClient(t *testing.T) *TestClient {
	cfg, err := config.Load(".env.test")
	assert.Nil(t, err)
	mem := memory.NewMemory()
	repos := memory.NewInMemoryRepositories(mem)
	tokens := memory.NewInMemoryTokenRepository(mem, repos.AuthUsers)
	application := app.NewApplication(repos)
	authService := auth.New(tokens, repos.AuthUsers)
	server := httpport.NewHttpServer(application, cfg, authService, repos.AuthUsers)
	return &TestClient{
		server: httptest.NewServer(server.Routes()),
	}
}

func TestServer_Healthy(t *testing.T) {
	t.Parallel()

	tc := NewTestClient(t)
	defer tc.Close()

	res, err := tc.Get("/v1/health")
	assert.Nil(t, err)
	response, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, string(response), "Healthy\n")
}

func (c *TestClient) Close() {
	c.server.Close()
}

func (c *TestClient) Get(endpoint string) (*http.Response, error) {
	return http.Get(c.server.URL + endpoint)
}

func (c *TestClient) Post(endpoint string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return http.Post(c.server.URL+endpoint, "application/json", bytes.NewBuffer(jsonBody))
}

func (c *TestClient) Put(endpoint string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, c.server.URL+endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(req)
}

func (c *TestClient) Patch(endpoint string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, c.server.URL+endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return http.DefaultClient.Do(req)
}

func (c *TestClient) Delete(endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, c.server.URL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}

func (c *TestClient) ParseResponseBody(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

func (c *TestClient) GetResponseString(resp *http.Response) (string, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (c *TestClient) GetCookie(res *http.Response, cookieName string) *http.Cookie {
	for _, c := range res.Cookies() {
		if c.Name == cookieName {
			return c
		}
	}

	return nil
}

func (c *TestClient) Do(method, endpoint string, body any, cookies ...*http.Cookie) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, c.server.URL+endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	for _, c := range cookies {
		req.AddCookie(c)
	}

	return http.DefaultClient.Do(req)
}

func (c *TestClient) GetWithCookies(endpoint string, cookies ...*http.Cookie) (*http.Response, error) {
	return c.Do(http.MethodGet, endpoint, nil, cookies...)
}
func (c *TestClient) PostWithCookies(endpoint string, body any, cookies ...*http.Cookie) (*http.Response, error) {
	return c.Do(http.MethodPost, endpoint, body, cookies...)
}

func (c *TestClient) PatchWithCookies(endpoint string, body any, cookies ...*http.Cookie) (*http.Response, error) {
	return c.Do(http.MethodPatch, endpoint, body, cookies...)
}

func (c *TestClient) DeleteWithCookies(endpoint string, body any, cookies ...*http.Cookie) (*http.Response, error) {
	return c.Do(http.MethodDelete, endpoint, body, cookies...)
}
