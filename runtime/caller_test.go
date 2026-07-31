package runtime_test

import (
	"fmt"
	"testing"

	"github.com/foomo/go/runtime"
	"github.com/stretchr/testify/assert"
)

type caller struct{}

func (c *caller) caller() (string, string, string, int, bool) {
	return runtime.Caller(0)
}

func ExampleCaller() {
	shortName, fullName, file, line, ok := runtime.Caller(0)

	fmt.Println(shortName)
	fmt.Println(fullName)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)

	// Output:
	// runtime_test.ExampleCaller
	// github.com/foomo/go/runtime_test.ExampleCaller
	// runtime/caller_test.go
	// 18
	// true
}

func TestCaller(t *testing.T) {
	t.Parallel()

	shortName, fullName, file, line, _ := runtime.Caller(0)

	assert.Equal(t, "runtime_test.TestCaller", shortName)
	assert.Equal(t, "github.com/foomo/go/runtime_test.TestCaller", fullName)
	assert.Equal(t, "runtime/caller_test.go", file)
	assert.Equal(t, 37, line)

	c := new(caller)
	shortName, fullName, file, line, _ = c.caller()

	assert.Equal(t, "runtime_test.(*caller).caller", shortName)
	assert.Equal(t, "github.com/foomo/go/runtime_test.(*caller).caller", fullName)
	assert.Equal(t, "runtime/caller_test.go", file)
	assert.Equal(t, 14, line)
}

// inlinedCaller is a leaf small enough for the compiler to inline into its
// caller. With FuncForPC the reported name would be the inliner; via
// CallersFrames it is correctly attributed to inlinedCaller.
func inlinedCaller() string {
	_, fullName, _, _, _ := runtime.Caller(0) //nolint:dogsled
	return fullName
}

func TestCallerInlined(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "github.com/foomo/go/runtime_test.inlinedCaller", inlinedCaller())
}
