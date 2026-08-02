package strings_test

import (
	"fmt"
	"strings"

	gostrings "github.com/foomo/go/strings"
)

func ExampleCompose() {
	result := gostrings.Compose("HELLO", strings.ToLower, gostrings.FirstToUpper)
	fmt.Println(result)
	// Output: Hello
}
