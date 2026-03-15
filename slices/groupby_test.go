package slices_test

import (
	"fmt"

	"github.com/foomo/go/slices"
)

func ExampleGroupBy() {
	type item struct {
		name     string
		category string
	}

	items := []item{
		{"apple", "fruit"},
		{"carrot", "vegetable"},
		{"banana", "fruit"},
	}
	groups := slices.GroupBy(items, func(i item) string {
		return i.category
	})
	fmt.Println(len(groups))
	fmt.Println(len(groups["fruit"]))
	fmt.Println(len(groups["vegetable"]))
	// Output:
	// 2
	// 2
	// 1
}
