package runtime

import "fmt"

// PanicError represents a recovered panic with captured runtime context.
type PanicError struct {
	// Value is the original value passed to panic().
	Value any
	// Stack is the full stack trace at the point of the panic.
	Stack string
}

// Error returns a string representation including the panic value and stack trace.
func (e *PanicError) Error() string {
	return fmt.Sprintf("panic: %v\n%s", e.Value, e.Stack)
}

// Unwrap returns the panic value if it implements error, nil otherwise.
// This enables errors.Is and errors.As to reach through the PanicError wrapper.
func (e *PanicError) Unwrap() error {
	if err, ok := e.Value.(error); ok {
		return err
	}

	return nil
}
