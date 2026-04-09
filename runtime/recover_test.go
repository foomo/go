package runtime_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/foomo/go/runtime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecover(t *testing.T) {
	t.Parallel()

	t.Run("no panic", func(t *testing.T) {
		t.Parallel()

		err := runtime.Recover(func() {
			// no panic
		})

		assert.NoError(t, err)
	})

	t.Run("string panic", func(t *testing.T) {
		t.Parallel()

		err := runtime.Recover(func() {
			panic("something went wrong")
		})

		require.Error(t, err)
		assert.Contains(t, err.Error(), "panic: something went wrong")

		var pe *runtime.PanicError
		require.ErrorAs(t, err, &pe)
		assert.Equal(t, "something went wrong", pe.Value)
		require.NoError(t, pe.Unwrap())
		assert.NotEmpty(t, pe.Stack)
	})

	t.Run("error panic", func(t *testing.T) {
		t.Parallel()

		target := errors.New("original error")
		err := runtime.Recover(func() {
			panic(target)
		})

		require.Error(t, err)
		require.ErrorIs(t, err, target)

		var pe *runtime.PanicError
		require.ErrorAs(t, err, &pe)
		assert.Equal(t, target, pe.Value)
		assert.Equal(t, target, pe.Unwrap())
	})

	t.Run("nil panic", func(t *testing.T) {
		t.Parallel()

		err := runtime.Recover(func() {
			panic(nil)
		})

		// Go 1.21+ wraps panic(nil) as *runtime.PanicNilError
		require.Error(t, err)

		var pe *runtime.PanicError
		require.ErrorAs(t, err, &pe)
		assert.NotNil(t, pe.Value)
	})

	t.Run("stack trace", func(t *testing.T) {
		t.Parallel()

		err := runtime.Recover(func() {
			panic("boom")
		})

		var pe *runtime.PanicError
		require.ErrorAs(t, err, &pe)
		assert.Contains(t, pe.Stack, "runtime_test.TestRecover")
	})
}

func ExampleRecover() {
	err := runtime.Recover(func() {
		panic("something went wrong")
	})

	var pe *runtime.PanicError
	if errors.As(err, &pe) {
		fmt.Println(pe.Value)
	}

	// Output:
	// something went wrong
}
