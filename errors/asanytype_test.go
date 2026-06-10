package errors_test

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"

	pkgerrors "github.com/foomo/go/errors"
)

func ExampleAsAnyType() {
	_, err := os.Open("/nonexistent/path/for/example")
	fmt.Println(pkgerrors.AsAnyType(err, &fs.PathError{}, &strconv.NumError{}))
	// Output: true
}
