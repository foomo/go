package slices_test

import (
	"strconv"
	"testing"

	"github.com/foomo/go/slices"
)

func BenchmarkMap(b *testing.B) {
	items := make([]int, 1000)
	for i := range items {
		items[i] = i
	}

	b.ResetTimer()

	for b.Loop() {
		slices.Map(items, strconv.Itoa)
	}
}

func BenchmarkFilter(b *testing.B) {
	items := make([]int, 1000)
	for i := range items {
		items[i] = i
	}

	b.ResetTimer()

	for b.Loop() {
		slices.Filter(items, func(i int) bool {
			return i%2 == 0
		})
	}
}
