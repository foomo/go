package errors

import "github.com/foomo/go/types"

// Cause unwraps err through successive [types.Causer] implementations and
// returns the root cause. It stops if an error is not a Causer, or if
// Cause returns the same error instance (which would otherwise loop forever).
func Cause(err error) error {
	for err != nil {
		causerErr, ok := err.(types.Causer)
		if !ok {
			break
		}

		nextErr := causerErr.Cause()

		// if Cause() returns nil or the same error instance,
		// we have reached the root. Break to avoid an infinite loop.
		if nextErr == nil || nextErr == err { //nolint:errorlint // explizit comparison
			break
		}

		err = nextErr
	}

	return err
}
