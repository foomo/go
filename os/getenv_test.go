package os_test

import (
	"fmt"
	"os"
	"sort"
	"time"

	goos "github.com/foomo/go/os"
)

// ---------------------------------- helpers ----------------------------------

func ExampleHasEnv() {
	_ = os.Unsetenv("FOO")

	fmt.Println(goos.HasEnv("FOO"))

	_ = os.Setenv("FOO", "bar")

	fmt.Println(goos.HasEnv("FOO"))

	// Output:
	// false
	// true
}

func ExampleMustHasEnv() {
	_ = os.Setenv("FOO", "bar")

	goos.MustHasEnv("FOO") // does not panic
	fmt.Println("ok")

	// Output:
	// ok
}

// -------------------------------- scalars -----------------------------------

func ExampleGetenv() {
	_ = os.Setenv("FOO", "")

	fmt.Println(goos.Getenv("FOO", "fallback"))

	_ = os.Setenv("FOO", "bar")

	fmt.Println(goos.Getenv("FOO", "fallback"))

	// Output:
	// fallback
	// bar
}

func ExampleGetenvBool() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvBool("FOO", false)
	fmt.Println(v)

	_ = os.Setenv("FOO", "true")
	v, _ = goos.GetenvBool("FOO", false)
	fmt.Println(v)

	// Output:
	// false
	// true
}

func ExampleGetenvInt() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvInt("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "42")
	v, _ = goos.GetenvInt("FOO", 1)
	fmt.Println(v)

	// Output:
	// 1
	// 42
}

func ExampleGetenvInt8() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvInt8("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "127")
	v, _ = goos.GetenvInt8("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "128")
	_, err := goos.GetenvInt8("FOO", 0)
	fmt.Println(err != nil)

	// Output:
	// 1
	// 127
	// true
}

func ExampleGetenvInt16() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvInt16("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "1000")
	v, _ = goos.GetenvInt16("FOO", 1)
	fmt.Println(v)

	// Output:
	// 1
	// 1000
}

func ExampleGetenvInt32() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvInt32("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "2")
	v, _ = goos.GetenvInt32("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "0x1F")
	v, _ = goos.GetenvInt32("FOO", 0)
	fmt.Println(v)

	_ = os.Setenv("FOO", "2147483648")
	_, err := goos.GetenvInt32("FOO", 0)
	fmt.Println(err != nil)

	// Output:
	// 1
	// 2
	// 31
	// true
}

func ExampleGetenvInt64() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvInt64("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "2")
	v, _ = goos.GetenvInt64("FOO", 1)
	fmt.Println(v)

	// Output:
	// 1
	// 2
}

func ExampleGetenvUint() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvUint("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "42")
	v, _ = goos.GetenvUint("FOO", 1)
	fmt.Println(v)

	// Output:
	// 1
	// 42
}

func ExampleGetenvUint8() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvUint8("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "255")
	v, _ = goos.GetenvUint8("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "256")
	_, err := goos.GetenvUint8("FOO", 0)
	fmt.Println(err != nil)

	// Output:
	// 1
	// 255
	// true
}

func ExampleGetenvUint16() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvUint16("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "65535")
	v, _ = goos.GetenvUint16("FOO", 1)
	fmt.Println(v)

	// Output:
	// 1
	// 65535
}

func ExampleGetenvUint32() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvUint32("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "100000")
	v, _ = goos.GetenvUint32("FOO", 1)
	fmt.Println(v)

	// Output:
	// 1
	// 100000
}

func ExampleGetenvUint64() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvUint64("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "18446744073709551615")
	v, _ = goos.GetenvUint64("FOO", 1)
	fmt.Println(v)

	// Output:
	// 1
	// 18446744073709551615
}

func ExampleGetenvFloat32() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvFloat32("FOO", 0.5)
	fmt.Println(v)

	_ = os.Setenv("FOO", "1.5")
	v, _ = goos.GetenvFloat32("FOO", 0.5)
	fmt.Println(v)

	_ = os.Setenv("FOO", "not-a-number")
	_, err := goos.GetenvFloat32("FOO", 0)
	fmt.Println(err != nil)

	// Output:
	// 0.5
	// 1.5
	// true
}

func ExampleGetenvFloat64() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvFloat64("FOO", 0.1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "0.2")
	v, _ = goos.GetenvFloat64("FOO", 0.1)
	fmt.Println(v)

	// Output:
	// 0.1
	// 0.2
}

func ExampleGetenvDuration() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvDuration("FOO", 5*time.Second)
	fmt.Println(v)

	_ = os.Setenv("FOO", "100ms")
	v, _ = goos.GetenvDuration("FOO", 5*time.Second)
	fmt.Println(v)

	_ = os.Setenv("FOO", "invalid")
	_, err := goos.GetenvDuration("FOO", 0)
	fmt.Println(err != nil)

	// Output:
	// 5s
	// 100ms
	// true
}

