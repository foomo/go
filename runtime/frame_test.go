package runtime_test

import (
	"testing"

	"github.com/foomo/go/runtime"
	"github.com/stretchr/testify/assert"
)

func TestFrameName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		frame runtime.Frame
		want  string
	}{
		{"no pkg", runtime.Frame{Func: "main"}, "main"},
		{"package func", runtime.Frame{Pkg: "pkg", Func: "Process"}, "pkg.Process"},
		{"method", runtime.Frame{Pkg: "pkg", Inst: "Payment", Func: "Process"}, "pkg.Payment.Process"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.frame.Name())
		})
	}
}

func TestFrameShort(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "Process", runtime.Frame{Pkg: "pkg", Func: "Process"}.Short())
	assert.Equal(t, "Payment.Process", runtime.Frame{Pkg: "pkg", Inst: "Payment", Func: "Process"}.Short())
}

func TestFrameZero(t *testing.T) {
	t.Parallel()

	assert.True(t, runtime.Frame{}.Zero())
	assert.False(t, runtime.Frame{Func: "x"}.Zero())
	assert.False(t, runtime.Frame{Pkg: "pkg"}.Zero())
}
