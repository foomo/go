package time

import (
	"context"
	"time"
)

// WaitFor repeatedly invokes the provided function until it returns true, an error, or the timeout duration elapses.
func WaitFor(ctx context.Context, fn func(context.Context) (bool, error), timeout, interval time.Duration) error {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		ok, err := fn(ctx)
		if ok {
			return err
		}

		if err := Sleep(ctx, interval); err != nil {
			return err
		}
	}

	return context.DeadlineExceeded
}
