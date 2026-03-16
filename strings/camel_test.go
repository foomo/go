package strings_test

import (
	"fmt"

	"github.com/foomo/go/strings"
)

func ExampleToCamel() {
	result := strings.ToCamel("hello_world_example")
	fmt.Println(result)
	// Output: HelloWorldExample
}

func ExampleToLowerCamel() {
	result := strings.ToLowerCamel("hello_world_example")
	fmt.Println(result)
	// Output: helloWorldExample
}
