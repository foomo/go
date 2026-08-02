package os_test

import (
	"os"
	"path/filepath"
	"testing"

	goos "github.com/foomo/go/os"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpand(t *testing.T) {
	t.Run("home prefix", func(t *testing.T) {
		home, err := os.UserHomeDir()
		require.NoError(t, err)

		got, err := goos.Expand("~/foo/bar")
		require.NoError(t, err)
		assert.Equal(t, filepath.Join(home, "foo", "bar"), got)
	})

	t.Run("env var", func(t *testing.T) {
		t.Setenv("EXPAND_TEST_VAR", "value")

		got, err := goos.Expand("prefix-$EXPAND_TEST_VAR-${EXPAND_TEST_VAR}")
		require.NoError(t, err)
		assert.Equal(t, "prefix-value-value", got)
	})

	t.Run("plain passthrough", func(t *testing.T) {
		got, err := goos.Expand("/absolute/path")
		require.NoError(t, err)
		assert.Equal(t, "/absolute/path", got)
	})
}
