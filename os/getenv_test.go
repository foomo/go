package os_test

import (
	"fmt"
	"os"
	"sort"
	"time"

	osx "github.com/foomo/go/os"
)

// ---------------------------------- helpers ----------------------------------

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

// -------------------------------- scalars -----------------------------------

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

func ExampleGetenvInt() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvInt("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "42")
	v, _ = osx.GetenvInt("FOO", 1)
	fmt.Println(v)

	// Output:
	// 1
	// 42
}

func ExampleGetenvInt8() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvInt8("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "127")
	v, _ = osx.GetenvInt8("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "128")
	_, err := osx.GetenvInt8("FOO", 0)
	fmt.Println(err != nil)

	// Output:
	// 1
	// 127
	// true
}

func ExampleGetenvInt16() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvInt16("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "1000")
	v, _ = osx.GetenvInt16("FOO", 1)
	fmt.Println(v)

	// Output:
	// 1
	// 1000
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

func ExampleGetenvUint() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvUint("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "42")
	v, _ = osx.GetenvUint("FOO", 1)
	fmt.Println(v)

	// Output:
	// 1
	// 42
}

func ExampleGetenvUint8() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvUint8("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "255")
	v, _ = osx.GetenvUint8("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "256")
	_, err := osx.GetenvUint8("FOO", 0)
	fmt.Println(err != nil)

	// Output:
	// 1
	// 255
	// true
}

func ExampleGetenvUint16() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvUint16("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "65535")
	v, _ = osx.GetenvUint16("FOO", 1)
	fmt.Println(v)

	// Output:
	// 1
	// 65535
}

func ExampleGetenvUint32() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvUint32("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "100000")
	v, _ = osx.GetenvUint32("FOO", 1)
	fmt.Println(v)

	// Output:
	// 1
	// 100000
}

func ExampleGetenvUint64() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvUint64("FOO", 1)
	fmt.Println(v)

	_ = os.Setenv("FOO", "18446744073709551615")
	v, _ = osx.GetenvUint64("FOO", 1)
	fmt.Println(v)

	// Output:
	// 1
	// 18446744073709551615
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

func ExampleGetenvDuration() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvDuration("FOO", 5*time.Second)
	fmt.Println(v)

	_ = os.Setenv("FOO", "100ms")
	v, _ = osx.GetenvDuration("FOO", 5*time.Second)
	fmt.Println(v)

	_ = os.Setenv("FOO", "invalid")
	_, err := osx.GetenvDuration("FOO", 0)
	fmt.Println(err != nil)

	// Output:
	// 5s
	// 100ms
	// true
}

// ------------------------------ must scalars --------------------------------

func ExampleMustGetenv() {
	_ = os.Setenv("FOO", "bar")

	fmt.Println(osx.MustGetenv("FOO"))

	// Output:
	// bar
}

func ExampleMustGetenvBool() {
	_ = os.Setenv("FOO", "true")

	fmt.Println(osx.MustGetenvBool("FOO"))

	// Output:
	// true
}

func ExampleMustGetenvInt() {
	_ = os.Setenv("FOO", "42")

	fmt.Println(osx.MustGetenvInt("FOO"))

	// Output:
	// 42
}

func ExampleMustGetenvInt8() {
	_ = os.Setenv("FOO", "127")

	fmt.Println(osx.MustGetenvInt8("FOO"))

	// Output:
	// 127
}

func ExampleMustGetenvInt16() {
	_ = os.Setenv("FOO", "1000")

	fmt.Println(osx.MustGetenvInt16("FOO"))

	// Output:
	// 1000
}

func ExampleMustGetenvInt32() {
	_ = os.Setenv("FOO", "100000")

	fmt.Println(osx.MustGetenvInt32("FOO"))

	// Output:
	// 100000
}

func ExampleMustGetenvInt64() {
	_ = os.Setenv("FOO", "100000")

	fmt.Println(osx.MustGetenvInt64("FOO"))

	// Output:
	// 100000
}

func ExampleMustGetenvUint() {
	_ = os.Setenv("FOO", "42")

	fmt.Println(osx.MustGetenvUint("FOO"))

	// Output:
	// 42
}

func ExampleMustGetenvUint8() {
	_ = os.Setenv("FOO", "255")

	fmt.Println(osx.MustGetenvUint8("FOO"))

	// Output:
	// 255
}

func ExampleMustGetenvUint16() {
	_ = os.Setenv("FOO", "65535")

	fmt.Println(osx.MustGetenvUint16("FOO"))

	// Output:
	// 65535
}

func ExampleMustGetenvUint32() {
	_ = os.Setenv("FOO", "100000")

	fmt.Println(osx.MustGetenvUint32("FOO"))

	// Output:
	// 100000
}

func ExampleMustGetenvUint64() {
	_ = os.Setenv("FOO", "18446744073709551615")

	fmt.Println(osx.MustGetenvUint64("FOO"))

	// Output:
	// 18446744073709551615
}

func ExampleMustGetenvFloat32() {
	_ = os.Setenv("FOO", "1.5")

	fmt.Println(osx.MustGetenvFloat32("FOO"))

	// Output:
	// 1.5
}

func ExampleMustGetenvFloat64() {
	_ = os.Setenv("FOO", "3.14")

	fmt.Println(osx.MustGetenvFloat64("FOO"))

	// Output:
	// 3.14
}

func ExampleMustGetenvDuration() {
	_ = os.Setenv("FOO", "5s")

	fmt.Println(osx.MustGetenvDuration("FOO"))

	// Output:
	// 5s
}

// --------------------------------- slices -----------------------------------

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

func ExampleGetenvBoolSlice() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvBoolSlice("FOO", nil)
	fmt.Println(v)

	_ = os.Setenv("FOO", "true,false,true")
	v, _ = osx.GetenvBoolSlice("FOO", nil)
	fmt.Println(v)

	// Output:
	// []
	// [true false true]
}

func ExampleGetenvIntSlice() {
	_ = os.Setenv("FOO", "")
	v, _ := osx.GetenvIntSlice("FOO", nil)
	fmt.Println(v)

	_ = os.Setenv("FOO", "1, 2, 3")
	v, _ = osx.GetenvIntSlice("FOO", nil)
	fmt.Println(v)

	// Output:
	// []
	// [1 2 3]
}

func ExampleGetenvInt8Slice() {
	_ = os.Setenv("FOO", "1, 2, 3")
	v, _ := osx.GetenvInt8Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1 2 3]
}

