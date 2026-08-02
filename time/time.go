package time

import (
	"time"
)

var (
	// Now returns the current time. It is a package-level variable so it can be
	// swapped out to control time, e.g., for time travel or deterministic unit
	// test outputs. Call timex.Now() instead of the standard library time.Now(),
	// then use Static or Incremental to take control of the clock.
	//
	// It defaults to the standard library time.Now.
	Now = time.Now

	// NowStaticNSec is the fixed instant, in Unix nanoseconds, returned by the
	// static provider installed via Static. It also seeds NowIncrementalNSec.
	// Defaults to 2021-01-01 12:00:00 (11:00:00 UTC).
	NowStaticNSec = int64(1609498800e9) // 2021-01-01 12:00:00

	// NowIncrementalNSec is the current cursor, in Unix nanoseconds, used by the
	// incremental provider installed via Incremental. It advances by one
	// nanosecond on every call to Now and can be rewound with ResetIncremental.
	NowIncrementalNSec = NowStaticNSec
)

// Static points Now at a static time provider: every call to Now returns the
// same instant, time.Unix(0, NowStaticNSec).
func Static() {
	Now = static
}

// Incremental points Now at an incremental time provider: each call to Now
// returns a strictly increasing instant, starting at time.Unix(0,
// NowIncrementalNSec) and advancing by one nanosecond per call.
func Incremental() {
	Now = incremental
}

func static() time.Time {
	return time.Unix(0, NowStaticNSec)
}

func incremental() time.Time {
	t := time.Unix(0, NowIncrementalNSec)
	NowIncrementalNSec++

	return t
}

// ResetIncremental rewinds the incremental provider's cursor back to the static
// default (NowStaticNSec).
func ResetIncremental() {
	NowIncrementalNSec = NowStaticNSec
}
