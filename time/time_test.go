package time_test

import (
	"fmt"
	"testing"
	"time"

	gotime "github.com/foomo/go/time"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Restore captures the mutable package globals and reinstates them on cleanup so
// that controlling the clock in one test does not leak into sibling tests. The
// globals are shared state, so these tests must not run in parallel.
func restore(t *testing.T) {
	t.Helper()

	now := gotime.Now
	static := gotime.NowStaticNSec
	incremental := gotime.NowIncrementalNSec

	t.Cleanup(func() {
		gotime.Now = now
		gotime.NowStaticNSec = static
		gotime.NowIncrementalNSec = incremental
	})
}

func TestStatic(t *testing.T) {
	restore(t)

	gotime.Static()

	want := time.Unix(0, gotime.NowStaticNSec)

	first := gotime.Now()
	second := gotime.Now()

	assert.Equal(t, want, first)
	// Every call returns the very same instant.
	assert.Equal(t, first, second)
}

func TestIncremental(t *testing.T) {
	restore(t)

	gotime.Incremental()

	first := gotime.Now()
	require.Equal(t, time.Unix(0, gotime.NowStaticNSec), first)

	second := gotime.Now()
	// Strictly increasing by one nanosecond per call.
	assert.Equal(t, first.Add(time.Nanosecond), second)
	assert.True(t, second.After(first))
}

func TestResetIncremental(t *testing.T) {
	restore(t)

	gotime.Incremental()
	gotime.Now()
	gotime.Now()
	require.Greater(t, gotime.NowIncrementalNSec, gotime.NowStaticNSec)

	gotime.ResetIncremental()
	assert.Equal(t, gotime.NowStaticNSec, gotime.NowIncrementalNSec)
}

func TestNowDefault(t *testing.T) {
	restore(t)

	// The default provider is a live clock.
	assert.WithinDuration(t, time.Now(), gotime.Now(), time.Minute)
}

func ExampleStatic() {
	// Save and restore the default clock so the example is self-contained.
	defer func() { gotime.Now = time.Now }()

	gotime.Static()

	fmt.Println(gotime.Now().UTC().Format(time.RFC3339))
	fmt.Println(gotime.Now().UTC().Format(time.RFC3339))

	// Output:
	// 2021-01-01T11:00:00Z
	// 2021-01-01T11:00:00Z
}
