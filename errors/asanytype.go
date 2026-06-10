package errors

import (
	"errors"
)

// AsAnyType reports whether err matches any of the targets via errors.As.
func AsAnyType(err error, targets ...any) bool {
	for _, target := range targets {
		if errors.As(err, &target) {
			return true
		}
	}

	return false
}
