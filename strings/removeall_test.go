package strings_test

import (
	"fmt"

	"github.com/foomo/go/strings"
)

func ExampleRemoveAll() {
	result := strings.RemoveAll("hello world", "o", "l")
	fmt.Println(result)
	// Output: he wrd
}
