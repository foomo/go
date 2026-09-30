package testing

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/foomo/go/runtime"
)

// ExampleTB implements [testing.TB] for use in godoc Example functions, which
// receive no *testing.T. It is not concurrency-safe and does not support
// benchmarking; Fatal and Fatalf terminate the process via os.Exit instead of
// a goroutine-local exit, since examples run on the main goroutine.
type ExampleTB struct {
	testing.TB // nil embed ok for examples
	name       string
	failed     bool
	skipped    bool
}

// NewExampleTB creates one for examples
func NewExampleTB() *ExampleTB {
	name, _ := runtime.CallerFunc(0)
	return &ExampleTB{name: name}
}

// Name returns the name of the running example.
func (t *ExampleTB) Name() string {
	return t.name
}

// Helper is a no-op; examples have no call stack to mark.
func (t *ExampleTB) Helper() {}

// Cleanup is a no-op; examples have no cleanup phase.
func (t *ExampleTB) Cleanup(_ func()) {}

// Attr is a no-op; examples do not report attributes.
func (t *ExampleTB) Attr(_, _ string) {}

// Fail marks the example as having failed.
func (t *ExampleTB) Fail() {
	t.failed = true
}

// FailNow marks the example as having failed. Unlike [testing.T.FailNow], it
// does not stop execution of the current goroutine.
func (t *ExampleTB) FailNow() {
	t.failed = true
}

// Failed reports whether Fail, FailNow, Error, Errorf, Fatal, or Fatalf has
// been called.
func (t *ExampleTB) Failed() bool {
	return t.failed
}

// Error marks the example as failed and logs args to Output.
func (t *ExampleTB) Error(args ...any) {
	t.failed = true
	_, _ = fmt.Fprint(t.Output(), "error: ")
	_, _ = fmt.Fprint(t.Output(), args...)
	_, _ = fmt.Fprint(t.Output(), "\n")
}

// Errorf marks the example as failed and logs a formatted message to Output.
func (t *ExampleTB) Errorf(format string, args ...any) {
	t.failed = true
	_, _ = fmt.Fprint(t.Output(), "error: ")
	_, _ = fmt.Fprintf(t.Output(), "error: "+format, args...)
	_, _ = fmt.Fprint(t.Output(), "\n")
}

// Fatal marks the example as failed, logs args to Output, and calls os.Exit(1).
func (t *ExampleTB) Fatal(args ...any) {
	t.failed = true
	_, _ = fmt.Fprint(t.Output(), "fatal: ")
	_, _ = fmt.Fprint(t.Output(), args...)
	_, _ = fmt.Fprint(t.Output(), "\n")

	os.Exit(1)
}

// Fatalf marks the example as failed, logs a formatted message to Output, and
// calls os.Exit(1).
func (t *ExampleTB) Fatalf(format string, args ...any) {
	t.failed = true
	_, _ = fmt.Fprintf(t.Output(), "fatal: "+format, args...)
	_, _ = fmt.Fprint(t.Output(), "\n")

	os.Exit(1)
}

// Log logs args to Output.
func (t *ExampleTB) Log(args ...any) {
	_, _ = fmt.Fprint(t.Output(), args...)
	_, _ = fmt.Fprint(t.Output(), "\n")
}

// Logf logs a formatted message to Output.
func (t *ExampleTB) Logf(format string, args ...any) {
	_, _ = fmt.Fprintf(t.Output(), format, args...)
	_, _ = fmt.Fprint(t.Output(), "\n")
}

// Skip marks the example as skipped and logs args to Output.
func (t *ExampleTB) Skip(args ...any) {
	t.skipped = true
	_, _ = fmt.Fprint(t.Output(), "skip: ")
	_, _ = fmt.Fprint(t.Output(), args...)
	_, _ = fmt.Fprint(t.Output(), "\n")
}

// SkipNow marks the example as skipped. Unlike [testing.T.SkipNow], it does
// not stop execution of the current goroutine.
func (t *ExampleTB) SkipNow() {
	t.skipped = true
}

// Skipf marks the example as skipped and logs a formatted message to Output.
func (t *ExampleTB) Skipf(format string, args ...any) {
	t.skipped = true
	_, _ = fmt.Fprintf(t.Output(), "skip: "+format, args...)
	_, _ = fmt.Fprint(t.Output(), "\n")
}

// Skipped reports whether Skip, SkipNow, or Skipf has been called.
func (t *ExampleTB) Skipped() bool {
	return t.skipped
}

// Setenv sets an environment variable for the duration of the process; it is
// never restored, since examples have no cleanup phase.
func (t *ExampleTB) Setenv(key, value string) {
	_ = os.Setenv(key, value)
}

// Chdir changes the working directory for the duration of the process; it is
// never restored, since examples have no cleanup phase.
func (t *ExampleTB) Chdir(dir string) {
	_ = os.Chdir(dir)
}

// TempDir creates and returns a new temporary directory; it is never removed,
// since examples have no cleanup phase.
func (t *ExampleTB) TempDir() string {
	dir, _ := os.MkdirTemp("", "exampletb-*")
	return dir
}

// ArtifactDir creates and returns a new temporary directory for artifacts; it
// is never removed, since examples have no cleanup phase.
func (t *ExampleTB) ArtifactDir() string {
	dir, _ := os.MkdirTemp("", "exampletb-artifacts-*")
	return dir
}

// Context returns context.Background().
func (t *ExampleTB) Context() context.Context {
	return context.Background()
}

// Output returns os.Stdout, the writer used by Log, Error, and related methods.
func (t *ExampleTB) Output() io.Writer {
	return os.Stdout
}

// Run creates a sub-ExampleTB named t.Name()+"/"+name, runs f synchronously
// with it, and always returns true.
func (t *ExampleTB) Run(name string, f func(testing.TB)) bool {
	sub := &ExampleTB{name: t.name + "/" + name}
	f(sub)

	return true
}
