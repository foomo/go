package strings

// ToSnake converts a given string to snake_case format using underscores as delimiters.
func ToSnake(s string) string {
	return ToDelimited(s, '_')
}

// ToSnakeWithIgnore converts the input string `s` to snake_case format while ignoring characters specified in `ignore`.
func ToSnakeWithIgnore(s string, ignore string) string {
	return ToScreamingDelimited(s, '_', ignore, false)
}

// ToScreamingSnake converts a given string to SCREAMING_SNAKE_CASE format using underscores as delimiters.
func ToScreamingSnake(s string) string {
	return ToScreamingDelimited(s, '_', "", true)
}
