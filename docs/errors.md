# errors

Helpers extending the standard `errors` package.

## Import

```go
import goerrors "github.com/foomo/go/errors"
```

## API

### AsAny

```go
func AsAny(err error, targets ...any) bool
```

Reports whether `err` matches any of the targets via `errors.As`. Each target must be a non-nil pointer to either a type that implements `error`, or to any interface type.

### IsAny

```go
func IsAny(err error, targets ...error) bool
```

Reports whether `err` matches any of the targets via `errors.Is`.

### AsAnyType

```go
func AsAnyType(err error, targets ...any) bool
```

Reports whether `err` matches the concrete type of any of the targets via `errors.As`. Each target must be a pointer to the error type to match (e.g. `&fs.PathError{}`); unlike `AsAny`, the caller need not declare target variables and the matched value is discarded. `nil` targets are skipped.

### Cause

```go
func Cause(err error) error
```

Unwraps `err` through successive [`types.Causer`](/types#causer) implementations and returns the root cause. Stops if an error is not a `Causer`, or if `Cause()` returns the same error instance.

## Examples

### AsAny

```go
var (
	pathErr *fs.PathError
	numErr  *strconv.NumError
)

_, err := os.Open("/nonexistent/path")
if goerrors.AsAny(err, &numErr, &pathErr) {
	// err unwrapped into one of the targets
}
```

### IsAny

```go
err := fmt.Errorf("wrapped: %w", io.EOF)
if goerrors.IsAny(err, context.Canceled, io.EOF) {
	// err matches one of the sentinels
}
```

### AsAnyType

```go
_, err := os.Open("/nonexistent/path")
if goerrors.AsAnyType(err, &fs.PathError{}, &strconv.NumError{}) {
	// err's concrete type matched one of the targets
}
```

### Cause

```go
type wrapErr struct {
	msg   string
	cause error
}

func (e *wrapErr) Error() string { return e.msg }
func (e *wrapErr) Cause() error  { return e.cause }

err := &wrapErr{msg: "outer", cause: &wrapErr{msg: "inner", cause: io.EOF}}
root := goerrors.Cause(err) // io.EOF
```
