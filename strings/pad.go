package strings

import (
	"strings"
	"unicode/utf8"
)

// PadRight appends spaces to the right of the input string s until it reaches the specified rune length n.
func PadRight(s string, n int) string {
	return s + strings.Repeat(" ", max(0, n-utf8.RuneCountInString(s)))
}

// PadLeft prepends spaces to the left of the input string s until it reaches the specified rune length n.
func PadLeft(s string, n int) string {
	return strings.Repeat(" ", max(0, n-utf8.RuneCountInString(s))) + s
}
