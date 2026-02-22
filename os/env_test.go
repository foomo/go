package os_test

import (
	"os"
	"testing"

	osx "github.com/foomo/go/os"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestEnvExists(t *testing.T) {
	require.NoError(t, os.Unsetenv("FOO"))
	assert.False(t, osx.HasEnv("FOO"))

	t.Setenv("FOO", "bar")
	assert.True(t, osx.HasEnv("FOO"))
}

func TestMustEnvExists(t *testing.T) {
	require.NoError(t, os.Unsetenv("FOO"))
	assert.Panics(t, func() {
		osx.MustHasEnv("FOO")
	})

	t.Setenv("FOO", "bar")
	assert.NotPanics(t, func() {
		osx.HasEnv("FOO")
	})
}

func TestGetenv(t *testing.T) {
	t.Setenv("FOO", "")
	assert.Equal(t, "foo", osx.Getenv("FOO", "foo"))
	t.Setenv("FOO", "bar")
	assert.Equal(t, "bar", osx.Getenv("FOO", "foo"))
}

func TestGetenvBool(t *testing.T) {
	t.Setenv("FOO", "")

	if v, err := osx.GetenvBool("FOO", false); assert.NoError(t, err) {
		assert.False(t, v)
	}

	t.Setenv("FOO", "true")

	if v, err := osx.GetenvBool("FOO", false); assert.NoError(t, err) {
		assert.True(t, v)
	}
}

func TestGetenvInt64(t *testing.T) {
	t.Setenv("FOO", "")

	if v, err := osx.GetenvInt64("FOO", 1); assert.NoError(t, err) {
		assert.Equal(t, int64(1), v)
	}

	t.Setenv("FOO", "2")

	if v, err := osx.GetenvInt64("FOO", 1); assert.NoError(t, err) {
		assert.Equal(t, int64(2), v)
	}
}

func TestGetenvFloat64(t *testing.T) {
	t.Setenv("FOO", "")

	if v, err := osx.GetenvFloat64("FOO", 0.1); assert.NoError(t, err) {
		assert.Equal(t, 0.1, v)
	}

	t.Setenv("FOO", "0.2")

	if v, err := osx.GetenvFloat64("FOO", 0.1); assert.NoError(t, err) {
		assert.Equal(t, 0.2, v)
	}
}

func TestGetenvStringSlice(t *testing.T) {
	t.Setenv("FOO", "")
	assert.Nil(t, osx.GetenvStringSlice("FOO", nil))

	t.Setenv("FOO", "foo")
	assert.Equal(t, []string{"foo"}, osx.GetenvStringSlice("FOO", nil))

	t.Setenv("FOO", "foo,bar")
	assert.Equal(t, []string{"foo", "bar"}, osx.GetenvStringSlice("FOO", nil))
}
