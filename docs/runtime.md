# runtime

Runtime introspection utilities.

## Import

```go
import runtimex "github.com/foomo/go/runtime"
```

## API

### Caller

```go
func Caller(skip int) (shortName, fullName, file string, line int, ok bool)
```

Enriched wrapper around `runtime.Caller`. Returns the short function name (e.g. `pkg.Func`), the fully qualified name, a simplified file path (`dir/file.go`), and the line number. The `skip` parameter works like `runtime.Caller` — use `0` for the immediate caller.

### CallerFunc

```go
func CallerFunc(skip int) (string, bool)
```

Returns only the bare function name (without the package path or receiver) for the caller `skip` frames up the call stack. Returns `"unknown", false` if the caller cannot be determined.

### StackTrace

```go
func StackTrace(size, skip int) string
```

Captures and formats a stack trace of up to `size` frames, skipping the first `skip` frames. Both values are clamped to `[0, 20]`. Each frame is formatted as:

```
package.Function
  dir/file.go:42
```

### Recover

```go
func Recover(fn func()) error
```

Calls `fn` and converts any panic into a `*PanicError`. Returns `nil` if `fn` does not panic.

### PanicError

```go
type PanicError struct {
	Value any    // the original value passed to panic()
	Stack string // full stack trace at the point of the panic
}

func (e *PanicError) Error() string
func (e *PanicError) Unwrap() error
```

Represents a recovered panic with captured runtime context. `Unwrap` returns the panic value if it implements `error`, enabling `errors.Is` and `errors.As` to reach through the wrapper.

## Examples

### Caller

```go
short, full, file, line, ok := runtimex.Caller(0)
fmt.Printf("%s (%s) at %s:%d\n", short, full, file, line)
// e.g. main.main (main.main) at main/main.go:12
```

### CallerFunc

```go
name, ok := runtimex.CallerFunc(0)
fmt.Println(name) // e.g. "main"
```

### Recover

```go
err := runtimex.Recover(func() {
	panic("something went wrong")
})

var pe *runtimex.PanicError
if errors.As(err, &pe) {
	fmt.Println(pe.Value) // "something went wrong"
	fmt.Println(pe.Stack) // full stack trace
}
```

### StackTrace

```go
trace := runtimex.StackTrace(5, 0)
fmt.Println(trace)
// main.handler
//   server/handler.go:28
// net/http.HandlerFunc.ServeHTTP
//   http/server.go:2136
// ...
```
