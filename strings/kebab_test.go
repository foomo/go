package strings_test

import (
	"fmt"

	"github.com/foomo/go/strings"
)

func ExampleToKebab() {
	result := strings.ToKebab("HelloWorldExample")
	fmt.Println(result)
	// Output: hello-world-example
}

func ExampleToScreamingKebab() {
	result := strings.ToScreamingKebab("HelloWorldExample")
	fmt.Println(result)
	// Output: HELLO-WORLD-EXAMPLE
}
