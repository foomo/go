package errors

import (
	"errors"
	"reflect"
)

// AsAnyType reports whether err matches the concrete type of any of the targets
// via errors.As. Each target must be a pointer to the error type to match
// (e.g. &fs.PathError{}); unlike AsAny, the caller need not declare target
// variables and the matched value is discarded. nil targets are skipped.
func AsAnyType(err error, targets ...any) bool {
	if err == nil {
		return false
	}

	for _, target := range targets {
		if target == nil {
			continue
		}
		// reflect.New yields a **T so errors.As has a settable *T target.
		ptr := reflect.New(reflect.TypeOf(target)).Interface()
		if errors.As(err, ptr) {
			return true
		}
	}

	return false
}
