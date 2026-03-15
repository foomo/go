# option

Functional options pattern support with generics.

## Import

```go
import "github.com/foomo/go/option"
```

## Types

```go
type Option[T any] func(T)
type OptionE[T any] func(T) error
```

## API

### Apply

```go
func Apply[T any](v T, opts ...Option[T])
```

Applies each option function to the value.

### ApplyE

```go
func ApplyE[T any](v T, opts ...OptionE[T]) error
```

Applies each option function to the value, stopping and returning on the first error.

## Example

```go
package main

import (
	"fmt"

	"github.com/foomo/go/option"
)

type Server struct {
	Name string
	Port int
}

func WithName(name string) option.Option[*Server] {
	return func(s *Server) {
		s.Name = name
	}
}

func WithPort(port int) option.Option[*Server] {
	return func(s *Server) {
		s.Port = port
	}
}

func main() {
	s := &Server{}
	option.Apply(s, WithName("localhost"), WithPort(8080))
	fmt.Println(s.Name) // localhost
	fmt.Println(s.Port) // 8080
}
```

### Error handling with ApplyE

```go
func WithValidatedPort(port int) option.OptionE[*Server] {
	return func(s *Server) error {
		if port < 1 || port > 65535 {
			return fmt.Errorf("invalid port: %d", port)
		}
		s.Port = port
		return nil
	}
}

err := option.ApplyE(s, WithValidatedPort(99999))
// err: invalid port: 99999
```
