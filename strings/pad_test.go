package strings_test

import (
	"fmt"

	"github.com/foomo/go/strings"
)

func ExamplePadRight() {
	result := strings.PadRight("hello", 10)
	fmt.Printf("'%s'", result)
	// Output: 'hello     '
}

func ExamplePadLeft() {
	result := strings.PadLeft("hello", 10)
	fmt.Printf("'%s'", result)
	// Output: '     hello'
}
