package os_test

import (
	"os"
	"testing"

	osx "github.com/foomo/go/os"

	"github.com/stretchr/testify/assert"
)

func TestGetenv(t *testing.T) {
	_ = os.Unsetenv("FOO")
	assert.Equal(t, "foo", osx.Getenv("FOO", "foo"))
	_ = os.Setenv("FOO", "bar")
	assert.Equal(t, "bar", osx.Getenv("FOO", "foo"))
}

func TestGetenvBool(t *testing.T) {
	_ = os.Unsetenv("FOO")
	if v, err := osx.GetenvBool("FOO", false); assert.NoError(t, err) {
		assert.Equal(t, false, v)
	}

	_ = os.Setenv("FOO", "true")
	if v, err := osx.GetenvBool("FOO", false); assert.NoError(t, err) {
		assert.Equal(t, true, v)
	}
}

func TestGetenvFloat(t *testing.T) {
	_ = os.Unsetenv("FOO")
	if v, err := osx.GetenvFloat("FOO", 0.1); assert.NoError(t, err) {
		assert.Equal(t, 0.1, v)
	}

	_ = os.Setenv("FOO", "0.2")
	if v, err := osx.GetenvFloat("FOO", 0.1); assert.NoError(t, err) {
		assert.Equal(t, 0.2, v)
	}
}

func TestGetenvInt(t *testing.T) {
	_ = os.Unsetenv("FOO")
	if v, err := osx.GetenvInt("FOO", 1); assert.NoError(t, err) {
		assert.Equal(t, int64(1), v)
	}

	_ = os.Setenv("FOO", "2")
	if v, err := osx.GetenvInt("FOO", 1); assert.NoError(t, err) {
		assert.Equal(t, int64(2), v)
	}
}

func TestGetenvStrings(t *testing.T) {
	_ = os.Unsetenv("FOO")
	assert.Nil(t, osx.GetenvStrings("FOO", nil))

	_ = os.Setenv("FOO", "foo")
	assert.Equal(t, []string{"foo"}, osx.GetenvStrings("FOO", nil))

	_ = os.Setenv("FOO", "foo,bar")
	assert.Equal(t, []string{"foo", "bar"}, osx.GetenvStrings("FOO", nil))
}
