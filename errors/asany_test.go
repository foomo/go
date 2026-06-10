package errors_test

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"

	pkgerrors "github.com/foomo/go/errors"
)

func ExampleAsAny() {
	var (
		pathErr *fs.PathError
		numErr  *strconv.NumError
	)

	_, err := os.Open("/nonexistent/path/for/example")
	fmt.Println(pkgerrors.AsAny(err, &numErr, &pathErr))
	// Output: true
}
