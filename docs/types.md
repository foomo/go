# types

Common interface contracts and function adapters used across the foomo ecosystem. Each lifecycle interface ships with a small family of adapters so plain functions can satisfy it without writing wrapper structs.

## Import

```go
import "github.com/foomo/go/types"
```

## Pattern

For each context-aware interface (e.g. `Starter`), the package provides:

- `<Name>Func`         — adapts `func()`
- `<Name>FuncErr`      — adapts `func() error`
- `<Name>FuncCtx`      — adapts `func(context.Context)`
- `<Name>FuncCtxErr`   — adapts `func(context.Context) error`
- `As<Name>(v any) (<Name>, bool)` — runtime type assertion that promotes any of the above (or an existing `<Name>`) into the interface

The function adapters never return errors when the underlying function has none, so wrapping a `func()` in `<Name>Func` is always safe.

## Interfaces

### Closer

```go
type Closer interface {
	Close(ctx context.Context) error
}
```

Adapters: `CloseFunc`, `CloseFuncErr`, `CloseFuncCtx`, `CloseFuncCtxErr`. Helper: `AsCloser`.

### Pinger

```go
type Pinger interface {
	Ping(ctx context.Context) error
}
```

Adapters: `PingFunc`, `PingFuncErr`, `PingFuncCtx`, `PingFuncCtxErr`. Helper: `AsPinger`.

### Shutdowner

```go
type Shutdowner interface {
	Shutdown(ctx context.Context) error
}
```

Adapters: `ShutdownFunc`, `ShutdownFuncErr`, `ShutdownFuncCtx`, `ShutdownFuncCtxErr`. Helper: `AsShutdowner`.

### Starter

```go
type Starter interface {
	Start(ctx context.Context) error
}
```

Adapters: `StartFunc`, `StartFuncErr`, `StartFuncCtx`, `StartFuncCtxErr`. Helper: `AsStarter`.

### Stopper

```go
type Stopper interface {
	Stop(ctx context.Context) error
}
```

Adapters: `StopFunc`, `StopFuncErr`, `StopFuncCtx`, `StopFuncCtxErr`. Helper: `AsStopper`.

### Unsubscriber

```go
type Unsubscriber interface {
	Unsubscribe(ctx context.Context) error
}
```

Adapters: `UnsubscribeFunc`, `UnsubscribeFuncErr`, `UnsubscribeFuncCtx`, `UnsubscribeFuncCtxErr`. Helper: `AsUnsubscriber`.

### Namer

```go
type Namer interface {
	Name() string
}
```

### Stringer

```go
type Stringer interface {
	String() string
}
```

## Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/foomo/go/types"
)

func main() {
	// Any of these forms can be promoted to a Starter.
	candidates := []any{
		func() { fmt.Println("plain") },
		func() error { fmt.Println("err"); return nil },
		func(ctx context.Context) { fmt.Println("ctx") },
		func(ctx context.Context) error { fmt.Println("ctx+err"); return nil },
	}

	for _, c := range candidates {
		s, ok := types.AsStarter(c)
		if !ok {
			continue
		}
		_ = s.Start(context.Background())
	}
}
```
