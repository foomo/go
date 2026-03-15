# net

Network utilities.

## Import

```go
import netx "github.com/foomo/go/net"
```

## API

### FreePort

```go
func FreePort() (string, error)
```

Returns a free port on localhost in `address:port` format. Useful for dynamically allocating ports in tests or local development.

## Example

```go
package main

import (
	"fmt"
	"log"

	netx "github.com/foomo/go/net"
)

func main() {
	addr, err := netx.FreePort()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addr) // e.g. "127.0.0.1:54321"
}
```
