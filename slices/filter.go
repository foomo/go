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

// FilterE returns elements where fn returns true
func FilterE[T any](items []T, fn func(T) (bool, error)) ([]T, error) {
	out := make([]T, 0, len(items))
	for _, item := range items {
		ok, err := fn(item)
		if err != nil {
			return nil, err
		}

		if ok {
			out = append(out, item)
		}
	}

	return out, nil
}
