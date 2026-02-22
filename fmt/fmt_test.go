package fmt_test

import (
	"testing"

	fmtx "github.com/foomo/go/fmt"
	"github.com/stretchr/testify/require"
)

func TestTprintf(t *testing.T) {
	t.Parallel()

	format := "%{.name} is %{.age} years old"
	result := fmtx.Tprintf(format, "name", "John", "age", "30")
	expected := "John is 30 years old"

	require.Equal(t, expected, result)
}
