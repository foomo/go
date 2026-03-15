package slices

// Filter returns elements where fn returns true.
func Filter[T any](items []T, fn func(T) bool) []T {
	out := make([]T, 0, len(items))
	for _, item := range items {
		if fn(item) {
			out = append(out, item)
		}
	}
	return out
}
