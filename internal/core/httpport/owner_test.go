package httpport_test

import (
	"testing"
)

func TestServer_RegisterOwner(t *testing.T) {
	t.Parallel()

	tc := NewTestClient(t)
	defer tc.Close()

	t.Run("post owner", func(t *testing.T) {

	})
}

func addUser(t *testing.T, tc *TestClient) {

}

func testPostOwner(t *testing.T, tc *TestClient) {

}
