package net_test

import (
	"context"
	stdnet "net"
	"strconv"
	"testing"
	"time"

	netx "github.com/foomo/go/net"
	"github.com/stretchr/testify/require"
)

func TestWaitFor(t *testing.T) {
	t.Parallel()

	l, err := (&stdnet.ListenConfig{}).Listen(t.Context(), "tcp", "127.0.0.1:0")
	require.NoError(t, err)

	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			conn, err := l.Accept()
			if err != nil {
				return
			}

			_ = conn.Close()
		}
	}()

	t.Cleanup(func() {
		_ = l.Close()

		<-done
	})

	require.NoError(t, netx.WaitFor(t.Context(), "tcp", l.Addr().String(), time.Second))
}

func TestWaitFor_Timeout(t *testing.T) {
	t.Parallel()

	port, err := netx.FreePort(t.Context())
	require.NoError(t, err)

	addr := stdnet.JoinHostPort("127.0.0.1", strconv.Itoa(port))

	err = netx.WaitFor(t.Context(), "tcp", addr, 200*time.Millisecond)
	require.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestWaitFor_CanceledContext(t *testing.T) {
	t.Parallel()

	port, err := netx.FreePort(t.Context())
	require.NoError(t, err)

	addr := stdnet.JoinHostPort("127.0.0.1", strconv.Itoa(port))

	ctx, cancel := context.WithCancel(t.Context())

	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	err = netx.WaitFor(ctx, "tcp", addr, time.Second)
	require.ErrorIs(t, err, context.Canceled)
}
