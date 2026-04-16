package runtime_test

import (
	"fmt"

	"github.com/foomo/go/runtime"
)

func ExampleCallerFunc() {
	name, _ := runtime.CallerFunc(0)

	fmt.Println(name)

	// Output: ExampleCallerFunc
}
