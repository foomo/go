package strings_test

import (
	"fmt"

	"github.com/foomo/go/strings"
)

func ExampleIsEmpty() {
	fmt.Println(strings.IsEmpty(""))
	// Output: true
}

func ExampleIsBlank() {
	fmt.Println(strings.IsBlank(" \t"))
	// Output: true
}

func ExampleIsAnyEmpty() {
	fmt.Println(strings.IsAnyEmpty("a", ""))
	// Output: true
}

func ExampleIsAnyBlank() {
	fmt.Println(strings.IsAnyBlank("a", " "))
	// Output: true
}

func ExampleIsAlpha() {
	fmt.Println(strings.IsAlpha("abc"))
	// Output: true
}

func ExampleIsAlphanumeric() {
	fmt.Println(strings.IsAlphanumeric("abc1"))
	// Output: true
}

func ExampleIsNumeric() {
	fmt.Println(strings.IsNumeric("123"))
	// Output: true
}

func ExampleIsNumerical() {
	fmt.Println(strings.IsNumerical("12.3"))
	// Output: true
}
