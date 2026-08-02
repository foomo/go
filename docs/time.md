# time

Context-aware time utilities, duration parsing, and a controllable clock.

## Import

```go
import timex "github.com/foomo/go/time"
```

## Time control

`Now` is a package-level variable (defaulting to `time.Now`) that acts as the application's
clock. Call `timex.Now()` instead of the standard library `time.Now()` throughout your code,
then swap the provider to control time — for time travel or deterministic unit test outputs.

### Now

```go
var Now = time.Now
```

The current-time provider. Invoke `timex.Now()` wherever you would call `time.Now()`.
Reassign it (directly, or via `Static`/`Incremental`) to take control of the clock.

### Static

```go
func Static()
```

Points `Now` at a static provider: every call returns the same instant,
`time.Unix(0, NowStaticNSec)` (default `2021-01-01 12:00:00`). Ideal for fixed, reproducible
timestamps in tests and golden files.

### Incremental

```go
func Incremental()
```

Points `Now` at an incremental provider: each call returns a strictly increasing instant,
starting at `time.Unix(0, NowIncrementalNSec)` and advancing one nanosecond per call. Useful
when you need distinct but deterministic timestamps (e.g. stable ordering).

### ResetIncremental

```go
func ResetIncremental()
```

Rewinds the incremental provider's cursor (`NowIncrementalNSec`) back to `NowStaticNSec`.

### Tunables

```go
var NowStaticNSec      = int64(1609498800e9) // 2021-01-01 12:00:00 (11:00:00 UTC), in Unix nanoseconds
var NowIncrementalNSec = NowStaticNSec       // current incremental cursor
```

Adjust `NowStaticNSec` to change the fixed instant used by `Static` (it also seeds the
incremental cursor). `NowIncrementalNSec` holds the incremental provider's current position.

## Example

```go
// In production code, always read the clock through timex.Now.
func stamp() time.Time { return timex.Now() }

// In a test, freeze time for deterministic output.
timex.Static()
fmt.Println(stamp().UTC().Format(time.RFC3339)) // 2021-01-01T11:00:00Z
fmt.Println(stamp().UTC().Format(time.RFC3339)) // 2021-01-01T11:00:00Z (unchanged)

// Or advance deterministically, one nanosecond per call.
timex.Incremental()
a := stamp()
b := stamp()
fmt.Println(b.After(a)) // true

// Restore the default clock when done.
timex.Now = time.Now
```

## API

### Sleep

```go
func Sleep(ctx context.Context, d time.Duration) error
```

Waits for the specified duration or until the context is canceled, whichever occurs first. Returns the context error if canceled before the delay elapses.

### WaitFor

```go
func WaitFor(ctx context.Context, fn func(context.Context) (bool, error), timeout, interval time.Duration) error
```

Polls `fn` until it returns `true`, returns an error, or the timeout deadline elapses. Sleeps `interval` between attempts using context-aware `Sleep`, so a canceled context aborts the wait. Returns `context.DeadlineExceeded` if the deadline is reached without success.

### ParseDuration

```go
func ParseDuration(s string) (time.Duration, error)
```

Parses a duration string like `time.ParseDuration`, but also accepts the `d` (day, 24h) and `w` (week, 168h) units. Valid units: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`, `d`, `w`.

## Examples

### Sleep

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

// Wait for 1 second (completes successfully)
err := timex.Sleep(ctx, 1*time.Second)
if err != nil {
	fmt.Println("Sleep failed:", err)
	return
}

fmt.Println("Sleep completed successfully")

// Sleep for 3 seconds with a 2-second timeout (context cancels first)
ctx2, cancel2 := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel2()

err = timex.Sleep(ctx2, 3*time.Second)
if err != nil {
	fmt.Println("Sleep cancelled:", err)
}

// Output:
// Sleep completed successfully
// Sleep cancelled: context deadline exceeded
```

### WaitFor

```go
// Poll a readiness check up to 10s, every 250ms.
err := timex.WaitFor(ctx, func(ctx context.Context) (bool, error) {
	return service.Ready(ctx), nil
}, 10*time.Second, 250*time.Millisecond)
if err != nil {
	log.Fatal(err)
}
```

### ParseDuration

```go
for _, s := range []string{"2w", "5d", "1w2d3h"} {
	d, _ := timex.ParseDuration(s)
	fmt.Printf("%s = %s\n", s, d)
}

// Output:
// 2w = 336h0m0s
// 5d = 120h0m0s
// 1w2d3h = 219h0m0s
```
