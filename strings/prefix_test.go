package strings_test

import (
	"fmt"

	"github.com/foomo/go/strings"
)

func ExampleHasAnyPrefix() {
	fmt.Println(strings.HasAnyPrefix("foobar", "foo", "baz"))
	// Output: true
}
