package types

// Causer is implemented by errors that wrap an underlying cause, distinct
// from the chain unwrapped by errors.Unwrap.
type Causer interface {
	Cause() error
}

// CauserFunc adapts a func() error to a Causer.
type CauserFunc func() error

// Cause calls f and returns its result.
func (f CauserFunc) Cause() error {
	return f()
}
