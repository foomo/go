package strings_test

import (
	"fmt"

	"github.com/foomo/go/strings"
)

func ExampleToSnake() {
	result := strings.ToSnake("HelloWorldExample")
	fmt.Println(result)
	// Output: hello_world_example
}

func ExampleToScreamingSnake() {
	result := strings.ToScreamingSnake("HelloWorldExample")
	fmt.Println(result)
	// Output: HELLO_WORLD_EXAMPLE
}
