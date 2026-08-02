package errors_test

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"testing"

	goerrors "github.com/foomo/go/errors"
)

func ExampleAsAnyType() {
	_, err := os.Open("/nonexistent/path/for/example")
	fmt.Println(goerrors.AsAnyType(err, &fs.PathError{}, &strconv.NumError{}))

	// Output: true
}

// customError is a distinct error type used to exercise type-based matching.
type customError struct{ msg string }

func (e *customError) Error() string { return e.msg }

func TestAsAnyType(t *testing.T) {
	pathErr := &fs.PathError{Op: "open", Path: "/x", Err: io.EOF}

	tests := []struct {
		name    string
		err     error
		targets []any
		want    bool
	}{
		{
			name:    "matching target type",
			err:     pathErr,
			targets: []any{&fs.PathError{}},
			want:    true,
		},
		{
			name:    "unrelated target type",
			err:     pathErr,
			targets: []any{&strconv.NumError{}},
			want:    false, // would be true before the AsAnyType fix
		},
		{
			name:    "nil err",
			err:     nil,
			targets: []any{&fs.PathError{}},
			want:    false,
		},
		{
			name:    "no targets",
			err:     pathErr,
			targets: nil,
			want:    false,
		},
		{
			name:    "match on first target",
			err:     pathErr,
			targets: []any{&fs.PathError{}, &strconv.NumError{}},
			want:    true,
		},
		{
			name:    "match on second target",
			err:     pathErr,
			targets: []any{&strconv.NumError{}, &fs.PathError{}},
			want:    true,
		},
		{
			name:    "matches none",
			err:     errors.New("plain"),
			targets: []any{&fs.PathError{}, &strconv.NumError{}},
			want:    false,
		},
		{
			name:    "wrapped error type match",
			err:     fmt.Errorf("wrapped: %w", pathErr),
			targets: []any{&strconv.NumError{}, &fs.PathError{}},
			want:    true,
		},
		{
			name:    "deeply wrapped custom error type match",
			err:     fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", &customError{msg: "boom"})),
			targets: []any{&fs.PathError{}, &customError{}},
			want:    true,
		},
		{
			name:    "nil target skipped then match",
			err:     pathErr,
			targets: []any{nil, &fs.PathError{}},
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goerrors.AsAnyType(tt.err, tt.targets...); got != tt.want {
				t.Errorf("AsAnyType() = %v, want %v", got, tt.want)
			}
		})
	}
}
