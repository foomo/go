package time_test

import (
	"fmt"
	"testing"
	"time"

	timex "github.com/foomo/go/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Restore captures the mutable package globals and reinstates them on cleanup so
// that controlling the clock in one test does not leak into sibling tests. The
// globals are shared state, so these tests must not run in parallel.
func restore(t *testing.T) {
	t.Helper()

	now := timex.Now
	static := timex.NowStaticNSec
	incremental := timex.NowIncrementalNSec

	t.Cleanup(func() {
		timex.Now = now
		timex.NowStaticNSec = static
		timex.NowIncrementalNSec = incremental
	})
}

func TestStatic(t *testing.T) {
	restore(t)

	timex.Static()

	want := time.Unix(0, timex.NowStaticNSec)

	first := timex.Now()
	second := timex.Now()

	assert.Equal(t, want, first)
	// Every call returns the very same instant.
	assert.Equal(t, first, second)
}

func TestIncremental(t *testing.T) {
	restore(t)

	timex.Incremental()

	first := timex.Now()
	require.Equal(t, time.Unix(0, timex.NowStaticNSec), first)

	second := timex.Now()
	// Strictly increasing by one nanosecond per call.
	assert.Equal(t, first.Add(time.Nanosecond), second)
	assert.True(t, second.After(first))
}

func TestResetIncremental(t *testing.T) {
	restore(t)

	timex.Incremental()
	timex.Now()
	timex.Now()
	require.Greater(t, timex.NowIncrementalNSec, timex.NowStaticNSec)

	timex.ResetIncremental()
	assert.Equal(t, timex.NowStaticNSec, timex.NowIncrementalNSec)
}

func TestNowDefault(t *testing.T) {
	restore(t)

	// The default provider is a live clock.
	assert.WithinDuration(t, time.Now(), timex.Now(), time.Minute)
}

func ExampleStatic() {
	// Save and restore the default clock so the example is self-contained.
	defer func() { timex.Now = time.Now }()

	timex.Static()

	fmt.Println(timex.Now().UTC().Format(time.RFC3339))
	fmt.Println(timex.Now().UTC().Format(time.RFC3339))

	// Output:
	// 2021-01-01T11:00:00Z
	// 2021-01-01T11:00:00Z
}
