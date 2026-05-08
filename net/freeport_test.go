package net_test

import (
	"context"
	stdnet "net"
	"testing"

	netx "github.com/foomo/go/net"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFreePort(t *testing.T) {
	t.Parallel()

	port, err := netx.FreePort(t.Context())
	require.NoError(t, err)
	assert.Positive(t, port)
}

func TestFreePort_CanceledContext(t *testing.T) {
	t.Skip("underlying Listen() does not support cancellation")
	t.Parallel()

	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	port, err := netx.FreePort(ctx)
	require.Error(t, err)
	assert.Zero(t, port)
}

func TestFreePorts(t *testing.T) {
	t.Parallel()

	ports, err := netx.FreePorts(t.Context(), 3)
	require.NoError(t, err)
	require.Len(t, ports, 3)

	seen := map[int]struct{}{}

	for _, p := range ports {
		assert.Positive(t, p)
		_, dup := seen[p]
		assert.False(t, dup, "duplicate port %d", p)
		seen[p] = struct{}{}
	}
}

func TestFreePorts_Zero(t *testing.T) {
	t.Parallel()

	ports, err := netx.FreePorts(t.Context(), 0)
	require.NoError(t, err)
	assert.Nil(t, ports)
}

func TestFreePorts_Negative(t *testing.T) {
	t.Parallel()

	ports, err := netx.FreePorts(t.Context(), -1)
	require.NoError(t, err)
	assert.Nil(t, ports)
}

func TestFreePorts_CanceledContext(t *testing.T) {
	t.Skip("underlying Listen() does not support cancellation")
	t.Parallel()

	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	ports, err := netx.FreePorts(ctx, 3)
	require.Error(t, err)
	assert.Nil(t, ports)
}

func TestIsFreePort(t *testing.T) {
	t.Parallel()

	port, err := netx.FreePort(t.Context())
	require.NoError(t, err)

	require.NoError(t, netx.IsFreePort(t.Context(), port))
}

func TestIsFreePort_InUse(t *testing.T) {
	t.Parallel()

	l, err := (&stdnet.ListenConfig{}).Listen(t.Context(), "tcp", "127.0.0.1:0")
	require.NoError(t, err)
	t.Cleanup(func() { _ = l.Close() })

	addr, ok := l.Addr().(*stdnet.TCPAddr)
	require.True(t, ok)
	require.Error(t, netx.IsFreePort(t.Context(), addr.Port))
}

func TestIsFreePort_CanceledContext(t *testing.T) {
	t.Skip("underlying Listen() does not support cancellation")
	t.Parallel()

	port, err := netx.FreePort(t.Context())
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(t.Context())
	cancel()

	require.Error(t, netx.IsFreePort(ctx, port))
}
