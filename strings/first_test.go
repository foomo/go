package strings_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/foomo/go/strings"
)

func ExampleFirstToUpper() {
	result := strings.FirstToUpper("hello")
	fmt.Println(result)
	// Output: Hello
}

func ExampleFirstToLower() {
	result := strings.FirstToLower("Hello")
	fmt.Println(result)
	// Output: hello
}

func TestFirstToUpper(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"non-empty", "hello", "Hello"},
		{"empty", "", ""},
		{"single char", "a", "A"},
		{"already upper", "Hello", "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, strings.FirstToUpper(tt.input))
		})
	}
}

func TestFirstToLower(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"non-empty", "Hello", "hello"},
		{"empty", "", ""},
		{"single char", "A", "a"},
		{"already lower", "hello", "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, strings.FirstToLower(tt.input))
		})
	}
}
