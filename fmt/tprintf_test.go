package fmt_test

import (
	"fmt"

	fmtx "github.com/foomo/go/fmt"
)

func ExampleTprintf() {
	format := "%{.name} is %{.age} years old"
	fmt.Println(fmtx.Tprintf(format, "name", "John", "age", "30"))

	// Output: John is 30 years old
}
