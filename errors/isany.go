package errors

import (
	"errors"
)

// IsAny reports whether err matches any of the targets via errors.Is.
func IsAny(err error, targets ...error) bool {
	for _, target := range targets {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}
