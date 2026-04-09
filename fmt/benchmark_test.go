package fmt_test

import (
	"testing"

	"github.com/foomo/go/fmt"
)

func BenchmarkTprintf(b *testing.B) {
	for b.Loop() {
		fmt.Tprintf("Hello %{.name}, you are %{.age} years old", "name", "World", "age", "30")
	}
}
