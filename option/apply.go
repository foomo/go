package option

import (
	"fmt"
)

// Apply applies a slice of functions to a value.
func Apply[T any](v T, opts ...Option[T]) {
	for _, opt := range opts {
		if opt != nil {
			opt(v)
		}
	}
}

// ApplyE applies a slice of error-returning functions to a value and stops on the first encountered error.
func ApplyE[T any](v T, opts ...OptionE[T]) error {
	for _, opt := range opts {
		if opt != nil {
			if err := opt(v); err != nil {
				return fmt.Errorf("apply option %T: %w", opt, err)
			}
		}
	}

	return nil
}
