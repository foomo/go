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

func ExampleAsAny() {
	var (
		pathErr *fs.PathError
		numErr  *strconv.NumError
	)

	_, err := os.Open("/nonexistent/path/for/example")
	fmt.Println(goerrors.AsAny(err, &numErr, &pathErr))
	// Output: true
}

func TestAsAny(t *testing.T) {
	pathErr := &fs.PathError{Op: "open", Path: "/x", Err: io.EOF}

	tests := []struct {
		name string
		err  error
		// targets returns fresh target pointers for each run.
		targets func() []any
		want    bool
	}{
		{
			name:    "nil err",
			err:     nil,
			targets: func() []any { var p *fs.PathError; return []any{&p} },
			want:    false,
		},
		{
			name:    "no targets",
			err:     pathErr,
			targets: func() []any { return nil },
			want:    false,
		},
		{
			name:    "matches none",
			err:     errors.New("plain"),
			targets: func() []any { var p *fs.PathError; var n *strconv.NumError; return []any{&p, &n} },
			want:    false,
		},
		{
			name:    "match on first target",
			err:     pathErr,
			targets: func() []any { var p *fs.PathError; var n *strconv.NumError; return []any{&p, &n} },
			want:    true,
		},
		{
			name:    "match on last target",
			err:     pathErr,
			targets: func() []any { var n *strconv.NumError; var p *fs.PathError; return []any{&n, &p} },
			want:    true,
		},
		{
			name:    "wrapped error matches",
			err:     fmt.Errorf("wrapped: %w", pathErr),
			targets: func() []any { var p *fs.PathError; return []any{&p} },
			want:    true,
		},
		{
			name:    "deeply wrapped error matches",
			err:     fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", pathErr)),
			targets: func() []any { var p *fs.PathError; return []any{&p} },
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goerrors.AsAny(tt.err, tt.targets()...); got != tt.want {
				t.Errorf("AsAny() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("populates matched target", func(t *testing.T) {
		var (
			numErr  *strconv.NumError
			pathTgt *fs.PathError
		)
		if !goerrors.AsAny(pathErr, &numErr, &pathTgt) {
			t.Fatal("AsAny() = false, want true")
		}

		if pathTgt == nil {
			t.Fatal("matched target *fs.PathError was not populated")
		}

		if pathTgt.Path != "/x" {
			t.Errorf("populated target Path = %q, want %q", pathTgt.Path, "/x")
		}

		if numErr != nil {
			t.Error("non-matching target *strconv.NumError should remain nil")
		}
	})
}
