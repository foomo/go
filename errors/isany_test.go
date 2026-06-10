package errors_test

import (
	"context"
	"fmt"
	"io"

	pkgerrors "github.com/foomo/go/errors"
)

func ExampleIsAny() {
	err := fmt.Errorf("wrapped: %w", io.EOF)
	fmt.Println(pkgerrors.IsAny(err, context.Canceled, io.EOF))
	// Output: true
}
