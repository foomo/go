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

### StackTrace

```go
func StackTrace(size, skip int) string
```

Captures and formats a stack trace of up to `size` frames, skipping the first `skip` frames. Both values are clamped to `[0, 20]`. Each frame is formatted as:

```
package.Function
  dir/file.go:42
```

## Examples

### Caller

```go
short, full, file, line, ok := runtimex.Caller(0)
fmt.Printf("%s (%s) at %s:%d\n", short, full, file, line)
// e.g. main.main (main.main) at main/main.go:12
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
