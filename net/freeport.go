package net

import (
	"context"
	"errors"
	"net"
	"strconv"
	"time"
)

// FreePort returns a port on localhost that is currently unused, by briefly
// binding to port 0 and releasing it. The port can be claimed by another
// process before the caller uses it.
func FreePort(ctx context.Context) (int, error) {
	ports, err := FreePorts(ctx, 1)
	if err != nil {
		return 0, err
	}

	return ports[0], nil
}

// FreePorts returns n distinct ports on localhost that are currently unused,
// by briefly binding to n separate port-0 listeners and releasing them all.
// It returns nil, nil if n <= 0. The ports can be claimed by another process
// before the caller uses them.
func FreePorts(ctx context.Context, n int) ([]int, error) {
	if n <= 0 {
		return nil, nil
	}

	ls := make([]net.Listener, n)

	defer func() {
		for _, l := range ls {
			if l != nil {
				_ = l.Close()
			}
		}
	}()

	for i := range n {
		l, err := (&net.ListenConfig{}).Listen(ctx, "tcp", "127.0.0.1:0")
		if err != nil {
			return nil, err
		}

		ls[i] = l
	}

	ports := make([]int, n)

	for i, l := range ls {
		addr, ok := l.Addr().(*net.TCPAddr)
		if !ok {
			return nil, errors.New("unexpected listener address type")
		}

		ports[i] = addr.Port
	}

	return ports, nil
}

// IsFreePort reports whether port is currently unused on localhost, by
// attempting to bind to it and immediately releasing it. It returns a non-nil
// error if the port is already in use or cannot be bound.
func IsFreePort(ctx context.Context, port int) error {
	addr := net.JoinHostPort("127.0.0.1", strconv.Itoa(port))

	l, err := (&net.ListenConfig{}).Listen(ctx, "tcp", addr)
	if err != nil {
		return err
	}

	defer func() {
		_ = l.Close()
	}()

	return nil
}

// WaitForFreePort attempts to establish a TCP connection to the specified address until the timeout or context cancellation occurs.
func WaitForFreePort(ctx context.Context, port int, timeout time.Duration) error {
	addr := net.JoinHostPort("127.0.0.1", strconv.Itoa(port))
	return WaitFor(ctx, "tcp", addr, timeout)
}
