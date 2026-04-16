package testing

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/foomo/go/runtime"
)

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

func (t *ExampleTB) Name() string {
	return t.name
}

func (t *ExampleTB) Helper() {}

func (t *ExampleTB) Cleanup(_ func()) {}

func (t *ExampleTB) Attr(_, _ string) {}

func (t *ExampleTB) Fail() {
	t.failed = true
}

func (t *ExampleTB) FailNow() {
	t.failed = true
}

func (t *ExampleTB) Failed() bool {
	return t.failed
}

func (t *ExampleTB) Error(args ...any) {
	t.failed = true
	_, _ = fmt.Fprint(t.Output(), "error: ")
	_, _ = fmt.Fprint(t.Output(), args...)
	_, _ = fmt.Fprint(t.Output(), "\n")
}

func (t *ExampleTB) Errorf(format string, args ...any) {
	t.failed = true
	_, _ = fmt.Fprint(t.Output(), "error: ")
	_, _ = fmt.Fprintf(t.Output(), "error: "+format, args...)
	_, _ = fmt.Fprint(t.Output(), "\n")
}

func (t *ExampleTB) Fatal(args ...any) {
	t.failed = true
	_, _ = fmt.Fprint(t.Output(), "fatal: ")
	_, _ = fmt.Fprint(t.Output(), args...)
	_, _ = fmt.Fprint(t.Output(), "\n")

	os.Exit(1)
}

func (t *ExampleTB) Fatalf(format string, args ...any) {
	t.failed = true
	_, _ = fmt.Fprintf(t.Output(), "fatal: "+format, args...)
	_, _ = fmt.Fprint(t.Output(), "\n")

	os.Exit(1)
}

func (t *ExampleTB) Log(args ...any) {
	_, _ = fmt.Fprint(t.Output(), args...)
	_, _ = fmt.Fprint(t.Output(), "\n")
}

func (t *ExampleTB) Logf(format string, args ...any) {
	_, _ = fmt.Fprintf(t.Output(), format, args...)
	_, _ = fmt.Fprint(t.Output(), "\n")
}

func (t *ExampleTB) Skip(args ...any) {
	t.skipped = true
	_, _ = fmt.Fprint(t.Output(), "skip: ")
	_, _ = fmt.Fprint(t.Output(), args...)
	_, _ = fmt.Fprint(t.Output(), "\n")
}

func (t *ExampleTB) SkipNow() {
	t.skipped = true
}

func (t *ExampleTB) Skipf(format string, args ...any) {
	t.skipped = true
	_, _ = fmt.Fprintf(t.Output(), "skip: "+format, args...)
	_, _ = fmt.Fprint(t.Output(), "\n")
}

func (t *ExampleTB) Skipped() bool {
	return t.skipped
}

func (t *ExampleTB) Setenv(key, value string) {
	_ = os.Setenv(key, value)
}

func (t *ExampleTB) Chdir(dir string) {
	_ = os.Chdir(dir)
}

func (t *ExampleTB) TempDir() string {
	dir, _ := os.MkdirTemp("", "exampletb-*")
	return dir
}

func (t *ExampleTB) ArtifactDir() string {
	dir, _ := os.MkdirTemp("", "exampletb-artifacts-*")
	return dir
}

func (t *ExampleTB) Context() context.Context {
	return context.Background()
}

func (t *ExampleTB) Output() io.Writer {
	return os.Stdout
}

func (t *ExampleTB) Run(name string, f func(testing.TB)) bool {
	sub := &ExampleTB{name: t.name + "/" + name}
	f(sub)

	return true
}
