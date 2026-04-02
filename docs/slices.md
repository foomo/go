# slices

Generic slice utility functions.

## Import

```go
import "github.com/foomo/go/slices"
```

## API

### Filter

```go
func Filter[T any](items []T, fn func(T) bool) []T
```

Returns a new slice containing only the elements for which the predicate returns `true`.

### Map

```go
func Map[T, U any](items []T, fn func(T) U) []U
```

Returns a new slice with each element transformed by the given function.

### GroupBy

```go
func GroupBy[T any, K comparable](items []T, keyFn func(T) K) map[K][]T
```

Groups elements into a map keyed by the result of the key function.

### FilterE

```go
func FilterE[T any](items []T, fn func(T) (bool, error)) ([]T, error)
```

Like `Filter`, but the predicate can return an error. Stops and returns the first error encountered.

### MapE

```go
func MapE[T, U any](items []T, fn func(T) (U, error)) ([]U, error)
```

Like `Map`, but the transform function can return an error. Stops and returns the first error encountered.

## Examples

### Filter

```go
evens := slices.Filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
	return n%2 == 0
})
fmt.Println(evens) // [2 4]
```

### Map

```go
doubled := slices.Map([]int{1, 2, 3}, func(n int) int {
	return n * 2
})
fmt.Println(doubled) // [2 4 6]
```

### GroupBy

```go
type item struct {
	name     string
	category string
}

items := []item{
	{"apple", "fruit"},
	{"carrot", "vegetable"},
	{"banana", "fruit"},
}

groups := slices.GroupBy(items, func(i item) string {
	return i.category
})

fmt.Println(len(groups))              // 2
fmt.Println(len(groups["fruit"]))     // 2
fmt.Println(len(groups["vegetable"])) // 1
```
