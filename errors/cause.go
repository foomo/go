package errors

import "github.com/foomo/go/types"

// Cause unwraps err through successive [types.Causer] implementations and
// returns the root cause. It stops if an error is not a Causer, nil or if
// Cause returns the same error instance (which would otherwise loop forever).
// ref: https://github.com/pkg/errors/pull/143
func Cause(err error) error {
	for err != nil {
		causerErr, ok := err.(types.Causer)
		if !ok {
			break
		}

		nextErr := causerErr.Cause()

		if nextErr == nil || nextErr == err { //nolint:errorlint
			break
		}

		err = nextErr
	}

	return err
}
