package testing

import (
	"testing"

	"github.com/foomo/go/net"
	"github.com/stretchr/testify/require"
)

// FreePort returns a free port on localhost
func FreePort(t *testing.T) string {
	t.Helper()

	addr, err := net.FreePort()
	require.NoError(t, err)

	return addr
}
