package strings_test

import (
	"fmt"

	"github.com/foomo/go/strings"
)

func ExampleHasAnySuffix() {
	fmt.Println(strings.HasAnySuffix("foobar", "bar", "baz"))
	// Output: true
}
