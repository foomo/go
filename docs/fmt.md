# fmt

Template string formatting utility.

## Import

```go
import fmtx "github.com/foomo/go/fmt"
```

## API

### Tprintf

```go
func Tprintf(format string, pairs ...string) string
```

Formats a string using template syntax with `%{.key}` delimiters. Takes a format string and variadic key-value pairs.

## Example

```go
package main

import (
	"fmt"

	fmtx "github.com/foomo/go/fmt"
)

func main() {
	result := fmtx.Tprintf("%{.name} is %{.age} years old", "name", "John", "age", "30")
	fmt.Println(result)
	// Output: John is 30 years old
}
```
