package os_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	osx "github.com/foomo/go/os"
)

func TestExists(t *testing.T) {
	require.NoError(t, os.Unsetenv("FOO"))
	assert.False(t, osx.Exists("FOO"))

	require.NoError(t, os.Setenv("FOO", "bar"))
	assert.True(t, osx.Exists("FOO"))
}

func TestMustExists(t *testing.T) {
	require.NoError(t, os.Unsetenv("FOO"))
	assert.Panics(t, func() {
		osx.MustExists("FOO")
	})

	require.NoError(t, os.Setenv("FOO", "bar"))
	assert.NotPanics(t, func() {
		osx.Exists("FOO")
	})
}
