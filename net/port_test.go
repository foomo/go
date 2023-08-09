package net_test

import (
	"testing"

	netx "github.com/foomo/go/net"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFreePort(t *testing.T) {
	port, err := netx.FreePort()
	require.NoError(t, err)
	assert.NotEqual(t, 0, port)
}
