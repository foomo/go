package runtime

import "runtime/debug"

// Recover calls fn and converts any panic into a *[PanicError].
// Returns nil if fn does not panic.
func Recover(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = &PanicError{
				Value: r,
				Stack: string(debug.Stack()),
			}
		}
	}()

	fn()

	return nil
}
