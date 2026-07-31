package time_test

import (
	"context"
	"errors"
	"testing"
	"time"

	timex "github.com/foomo/go/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWaitFor(t *testing.T) {
	t.Parallel()

	t.Run("success on first poll", func(t *testing.T) {
		t.Parallel()

		require.NoError(t, timex.WaitFor(t.Context(), func(context.Context) (bool, error) {
			return true, nil
		}, time.Second, 10*time.Millisecond))
	})

	t.Run("error passthrough when done", func(t *testing.T) {
		t.Parallel()

		errBoom := errors.New("boom")
		err := timex.WaitFor(t.Context(), func(context.Context) (bool, error) {
			return true, errBoom
		}, time.Second, 10*time.Millisecond)
		require.ErrorIs(t, err, errBoom)
	})

	t.Run("becomes true after retries", func(t *testing.T) {
		t.Parallel()

		calls := 0

		require.NoError(t, timex.WaitFor(t.Context(), func(context.Context) (bool, error) {
			calls++
			return calls >= 3, nil
		}, time.Second, 5*time.Millisecond))
		assert.GreaterOrEqual(t, calls, 3)
	})

	t.Run("timeout", func(t *testing.T) {
		t.Parallel()

		err := timex.WaitFor(t.Context(), func(context.Context) (bool, error) {
			return false, nil
		}, 100*time.Millisecond, 10*time.Millisecond)
		require.ErrorIs(t, err, context.DeadlineExceeded)
	})

	t.Run("canceled context", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(t.Context())
		cancel()

		err := timex.WaitFor(ctx, func(context.Context) (bool, error) {
			return false, nil
		}, time.Second, 10*time.Millisecond)
		require.ErrorIs(t, err, context.Canceled)
	})
}