// ------------------------------ must scalars --------------------------------

func ExampleMustGetenv() {
	_ = os.Setenv("FOO", "bar")

	fmt.Println(goos.MustGetenv("FOO"))

	// Output:
	// bar
}

func ExampleMustGetenvBool() {
	_ = os.Setenv("FOO", "true")

	fmt.Println(goos.MustGetenvBool("FOO"))

	// Output:
	// true
}

func ExampleMustGetenvInt() {
	_ = os.Setenv("FOO", "42")

	fmt.Println(goos.MustGetenvInt("FOO"))

	// Output:
	// 42
}

func ExampleMustGetenvInt8() {
	_ = os.Setenv("FOO", "127")

	fmt.Println(goos.MustGetenvInt8("FOO"))

	// Output:
	// 127
}

func ExampleMustGetenvInt16() {
	_ = os.Setenv("FOO", "1000")

	fmt.Println(goos.MustGetenvInt16("FOO"))

	// Output:
	// 1000
}

func ExampleMustGetenvInt32() {
	_ = os.Setenv("FOO", "100000")

	fmt.Println(goos.MustGetenvInt32("FOO"))

	// Output:
	// 100000
}

func ExampleMustGetenvInt64() {
	_ = os.Setenv("FOO", "100000")

	fmt.Println(goos.MustGetenvInt64("FOO"))

	// Output:
	// 100000
}

func ExampleMustGetenvUint() {
	_ = os.Setenv("FOO", "42")

	fmt.Println(goos.MustGetenvUint("FOO"))

	// Output:
	// 42
}

func ExampleMustGetenvUint8() {
	_ = os.Setenv("FOO", "255")

	fmt.Println(goos.MustGetenvUint8("FOO"))

	// Output:
	// 255
}

func ExampleMustGetenvUint16() {
	_ = os.Setenv("FOO", "65535")

	fmt.Println(goos.MustGetenvUint16("FOO"))

	// Output:
	// 65535
}

func ExampleMustGetenvUint32() {
	_ = os.Setenv("FOO", "100000")

	fmt.Println(goos.MustGetenvUint32("FOO"))

	// Output:
	// 100000
}

func ExampleMustGetenvUint64() {
	_ = os.Setenv("FOO", "18446744073709551615")

	fmt.Println(goos.MustGetenvUint64("FOO"))

	// Output:
	// 18446744073709551615
}

func ExampleMustGetenvFloat32() {
	_ = os.Setenv("FOO", "1.5")

	fmt.Println(goos.MustGetenvFloat32("FOO"))

	// Output:
	// 1.5
}

func ExampleMustGetenvFloat64() {
	_ = os.Setenv("FOO", "3.14")

	fmt.Println(goos.MustGetenvFloat64("FOO"))

	// Output:
	// 3.14
}

func ExampleMustGetenvDuration() {
	_ = os.Setenv("FOO", "5s")

	fmt.Println(goos.MustGetenvDuration("FOO"))

	// Output:
	// 5s
}

