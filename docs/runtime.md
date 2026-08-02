# runtime

Runtime introspection utilities.

## Import

```go
import goruntime "github.com/foomo/go/runtime"
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

### CallFrame

```go
func CallFrame(skip int) Frame
```

Returns the call site `skip` levels above its own caller — `skip=0` is the function that called `CallFrame`, `skip=1` that function's caller, and so on. Symbolisation is memoised per program counter, so the steady-state cost is one `runtime.Callers` into a stack array plus a map load, with no allocation. Returns a zero `Frame` (see `Frame.Zero`) when the stack is unavailable.

### Frame

```go
type Frame struct {
	Pkg  string // full import path, e.g. "github.com/foomo/go/runtime"
	Inst string // receiver type name for methods; empty for package-level funcs
	Func string // function or method name, incl. suffixes like ".func1" or "-fm"
	File string // absolute path as recorded at build time
	Line int
}

func (f Frame) Name() string
func (f Frame) Short() string
func (f Frame) Zero() bool
```

Identifies a single call site. `Name` reassembles the qualified name (`pkg.Recv.Func`), suitable for the `code.function.name` attribute; the pointer-receiver marker is not preserved. `Short` renders the receiver-qualified name without the package path (`Recv.Func`). `Zero` reports whether the frame could not be resolved.

### Memo

```go
type Memo[T any] struct { /* ... */ }

func (mo *Memo[T]) Get(skip int, derive func(Frame) T) T
```

Caches a value of type `T` per call site, so per-site derived data — a span name, a preformatted attribute slice, a logger — is computed once for the life of the process rather than on every call. `Get` resolves the call site `skip` levels above its caller (`skip=0` is the function that called `Get`) and returns the memoised value, invoking `derive` at most once per site under normal conditions. `derive` must be pure and its result safe to share; values handed out by `Get` are shared across all calls from that site and must be treated as read-only. The zero value is ready to use and a `Memo` must not be copied after first use.

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
short, full, file, line, ok := goruntime.Caller(0)
fmt.Printf("%s (%s) at %s:%d\n", short, full, file, line)
// e.g. main.main (main.main) at main/main.go:12
```

### CallerFunc

```go
name, ok := goruntime.CallerFunc(0)
fmt.Println(name) // e.g. "main"
```

### CallFrame

```go
f := goruntime.CallFrame(0)
fmt.Println(f.Name())  // e.g. "main.main"
fmt.Println(f.Short()) // e.g. "main"
```

### Memo

```go
// spans caches a span name derived once per call site.
var spans goruntime.Memo[string]

func handle() {
	name := spans.Get(0, func(f goruntime.Frame) string {
		return f.Short() // computed once for this call site
	})
	_ = name
}
```

### Recover

```go
err := goruntime.Recover(func() {
	panic("something went wrong")
})

var pe *goruntime.PanicError
if errors.As(err, &pe) {
	fmt.Println(pe.Value) // "something went wrong"
	fmt.Println(pe.Stack) // full stack trace
}
```

### StackTrace

```go
trace := goruntime.StackTrace(5, 0)
fmt.Println(trace)
// main.handler
//   server/handler.go:28
// net/http.HandlerFunc.ServeHTTP
//   http/server.go:2136
// ...
```
