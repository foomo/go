package os_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/foomo/go/os"
)

func TestExists(t *testing.T) {
	assert.True(t, os.Exists("GOROOT"))
	assert.False(t, os.Exists("UNDEFINED"))
}

func TestMustExists(t *testing.T) {
	assert.Panics(t, func() {
		os.MustExists("UNDEFINED")
	})
	assert.NotPanics(t, func() {
		os.Exists("GOROOT")
	})
}
