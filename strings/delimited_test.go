package strings_test

import (
	"fmt"

	"github.com/foomo/go/strings"
)

func ExampleToDelimited() {
	result := strings.ToDelimited("HelloWorldExample", '.')
	fmt.Println(result)
	// Output: hello.world.example
}

func ExampleToScreamingDelimited() {
	result := strings.ToScreamingDelimited("HelloWorldExample", '.', "", true)
	fmt.Println(result)
	// Output: HELLO.WORLD.EXAMPLE
}
