package os_test

import (
	"fmt"
	"os"
	"sort"

	osx "github.com/foomo/go/os"
)

func ExampleHasEnv() {
	_ = os.Unsetenv("FOO")

	fmt.Println(osx.HasEnv("FOO"))

	_ = os.Setenv("FOO", "bar")

	fmt.Println(osx.HasEnv("FOO"))

	// Output:
	// false
	// true
}

func ExampleMustHasEnv() {
	_ = os.Setenv("FOO", "bar")

	osx.MustHasEnv("FOO") // does not panic
	fmt.Println("ok")

	// Output:
	// ok
}

func ExampleGetenv() {
	_ = os.Setenv("FOO", "")

	fmt.Println(osx.Getenv("FOO", "fallback"))

	_ = os.Setenv("FOO", "bar")

	fmt.Println(osx.Getenv("FOO", "fallback"))

	// Output:
	// fallback
	// bar
}

func ExampleGetenvBool() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvBool("FOO", false)
	fmt.Println(v)

	_ = os.Setenv("FOO", "true")
	v, _ = osx.GetenvBool("FOO", false)
	fmt.Println(v)

	// Output:
	// false
	// true
}

func ExampleGetenvInt64() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvInt64("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "2")
	v, _ = osx.GetenvInt64("FOO", 1)
	fmt.Println(v)

	// Output:
	// 1
	// 2
}

func ExampleGetenvInt32() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvInt32("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "2")
	v, _ = osx.GetenvInt32("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "0x1F")
	v, _ = osx.GetenvInt32("FOO", 0)
	fmt.Println(v)

	_ = os.Setenv("FOO", "2147483648")
	_, err := osx.GetenvInt32("FOO", 0)
	fmt.Println(err != nil)

	// Output:
	// 1
	// 2
	// 31
	// true
}

func ExampleGetenvFloat64() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvFloat64("FOO", 0.1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "0.2")
	v, _ = osx.GetenvFloat64("FOO", 0.1)
	fmt.Println(v)

	// Output:
	// 0.1
	// 0.2
}

func ExampleGetenvFloat32() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvFloat32("FOO", 0.5)
	fmt.Println(v)

	_ = os.Setenv("FOO", "1.5")
	v, _ = osx.GetenvFloat32("FOO", 0.5)
	fmt.Println(v)

	_ = os.Setenv("FOO", "not-a-number")
	_, err := osx.GetenvFloat32("FOO", 0)
	fmt.Println(err != nil)

	// Output:
	// 0.5
	// 1.5
	// true
}

func ExampleGetenvStringSlice() {
	_ = os.Setenv("FOO", "")

	fmt.Println(osx.GetenvStringSlice("FOO", nil))

	_ = os.Setenv("FOO", "foo")

	fmt.Println(osx.GetenvStringSlice("FOO", nil))

	_ = os.Setenv("FOO", "foo,bar")

	fmt.Println(osx.GetenvStringSlice("FOO", nil))

	// Output:
	// []
	// [foo]
	// [foo bar]
}

func ExampleGetenvStringMapString() {
	_ = os.Setenv("FOO", "a:1")
	v, _ := osx.GetenvStringMapString("FOO", nil)
	fmt.Println(v)

	_ = os.Setenv("FOO", " x : hello , y : world ")
	v, _ = osx.GetenvStringMapString("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%s\n", k, v[k])
	}

	_ = os.Setenv("FOO", "invalid")
	_, err := osx.GetenvStringMapString("FOO", nil)
	fmt.Println(err != nil)

	// Output:
	// map[a:1]
	// x:hello
	// y:world
	// true
}
