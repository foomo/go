package errors

import "github.com/foomo/go/types"

// Cause protects against the infinite loop
func Cause(err error) error {
	for err != nil {
		causerErr, ok := err.(types.Causer)
		if !ok {
			break
		}

		nextErr := causerErr.Cause()

		// if Cause() returns the same error instance,
		// we have reached the root. Break to avoid an infinite loop.
		if nextErr == err { //nolint:errorlint // explizit comparison
			break
		}

		err = nextErr
	}

	return err
}
