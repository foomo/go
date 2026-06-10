# errors

Helpers extending the standard `errors` package.

## Import

```go
import errorsx "github.com/foomo/go/errors"
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

## Examples

### AsAny

```go
var (
	pathErr *fs.PathError
	numErr  *strconv.NumError
)

_, err := os.Open("/nonexistent/path")
if errorsx.AsAny(err, &numErr, &pathErr) {
	// err unwrapped into one of the targets
}
```

### IsAny

```go
err := fmt.Errorf("wrapped: %w", io.EOF)
if errorsx.IsAny(err, context.Canceled, io.EOF) {
	// err matches one of the sentinels
}
```
