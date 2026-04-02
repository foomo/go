package runtime_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/foomo/go/runtime"
	"github.com/stretchr/testify/assert"
)

func ExampleStackTrace() {
	stack := runtime.StackTrace(2, 0)

	fmt.Println(stack)

	// Output:
	// github.com/foomo/go/runtime_test.ExampleStackTrace
	//   runtime/stacktrace_test.go:13
	// testing.runExample
	//   testing/run_example.go:63
}

func TestStackTrace(t *testing.T) {
	t.Parallel()

	stack := runtime.StackTrace(2, 0)
	parts := strings.Split(stack, "\n")

	assert.Len(t, parts, 4)
	assert.Equal(t, "github.com/foomo/go/runtime_test.TestStackTrace", parts[0])
}

func TestStackTrace_size(t *testing.T) {
	t.Parallel()

	t.Run("zero", func(t *testing.T) {
		stack := runtime.StackTrace(0, 0)
		parts := strings.Split(stack, "\n")

		assert.Len(t, parts, 2)
		assert.Equal(t, "github.com/foomo/go/runtime_test.TestStackTrace_size.func1", parts[0])
	})

	t.Run("negative", func(t *testing.T) {
		stack := runtime.StackTrace(-2, 0)
		parts := strings.Split(stack, "\n")

		assert.Len(t, parts, 2)
		assert.Equal(t, "github.com/foomo/go/runtime_test.TestStackTrace_size.func2", parts[0])
	})

	t.Run("too large", func(t *testing.T) {
		stack := runtime.StackTrace(100_000, 0)
		parts := strings.Split(stack, "\n")

		assert.Len(t, parts, 6)
		assert.Equal(t, "github.com/foomo/go/runtime_test.TestStackTrace_size.func3", parts[0])
	})
}

func TestStackTrace_skip(t *testing.T) {
	t.Parallel()

	t.Run("zero", func(t *testing.T) {
		stack := runtime.StackTrace(2, 0)
		parts := strings.Split(stack, "\n")

		assert.Len(t, parts, 4)
		assert.Equal(t, "github.com/foomo/go/runtime_test.TestStackTrace_skip.func1", parts[0])
	})

	t.Run("negative", func(t *testing.T) {
		stack := runtime.StackTrace(2, -1)
		parts := strings.Split(stack, "\n")

		assert.Len(t, parts, 4)
		assert.Equal(t, "github.com/foomo/go/runtime_test.TestStackTrace_skip.func2", parts[0])
	})

	t.Run("too large", func(t *testing.T) {
		stack := runtime.StackTrace(2, 100_000)
		assert.Equal(t, "stack trace out of bounds", stack)
	})
}
