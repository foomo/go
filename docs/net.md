# net

Network utilities — free port allocation and TCP connection wait helpers.

## Import

```go
import netx "github.com/foomo/go/net"
```

## API

### FreePort

```go
func FreePort(ctx context.Context) (int, error)
```

Returns a free port on localhost. Useful for dynamically allocating ports in tests or local development.

### FreePorts

```go
func FreePorts(ctx context.Context, n int) ([]int, error)
```

Returns `n` free ports on localhost. All listeners are held open until the function returns, so the returned ports are guaranteed not to collide with each other.

### IsFreePort

```go
func IsFreePort(ctx context.Context, port int) error
```

Returns `nil` if the given port can be bound on `127.0.0.1`, otherwise returns the underlying `Listen` error.

### WaitFor

```go
func WaitFor(ctx context.Context, network, address string, timeout time.Duration) error
```

Repeatedly attempts to dial `network`/`address` until a connection succeeds, the context is canceled, or the timeout elapses. Polls every 100ms; each individual dial uses a 1s timeout. Returns `context.DeadlineExceeded` on timeout.

### WaitForFreePort

```go
func WaitForFreePort(ctx context.Context, port int, timeout time.Duration) error
```

Convenience wrapper around `WaitFor` that targets `127.0.0.1:port` over TCP.

## Example

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	netx "github.com/foomo/go/net"
)

func main() {
	ctx := context.Background()

	port, err := netx.FreePort(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(port) // e.g. 54321

	// Wait for a service to come up on the port (max 5s).
	if err := netx.WaitForFreePort(ctx, port, 5*time.Second); err != nil {
		log.Fatal(err)
	}
}
```
