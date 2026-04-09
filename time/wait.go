package time

import (
	"context"
	"time"
)

// Sleep waits for the specified delay duration or until the context is canceled, whichever occurs first.
// Returns an error if the context is canceled before the delay elapses.
func Sleep(ctx context.Context, d time.Duration) error {
	t := time.NewTimer(d)
	select {
	case <-ctx.Done():
		t.Stop()
		return ctx.Err()
	case <-t.C:
		return nil
	}
}
