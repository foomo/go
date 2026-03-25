package slices

// Map transforms each element using fn.
func Map[T, U any](items []T, fn func(T) U) []U {
	out := make([]U, len(items))
	for i, item := range items {
		out[i] = fn(item)
	}

	return out
}

// MapE transforms each element using fn and returns an error if any.
func MapE[T, U any](items []T, fn func(T) (U, error)) ([]U, error) {
	var err error

	out := make([]U, len(items))
	for i, item := range items {
		out[i], err = fn(item)
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}
