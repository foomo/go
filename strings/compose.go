package strings

// Compose chains multiple functions to a string.
func Compose(s string, funcs ...func(string) string) string {
	for _, fn := range funcs {
		s = fn(s)
	}

	return s
}
