package runtime_test

import (
	"fmt"

	"github.com/foomo/go/runtime"
)

func ExampleCallerFunc() {
	name, _ := runtime.CallerFunc(0)

	fmt.Println(name)

	// Output:
	// runtime_test.ExampleCaller
	// github.com/foomo/go/runtime_test.ExampleCaller
	// runtime/caller_test.go
	// 18
	// true
}