// --------------------------------- slices -----------------------------------

func ExampleGetenvStringSlice() {
	_ = os.Setenv("FOO", "")

	fmt.Println(goos.GetenvStringSlice("FOO", nil))

	_ = os.Setenv("FOO", "foo")

	fmt.Println(goos.GetenvStringSlice("FOO", nil))

	_ = os.Setenv("FOO", "foo,bar")

	fmt.Println(goos.GetenvStringSlice("FOO", nil))

	// Output:
	// []
	// [foo]
	// [foo bar]
}

func ExampleGetenvBoolSlice() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvBoolSlice("FOO", nil)
	fmt.Println(v)

	_ = os.Setenv("FOO", "true,false,true")
	v, _ = goos.GetenvBoolSlice("FOO", nil)
	fmt.Println(v)

	// Output:
	// []
	// [true false true]
}

func ExampleGetenvIntSlice() {
	_ = os.Setenv("FOO", "")
	v, _ := goos.GetenvIntSlice("FOO", nil)
	fmt.Println(v)

	_ = os.Setenv("FOO", "1, 2, 3")
	v, _ = goos.GetenvIntSlice("FOO", nil)
	fmt.Println(v)

	// Output:
	// []
	// [1 2 3]
}

func ExampleGetenvInt8Slice() {
	_ = os.Setenv("FOO", "1, 2, 3")
	v, _ := goos.GetenvInt8Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1 2 3]
}

func ExampleGetenvInt16Slice() {
	_ = os.Setenv("FOO", "100, 200, 300")
	v, _ := goos.GetenvInt16Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [100 200 300]
}

func ExampleGetenvInt32Slice() {
	_ = os.Setenv("FOO", "100, 200, 300")
	v, _ := goos.GetenvInt32Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [100 200 300]
}

func ExampleGetenvInt64Slice() {
	_ = os.Setenv("FOO", "100, 200, 300")
	v, _ := goos.GetenvInt64Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [100 200 300]
}

func ExampleGetenvUintSlice() {
	_ = os.Setenv("FOO", "1, 2, 3")
	v, _ := goos.GetenvUintSlice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1 2 3]
}

func ExampleGetenvUint8Slice() {
	_ = os.Setenv("FOO", "1, 2, 3")
	v, _ := goos.GetenvUint8Slice("FOO", nil)
	fmt.Println(v) //nolint:staticcheck // QF1010

	// Output:
	// [1 2 3]
}

func ExampleGetenvUint16Slice() {
	_ = os.Setenv("FOO", "1, 2, 3")
	v, _ := goos.GetenvUint16Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1 2 3]
}

func ExampleGetenvUint32Slice() {
	_ = os.Setenv("FOO", "1, 2, 3")
	v, _ := goos.GetenvUint32Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1 2 3]
}

func ExampleGetenvUint64Slice() {
	_ = os.Setenv("FOO", "1, 2, 3")
	v, _ := goos.GetenvUint64Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1 2 3]
}

func ExampleGetenvFloat32Slice() {
	_ = os.Setenv("FOO", "1.1, 2.2, 3.3")
	v, _ := goos.GetenvFloat32Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1.1 2.2 3.3]
}

func ExampleGetenvFloat64Slice() {
	_ = os.Setenv("FOO", "1.1, 2.2, 3.3")
	v, _ := goos.GetenvFloat64Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1.1 2.2 3.3]
}

func ExampleGetenvDurationSlice() {
	_ = os.Setenv("FOO", "1s, 500ms, 2m")
	v, _ := goos.GetenvDurationSlice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1s 500ms 2m0s]
}

// ---------------------------------- maps ------------------------------------

func ExampleGetenvStringMap() {
	_ = os.Setenv("FOO", "a:1")
	v, _ := goos.GetenvStringMap("FOO", nil)
	fmt.Println(v)

	_ = os.Setenv("FOO", " x : hello , y : world ")
	v, _ = goos.GetenvStringMap("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%s\n", k, v[k])
	}

	_ = os.Setenv("FOO", "invalid")
	_, err := goos.GetenvStringMap("FOO", nil)
	fmt.Println(err != nil)

	// Output:
	// map[a:1]
	// x:hello
	// y:world
	// true
}

