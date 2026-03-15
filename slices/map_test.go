package slices_test

import (
	"fmt"

	"github.com/foomo/go/slices"
)

func ExampleMap() {
	result := slices.Map([]int{1, 2, 3}, func(n int) int {
		return n * 2
	})
	fmt.Println(result)
	// Output: [2 4 6]
}
