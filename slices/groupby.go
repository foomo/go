package slices

// GroupBy groups elements by key.
func GroupBy[T any, K comparable](items []T, keyFn func(T) K) map[K][]T {
	groups := make(map[K][]T)

	for _, item := range items {
		key := keyFn(item)
		groups[key] = append(groups[key], item)
	}

	return groups
}
