package testing

import (
	"sync"
	"testing"
	"time"

	"github.com/foomo/go/net"
	"github.com/stretchr/testify/require"
)

// FreePort returns a free port on localhost, failing tb if none is available.
func FreePort(tb testing.TB) int {
	tb.Helper()

	port, err := net.FreePort(tb.Context())
	require.NoError(tb, err)

	return port
}

// FreePorts returns n free ports on localhost, failing tb if not all are available.
func FreePorts(tb testing.TB, n int) []int {
	tb.Helper()

	ports, err := net.FreePorts(tb.Context(), n)
	require.NoError(tb, err)

	return ports
}

// WaitForFreePorts blocks until each of ports becomes free on localhost,
// failing tb if any does not become free within 10 seconds.
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
