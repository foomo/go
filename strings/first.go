package strings

import (
	"strings"
)

// FirstToUpper converts the first character of a string to uppercase.
func FirstToUpper(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

// FirstToLower converts the first character of a string to lowercase.
func FirstToLower(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}