func ExampleGetenvInt16Slice() {
	_ = os.Setenv("FOO", "100, 200, 300")
	v, _ := osx.GetenvInt16Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [100 200 300]
}

func ExampleGetenvInt32Slice() {
	_ = os.Setenv("FOO", "100, 200, 300")
	v, _ := osx.GetenvInt32Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [100 200 300]
}

func ExampleGetenvInt64Slice() {
	_ = os.Setenv("FOO", "100, 200, 300")
	v, _ := osx.GetenvInt64Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [100 200 300]
}

func ExampleGetenvUintSlice() {
	_ = os.Setenv("FOO", "1, 2, 3")
	v, _ := osx.GetenvUintSlice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1 2 3]
}

func ExampleGetenvUint8Slice() {
	_ = os.Setenv("FOO", "1, 2, 3")
	v, _ := osx.GetenvUint8Slice("FOO", nil)
	fmt.Println(string(v))

	// Output:
	// [1 2 3]
}

func ExampleGetenvUint16Slice() {
	_ = os.Setenv("FOO", "1, 2, 3")
	v, _ := osx.GetenvUint16Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1 2 3]
}

func ExampleGetenvUint32Slice() {
	_ = os.Setenv("FOO", "1, 2, 3")
	v, _ := osx.GetenvUint32Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1 2 3]
}

func ExampleGetenvUint64Slice() {
	_ = os.Setenv("FOO", "1, 2, 3")
	v, _ := osx.GetenvUint64Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1 2 3]
}

func ExampleGetenvFloat32Slice() {
	_ = os.Setenv("FOO", "1.1, 2.2, 3.3")
	v, _ := osx.GetenvFloat32Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1.1 2.2 3.3]
}

func ExampleGetenvFloat64Slice() {
	_ = os.Setenv("FOO", "1.1, 2.2, 3.3")
	v, _ := osx.GetenvFloat64Slice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1.1 2.2 3.3]
}

func ExampleGetenvDurationSlice() {
	_ = os.Setenv("FOO", "1s, 500ms, 2m")
	v, _ := osx.GetenvDurationSlice("FOO", nil)
	fmt.Println(v)

	// Output:
	// [1s 500ms 2m0s]
}

// ---------------------------------- maps ------------------------------------

func ExampleGetenvStringMap() {
	_ = os.Setenv("FOO", "a:1")
	v, _ := osx.GetenvStringMap("FOO", nil)
	fmt.Println(v)

	_ = os.Setenv("FOO", " x : hello , y : world ")
	v, _ = osx.GetenvStringMap("FOO", nil)

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s:%s\n", k, v[k])
	}

	_ = os.Setenv("FOO", "invalid")
	_, err := osx.GetenvStringMap("FOO", nil)
	fmt.Println(err != nil)

	// Output:
	// map[a:1]
	// x:hello
	// y:world
	// true
}

func ExampleGetenvStringMapString() {
	_ = os.Setenv("FOO", "a:1")
	v, _ := osx.GetenvStringMapString("FOO", nil)
	fmt.Println(v)

	// Output:
	// map[a:1]
}

func ExampleGetenvBoolMap() {
	_ = os.Setenv("FOO", "debug:true, verbose:false")
	v, _ := osx.GetenvBoolMap("FOO", nil)

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
	v, _ := osx.GetenvIntMap("FOO", nil)

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
	v, _ := osx.GetenvInt8Map("FOO", nil)

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
	v, _ := osx.GetenvInt16Map("FOO", nil)

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
	v, _ := osx.GetenvInt32Map("FOO", nil)

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
	v, _ := osx.GetenvInt64Map("FOO", nil)

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
	v, _ := osx.GetenvUintMap("FOO", nil)

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
	v, _ := osx.GetenvUint8Map("FOO", nil)

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
	v, _ := osx.GetenvUint16Map("FOO", nil)

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
	v, _ := osx.GetenvUint32Map("FOO", nil)

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
	v, _ := osx.GetenvUint64Map("FOO", nil)

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
	v, _ := osx.GetenvFloat32Map("FOO", nil)

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
	v, _ := osx.GetenvFloat64Map("FOO", nil)

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
	v, _ := osx.GetenvDurationMap("FOO", nil)

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
