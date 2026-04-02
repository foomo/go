# os

Typed environment variable parsing with defaults.

## Import

```go
import osx "github.com/foomo/go/os"
```

## Package Variables

```go
var SliceSeperator  = ","   // delimiter for slice elements
var MapSeperator    = ","   // delimiter for map key-value pairs
var MapKVSeperator  = ":"   // delimiter between key and value
```

## API

### Environment Check

```go
func HasEnv(key string) bool
func MustHasEnv(key string)             // panics if not defined
```

### Getenv — Scalars

All `Getenv*` functions return the parsed value or the given default if the variable is empty or undefined.

```go
func Getenv(key string, def string) string
func GetenvBool(key string, def bool) (bool, error)
func GetenvInt(key string, def int) (int, error)
func GetenvInt8(key string, def int8) (int8, error)
func GetenvInt16(key string, def int16) (int16, error)
func GetenvInt32(key string, def int32) (int32, error)
func GetenvInt64(key string, def int64) (int64, error)
func GetenvUint(key string, def uint) (uint, error)
func GetenvUint8(key string, def uint8) (uint8, error)
func GetenvUint16(key string, def uint16) (uint16, error)
func GetenvUint32(key string, def uint32) (uint32, error)
func GetenvUint64(key string, def uint64) (uint64, error)
func GetenvFloat32(key string, def float32) (float32, error)
func GetenvFloat64(key string, def float64) (float64, error)
func GetenvDuration(key string, def time.Duration) (time.Duration, error)
```

Integer types support the `0x` hex prefix.

### MustGetenv — Scalars

`Must` variants panic if the variable is not defined or cannot be parsed. No default value is needed.

```go
func MustGetenv(key string) string
func MustGetenvBool(key string) bool
func MustGetenvInt(key string) int
func MustGetenvInt8(key string) int8
func MustGetenvInt16(key string) int16
func MustGetenvInt32(key string) int32
func MustGetenvInt64(key string) int64
func MustGetenvUint(key string) uint
func MustGetenvUint8(key string) uint8
func MustGetenvUint16(key string) uint16
func MustGetenvUint32(key string) uint32
func MustGetenvUint64(key string) uint64
func MustGetenvFloat32(key string) float32
func MustGetenvFloat64(key string) float64
func MustGetenvDuration(key string) time.Duration
```

### Getenv — Slices

Parse comma-separated values into typed slices. Values are space-trimmed.

```go
func GetenvStringSlice(key string, def []string) []string
func GetenvBoolSlice(key string, def []bool) ([]bool, error)
func GetenvIntSlice(key string, def []int) ([]int, error)
func GetenvInt8Slice(key string, def []int8) ([]int8, error)
func GetenvInt16Slice(key string, def []int16) ([]int16, error)
func GetenvInt32Slice(key string, def []int32) ([]int32, error)
func GetenvInt64Slice(key string, def []int64) ([]int64, error)
func GetenvUintSlice(key string, def []uint) ([]uint, error)
func GetenvUint8Slice(key string, def []uint8) ([]uint8, error)
func GetenvUint16Slice(key string, def []uint16) ([]uint16, error)
func GetenvUint32Slice(key string, def []uint32) ([]uint32, error)
func GetenvUint64Slice(key string, def []uint64) ([]uint64, error)
func GetenvFloat32Slice(key string, def []float32) ([]float32, error)
func GetenvFloat64Slice(key string, def []float64) ([]float64, error)
func GetenvDurationSlice(key string, def []time.Duration) ([]time.Duration, error)
```

### Getenv — Maps

Parse comma-separated `key:value` pairs into typed maps. Keys and values are space-trimmed.

```go
func GetenvStringMap(key string, def map[string]string) (map[string]string, error)
func GetenvBoolMap(key string, def map[string]bool) (map[string]bool, error)
func GetenvIntMap(key string, def map[string]int) (map[string]int, error)
func GetenvInt8Map(key string, def map[string]int8) (map[string]int8, error)
func GetenvInt16Map(key string, def map[string]int16) (map[string]int16, error)
func GetenvInt32Map(key string, def map[string]int32) (map[string]int32, error)
func GetenvInt64Map(key string, def map[string]int64) (map[string]int64, error)
func GetenvUintMap(key string, def map[string]uint) (map[string]uint, error)
func GetenvUint8Map(key string, def map[string]uint8) (map[string]uint8, error)
func GetenvUint16Map(key string, def map[string]uint16) (map[string]uint16, error)
func GetenvUint32Map(key string, def map[string]uint32) (map[string]uint32, error)
func GetenvUint64Map(key string, def map[string]uint64) (map[string]uint64, error)
func GetenvFloat32Map(key string, def map[string]float32) (map[string]float32, error)
func GetenvFloat64Map(key string, def map[string]float64) (map[string]float64, error)
func GetenvDurationMap(key string, def map[string]time.Duration) (map[string]time.Duration, error)
```

::: warning Deprecated
`GetenvStringMapString` is deprecated — use `GetenvStringMap` instead.
:::

## Examples

### Basic usage

```go
package main

import (
	"fmt"
	"os"

	osx "github.com/foomo/go/os"
)

func main() {
	// String with default
	host := osx.Getenv("APP_HOST", "localhost")
	fmt.Println(host)

	// Check existence
	if osx.HasEnv("DATABASE_URL") {
		fmt.Println("database configured")
	}

	// Typed parsing
	os.Setenv("DEBUG", "true")
	debug, _ := osx.GetenvBool("DEBUG", false)
	fmt.Println(debug) // true

	os.Setenv("PORT", "8080")
	port, _ := osx.GetenvInt32("PORT", 3000)
	fmt.Println(port) // 8080

	// Duration
	os.Setenv("TIMEOUT", "30s")
	timeout, _ := osx.GetenvDuration("TIMEOUT", 10*time.Second)
	fmt.Println(timeout) // 30s
}
```

### Must variants

```go
// Panics if DATABASE_URL is not set
dsn := osx.MustGetenv("DATABASE_URL")

// Panics if PORT is not set or not a valid int
port := osx.MustGetenvInt("PORT")
```

### Slices and maps

```go
os.Setenv("ALLOWED_ORIGINS", "foo.com, bar.com, baz.com")
origins := osx.GetenvStringSlice("ALLOWED_ORIGINS", nil)
// origins = ["foo.com", "bar.com", "baz.com"]

os.Setenv("PORTS", "8080,8081,8082")
ports, _ := osx.GetenvIntSlice("PORTS", nil)
// ports = [8080, 8081, 8082]

os.Setenv("LABELS", "env:prod, region:eu")
labels, _ := osx.GetenvStringMap("LABELS", nil)
// labels = map["env":"prod", "region":"eu"]

os.Setenv("TIMEOUTS", "read:5s,write:10s")
timeouts, _ := osx.GetenvDurationMap("TIMEOUTS", nil)
// timeouts = map["read":5s, "write":10s]
```
