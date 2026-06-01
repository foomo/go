package testing

import (
	"context"
	"testing"
	"time"

	gotime "github.com/foomo/go/time"
	"github.com/stretchr/testify/require"
)

// WaitFor repeatedly invokes cond until it returns true or the timeout elapses, failing the test on timeout.
func WaitFor(tb testing.TB, timeout time.Duration, cond func() bool) {
	tb.Helper()

	err := gotime.WaitFor(tb.Context(), func(ctx context.Context) (bool, error) {
		return cond(), nil
	}, timeout, 25*time.Millisecond)
	require.NoError(tb, err)
}
