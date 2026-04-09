package fmt_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	fmtx "github.com/foomo/go/fmt"
)

func ExampleTprintf() {
	format := "%{.name} is %{.age} years old"
	fmt.Println(fmtx.Tprintf(format, "name", "John", "age", "30"))

	// Output: John is 30 years old
}

func TestTprintf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		format string
		pairs  []string
		want   string
	}{
		{
			name:   "simple substitution",
			format: "%{.name} is %{.age} years old",
			pairs:  []string{"name", "John", "age", "30"},
			want:   "John is 30 years old",
		},
		{
			name:   "missing key outputs empty",
			format: "hello %{.missing}!",
			pairs:  []string{"name", "John"},
			want:   "hello !",
		},
		{
			name:   "no placeholders",
			format: "just a string",
			pairs:  []string{"key", "val"},
			want:   "just a string",
		},
		{
			name:   "empty format",
			format: "",
			want:   "",
		},
		{
			name:   "odd pairs drops trailing key",
			format: "%{.a} %{.b}",
			pairs:  []string{"a", "1", "b"},
			want:   "1 ",
		},
		{
			name:   "unclosed delimiter is literal",
			format: "hello %{.name",
			pairs:  []string{"name", "world"},
			want:   "hello %{.name",
		},
		{
			name:   "no dot prefix is literal",
			format: "%{name}",
			pairs:  []string{"name", "John"},
			want:   "%{name}",
		},
		{
			name:   "adjacent placeholders",
			format: "%{.a}%{.b}",
			pairs:  []string{"a", "1", "b", "2"},
			want:   "12",
		},
		{
			name:   "percent without brace is literal",
			format: "100% done %{.status}",
			pairs:  []string{"status", "ok"},
			want:   "100% done ok",
		},
		{
			name:   "no pairs",
			format: "%{.name}",
			want:   "",
		},
		{
			name:   "empty key",
			format: "%{.}",
			pairs:  []string{"", "val"},
			want:   "val",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, fmtx.Tprintf(tt.format, tt.pairs...))
		})
	}
}
