package strings_test

import (
	"fmt"

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
