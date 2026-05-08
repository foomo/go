package net

import (
	"context"
	"net"
	"time"

	timex "github.com/foomo/go/time"
)

// WaitFor attempts to establish a connection to the specified address until the timeout or context cancellation occurs.
func WaitFor(ctx context.Context, network, address string, timeout time.Duration) error {
	return timex.WaitFor(ctx, func(ctx context.Context) (bool, error) {
		conn, err := (&net.Dialer{Timeout: time.Second}).DialContext(ctx, network, address)
		if err != nil {
			return false, err
		}

		_ = conn.Close()

		return true, nil
	}, timeout, 100*time.Millisecond)
}
