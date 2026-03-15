# strings

String manipulation utilities including case conversions, padding, and substring removal.

## Import

```go
import stringsx "github.com/foomo/go/strings"
```

## API

### Case Conversions

```go
func ToSnake(s string) string
func ToSnakeWithIgnore(s string, ignore string) string
func ToScreamingSnake(s string) string
func ToKebab(s string) string
func ToScreamingKebab(s string) string
func ToCamel(s string) string
func ToLowerCamel(s string) string
func ToDelimited(s string, delimiter rune) string
func ToScreamingDelimited(s string, delimiter rune, acronym string, screaming bool) string
```

Re-exported from [strcase](https://github.com/iancoleman/strcase) for convenience.

| Function | Input | Output |
|----------|-------|--------|
| `ToSnake` | `HelloWorldExample` | `hello_world_example` |
| `ToScreamingSnake` | `HelloWorldExample` | `HELLO_WORLD_EXAMPLE` |
| `ToKebab` | `HelloWorldExample` | `hello-world-example` |
| `ToScreamingKebab` | `HelloWorldExample` | `HELLO-WORLD-EXAMPLE` |
| `ToCamel` | `hello_world_example` | `HelloWorldExample` |
| `ToLowerCamel` | `hello_world_example` | `helloWorldExample` |
| `ToDelimited` | `HelloWorldExample` (`.`) | `hello.world.example` |
| `ToScreamingDelimited` | `HelloWorldExample` (`.`) | `HELLO.WORLD.EXAMPLE` |

### Padding

```go
func PadRight(s string, n int) string
func PadLeft(s string, n int) string
```

Pad a string with spaces to reach the specified rune length. If the string is already long enough, it is returned unchanged.

### RemoveAll

```go
func RemoveAll(s string, substrings ...string) string
```

Removes all occurrences of the given substrings in a single pass.

## Examples

### Case conversions

```go
stringsx.ToCamel("hello_world")         // "HelloWorld"
stringsx.ToSnake("HelloWorld")           // "hello_world"
stringsx.ToKebab("HelloWorld")           // "hello-world"
stringsx.ToScreamingSnake("HelloWorld")  // "HELLO_WORLD"
```

### Padding

```go
fmt.Printf("'%s'", stringsx.PadRight("hello", 10))
// Output: 'hello     '

fmt.Printf("'%s'", stringsx.PadLeft("hello", 10))
// Output: '     hello'
```

### RemoveAll

```go
result := stringsx.RemoveAll("hello world", "o", "l")
fmt.Println(result) // "he wrd"
```
