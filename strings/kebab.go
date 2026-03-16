package strings

// ToKebab converts a string to kebab-case format using hyphens as delimiters.
func ToKebab(s string) string {
	return ToDelimited(s, '-')
}

// ToScreamingKebab converts a string to SCREAMING-KEBAB-CASE format using hyphens as delimiters.
func ToScreamingKebab(s string) string {
	return ToScreamingDelimited(s, '-', "", true)
}
