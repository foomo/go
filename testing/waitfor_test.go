package testing_test

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	testingx "github.com/foomo/go/testing"
	"github.com/stretchr/testify/require"
)

func TestWaitFor(t *testing.T) {
	t.Parallel()

	var ready atomic.Bool

	go func() {
		time.Sleep(20 * time.Millisecond)
		ready.Store(true)
	}()

	testingx.WaitFor(t, time.Second, ready.Load)
	require.True(t, ready.Load())
}

func TestWaitFor_Timeout(t *testing.T) {
	t.Parallel()

	tb := testingx.NewExampleTB()
	testingx.WaitFor(tb, 50*time.Millisecond, func() bool { return false })
	require.True(t, tb.Failed())
}

func ExampleWaitFor() {
	tb := testingx.NewExampleTB()

	var ready atomic.Bool

	go func() {
		time.Sleep(20 * time.Millisecond)
		ready.Store(true)
	}()

	testingx.WaitFor(tb, time.Second, ready.Load)
	fmt.Println("ready:", ready.Load())
	// Output: ready: true
}
