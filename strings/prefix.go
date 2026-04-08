package strings

import (
	"slices"
	"strings"
)

// HasAnyPrefix checks if the given string has any of the provided prefixes and returns true if so.
func HasAnyPrefix(s string, prefixes ...string) bool {
	if IsEmpty(s) || len(prefixes) == 0 {
		return false
	}

	if slices.Contains(prefixes, s) {
		return true
	}

	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}

	return false
}
