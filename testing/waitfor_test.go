package testing_test

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	gotesting "github.com/foomo/go/testing"
	"github.com/stretchr/testify/require"
)

func TestWaitFor(t *testing.T) {
	t.Parallel()

	var ready atomic.Bool

	go func() {
		time.Sleep(20 * time.Millisecond)
		ready.Store(true)
	}()

	gotesting.WaitFor(t, time.Second, ready.Load)
	require.True(t, ready.Load())
}

func TestWaitFor_Timeout(t *testing.T) {
	t.Parallel()

	tb := gotesting.NewExampleTB()
	gotesting.WaitFor(tb, 50*time.Millisecond, func() bool { return false })
	require.True(t, tb.Failed())
}

func ExampleWaitFor() {
	tb := gotesting.NewExampleTB()

	var ready atomic.Bool

	go func() {
		time.Sleep(20 * time.Millisecond)
		ready.Store(true)
	}()

	gotesting.WaitFor(tb, time.Second, ready.Load)
	fmt.Println("ready:", ready.Load())
	// Output: ready: true
}
