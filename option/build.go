package option

import (
	"fmt"
)

// Build applies a slice of functions to a value.
func Build[T any](v T, builders ...interface{ List() []Option[T] }) {
	for _, b := range builders {
		if b != nil {
			Apply(v, b.List()...)
		}
	}
}

// BuildE applies a slice of error-returning functions to a value and stops on the first encountered error.
func BuildE[T any](v T, builders ...interface{ List() []OptionE[T] }) error {
	for _, b := range builders {
		if b != nil {
			if err := ApplyE(v, b.List()...); err != nil {
				return fmt.Errorf("build option: %w", err)
			}
		}
	}

	return nil
}
