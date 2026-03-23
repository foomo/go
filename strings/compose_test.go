package strings_test

import (
	"fmt"
	"strings"

	stringsx "github.com/foomo/go/strings"
)

func ExampleCompose() {
	result := stringsx.Compose("HELLO", strings.ToLower, stringsx.FirstToUpper)
	fmt.Println(result)
	// Output: Hello
}
