package testing

import (
	"sync"
	"testing"
	"time"

	"github.com/foomo/go/net"
	"github.com/stretchr/testify/require"
)

// FreePort returns a free port on localhost
func FreePort(tb testing.TB) int {
	tb.Helper()

	port, err := net.FreePort(tb.Context())
	require.NoError(tb, err)

	return port
}

// FreePorts returns a free ports on localhost
func FreePorts(tb testing.TB, n int) []int {
	tb.Helper()

	ports, err := net.FreePorts(tb.Context(), n)
	require.NoError(tb, err)

	return ports
}

// WaitForFreePorts returns a free port on localhost
func WaitForFreePorts(tb testing.TB, ports ...int) {
	tb.Helper()

	wg := sync.WaitGroup{}

	for _, port := range ports {
		wg.Add(1)

		go func(port int) {
			defer wg.Done()

			if err := net.WaitForFreePort(tb.Context(), port, 10*time.Second); err != nil {
				tb.Fatal(err)
			}
		}(port)
	}

	wg.Wait()
}
