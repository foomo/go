package testing

import (
	"testing"

	"github.com/foomo/go/net"
	"github.com/stretchr/testify/require"
)

// FreePort returns a free port on localhost
func FreePort(tb testing.TB) string {
	tb.Helper()

	addr, err := net.FreePort()
	require.NoError(tb, err)

	return addr
}
