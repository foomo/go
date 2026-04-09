# time

Context-aware time utilities.

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

## Example

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
