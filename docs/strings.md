# strings

String manipulation utilities including case conversions, padding, substring removal, validation, and prefix/suffix matching.

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

### FirstToUpper

```go
func FirstToUpper(s string) string
```

Converts the first character of a string to uppercase.

### FirstToLower

```go
func FirstToLower(s string) string
```

Converts the first character of a string to lowercase.

### Compose

```go
func Compose(s string, funcs ...func(string) string) string
```

Chains multiple string transformation functions, applying each in order.

### Validation

```go
func IsEmpty(s string) bool
func IsBlank(s string) bool
func IsAnyEmpty(s ...string) bool
func IsAnyBlank(strings ...string) bool
func IsAlpha(s string) bool
func IsAlphanumeric(s string) bool
func IsNumeric(s string) bool
func IsNumerical(s string) bool
```

| Function | Description |
|----------|-------------|
| `IsEmpty` | Returns `true` if the string has zero length. |
| `IsBlank` | Returns `true` if the string is empty or contains only whitespace. |
| `IsAnyEmpty` | Returns `true` if any of the provided strings is empty. |
| `IsAnyBlank` | Returns `true` if any of the provided strings is blank. |
| `IsAlpha` | Returns `true` if the string contains only Unicode letters. |
| `IsAlphanumeric` | Returns `true` if the string contains only Unicode letters and digits. |
| `IsNumeric` | Returns `true` if the string contains only Unicode digits. |
| `IsNumerical` | Returns `true` if the string represents a number (digits with an optional decimal point). |

Empty strings return `false` for `IsAlpha`, `IsAlphanumeric`, `IsNumeric`, and `IsNumerical`.

### Prefix / Suffix

```go
func HasAnyPrefix(s string, prefixes ...string) bool
func HasAnySuffix(s string, suffixes ...string) bool
```

Returns `true` if the string starts (or ends) with any of the provided values. Returns `false` when the string is empty or no candidates are given.

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

### Validation

```go
stringsx.IsEmpty("")           // true
stringsx.IsBlank("  \t")      // true
stringsx.IsAnyEmpty("a", "")  // true
stringsx.IsAlpha("Hello")     // true
stringsx.IsAlphanumeric("Go1") // true
stringsx.IsNumeric("12345")   // true
stringsx.IsNumerical("3.14")  // true
```

### Prefix / Suffix

```go
stringsx.HasAnyPrefix("/api/users", "/api", "/admin") // true
stringsx.HasAnySuffix("image.png", ".jpg", ".png")    // true
```
