# slog

Test-friendly `slog.Handler` that writes log output through `testing.TB`.

## Import

```go
import slogx "github.com/foomo/go/slog"
```

## API

### NewTestHandler

```go
func NewTestHandler(tb testing.TB, opts ...TestHandlerOption) slog.Handler
```

Returns an `slog.Handler` that writes log records to `tb.Output` in a compact format:

```
file.go:42: [LEVEL] msg key=value ...
```

Defaults to `slog.LevelDebug`. Supports `WithAttrs` and `WithGroup` for structured context.

### TestHandlerWithLevel

```go
func TestHandlerWithLevel(level slog.Leveler) TestHandlerOption
```

Sets the minimum log level for the test handler.

## Example

```go
func TestService(t *testing.T) {
	logger := slog.New(slogx.NewTestHandler(t))
	logger.Info("starting", "port", 8080)
	// Output via t.Log: testfile_test.go:4: [INFO] starting port=8080
}
```

### With minimum level

```go
func TestServiceWarn(t *testing.T) {
	logger := slog.New(slogx.NewTestHandler(t, slogx.TestHandlerWithLevel(slog.LevelWarn)))
	logger.Debug("ignored") // not printed
	logger.Warn("something happened", "err", "timeout")
}
```
