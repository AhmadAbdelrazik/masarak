package httpport_test

import (
	"testing"
)

func TestServer_RegisterOwner(t *testing.T) {
	t.Parallel()

	tc := NewTestClient(t)
	defer tc.Close()

}