func ExampleGetenvStringMapString() {
	_ = os.Setenv("FOO", "a:1")
	v, _ := goos.GetenvStringMapString("FOO", nil)
	fmt.Println(v)

	// Output:
	// map[a:1]
}

func ExampleGetenvBoolMap() {
	_ = os.Setenv("FOO", "debug:true, verbose:false")
	v, _ := goos.GetenvBoolMap("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%v\n", k, v[k])
	}

	// Output:
	// debug:true
	// verbose:false
}

func ExampleGetenvIntMap() {
	_ = os.Setenv("FOO", "a:1, b:2")
	v, _ := goos.GetenvIntMap("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%d\n", k, v[k])
	}

	// Output:
	// a:1
	// b:2
}

func ExampleGetenvInt8Map() {
	_ = os.Setenv("FOO", "a:1, b:2")
	v, _ := goos.GetenvInt8Map("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%d\n", k, v[k])
	}

	// Output:
	// a:1
	// b:2
}

func ExampleGetenvInt16Map() {
	_ = os.Setenv("FOO", "a:100, b:200")
	v, _ := goos.GetenvInt16Map("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%d\n", k, v[k])
	}

	// Output:
	// a:100
	// b:200
}

func ExampleGetenvInt32Map() {
	_ = os.Setenv("FOO", "a:100, b:200")
	v, _ := goos.GetenvInt32Map("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%d\n", k, v[k])
	}

	// Output:
	// a:100
	// b:200
}

func ExampleGetenvInt64Map() {
	_ = os.Setenv("FOO", "a:100, b:200")
	v, _ := goos.GetenvInt64Map("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%d\n", k, v[k])
	}

	// Output:
	// a:100
	// b:200
}

func ExampleGetenvUintMap() {
	_ = os.Setenv("FOO", "a:1, b:2")
	v, _ := goos.GetenvUintMap("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%d\n", k, v[k])
	}

	// Output:
	// a:1
	// b:2
}

func ExampleGetenvUint8Map() {
	_ = os.Setenv("FOO", "a:1, b:2")
	v, _ := goos.GetenvUint8Map("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%d\n", k, v[k])
	}

	// Output:
	// a:1
	// b:2
}

func ExampleGetenvUint16Map() {
	_ = os.Setenv("FOO", "a:1, b:2")
	v, _ := goos.GetenvUint16Map("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%d\n", k, v[k])
	}

	// Output:
	// a:1
	// b:2
}

func ExampleGetenvUint32Map() {
	_ = os.Setenv("FOO", "a:1, b:2")
	v, _ := goos.GetenvUint32Map("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%d\n", k, v[k])
	}

	// Output:
	// a:1
	// b:2
}

func ExampleGetenvUint64Map() {
	_ = os.Setenv("FOO", "a:1, b:2")
	v, _ := goos.GetenvUint64Map("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%d\n", k, v[k])
	}

	// Output:
	// a:1
	// b:2
}

func ExampleGetenvFloat32Map() {
	_ = os.Setenv("FOO", "a:1.5, b:2.5")
	v, _ := goos.GetenvFloat32Map("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%.1f\n", k, v[k])
	}

	// Output:
	// a:1.5
	// b:2.5
}

func ExampleGetenvFloat64Map() {
	_ = os.Setenv("FOO", "a:1.5, b:2.5")
	v, _ := goos.GetenvFloat64Map("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%.1f\n", k, v[k])
	}

	// Output:
	// a:1.5
	// b:2.5
}

func ExampleGetenvDurationMap() {
	_ = os.Setenv("FOO", "timeout:5s, interval:100ms")
	v, _ := goos.GetenvDurationMap("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%v\n", k, v[k])
	}

	// Output:
	// interval:100ms
	// timeout:5s
}
