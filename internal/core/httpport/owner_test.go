package httpport_test

import (
	"io"
	"testing"

	"github.com/ahmadabdelrazik/masarak/pkg/assert"
)

func TestServer_Healthy(t *testing.T) {
	tc := NewTestClient(t)
	defer tc.Close()

	res, err := tc.Get("/v1/health")
	assert.Nil(t, err)
	response, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, string(response), "Healthy\n")
}
