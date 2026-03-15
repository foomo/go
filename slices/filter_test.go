package slices_test

import (
	"fmt"

	"github.com/foomo/go/slices"
)

func ExampleFilter() {
	result := slices.Filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(result)
	// Output: [2 4]
}
