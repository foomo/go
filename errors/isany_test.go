package errors_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"testing"

	pkgerrors "github.com/foomo/go/errors"
)

func ExampleIsAny() {
	err := fmt.Errorf("wrapped: %w", io.EOF)
	fmt.Println(pkgerrors.IsAny(err, context.Canceled, io.EOF))
	// Output: true
}

func TestIsAny(t *testing.T) {
	sentinel := errors.New("sentinel")

	tests := []struct {
		name    string
		err     error
		targets []error
		want    bool
	}{
		{
			name:    "no targets",
			err:     io.EOF,
			targets: nil,
			want:    false,
		},
		{
			name:    "nil err and nil target match",
			err:     nil,
			targets: []error{nil},
			want:    true, // errors.Is(nil, nil) == true
		},
		{
			name:    "nil err with real target",
			err:     nil,
			targets: []error{io.EOF},
			want:    false,
		},
		{
			name:    "match on first target",
			err:     io.EOF,
			targets: []error{io.EOF, context.Canceled},
			want:    true,
		},
		{
			name:    "match on last target",
			err:     context.Canceled,
			targets: []error{io.EOF, sentinel, context.Canceled},
			want:    true,
		},
		{
			name:    "matches none",
			err:     sentinel,
			targets: []error{io.EOF, context.Canceled},
			want:    false,
		},
		{
			name:    "wrapped error matches",
			err:     fmt.Errorf("wrapped: %w", io.EOF),
			targets: []error{context.Canceled, io.EOF},
			want:    true,
		},
		{
			name:    "deeply wrapped error matches",
			err:     fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", sentinel)),
			targets: []error{io.EOF, sentinel},
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pkgerrors.IsAny(tt.err, tt.targets...); got != tt.want {
				t.Errorf("IsAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
