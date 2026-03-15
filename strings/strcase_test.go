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
