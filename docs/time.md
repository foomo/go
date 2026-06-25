# time

Context-aware time utilities and duration parsing.

## Import

```go
import timex "github.com/foomo/go/time"
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
