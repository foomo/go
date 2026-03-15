# os

Typed environment variable parsing with defaults.

## Import

```go
import osx "github.com/foomo/go/os"
```

## API

### HasEnv

```go
func HasEnv(key string) bool
```

Returns `true` if the environment variable is defined.

### MustHasEnv

```go
func MustHasEnv(key string)
```

Panics if the environment variable is not defined.

### Getenv

```go
func Getenv(key string, def string) string
```

Returns the environment variable value, or the default if empty or undefined.

### GetenvBool

```go
func GetenvBool(key string, def bool) (bool, error)
```

Parses the environment variable as a boolean, or returns the default.

### GetenvInt32

```go
func GetenvInt32(key string, def int32) (int32, error)
```

Parses the environment variable as an `int32` (supports `0x` hex prefix), or returns the default.

### GetenvInt64

```go
func GetenvInt64(key string, def int64) (int64, error)
```

Parses the environment variable as an `int64`, or returns the default.

### GetenvFloat32

```go
func GetenvFloat32(key string, def float32) (float32, error)
```

Parses the environment variable as a `float32`, or returns the default.

### GetenvFloat64

```go
func GetenvFloat64(key string, def float64) (float64, error)
```

Parses the environment variable as a `float64`, or returns the default.

### GetenvStringSlice

```go
func GetenvStringSlice(key string, def []string) []string
```

Parses a comma-separated environment variable into a string slice. Values are space-trimmed.

### GetenvStringMapString

```go
func GetenvStringMapString(key string, def map[string]string) (map[string]string, error)
```

Parses a comma-separated list of `key:value` pairs into a map. Keys and values are space-trimmed.

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
}
```

### Slices and maps

```go
os.Setenv("ALLOWED_ORIGINS", "foo.com, bar.com, baz.com")
origins := osx.GetenvStringSlice("ALLOWED_ORIGINS", nil)
// origins = ["foo.com", "bar.com", "baz.com"]

os.Setenv("LABELS", "env:prod, region:eu")
labels, _ := osx.GetenvStringMapString("LABELS", nil)
// labels = map["env":"prod", "region":"eu"]
```
