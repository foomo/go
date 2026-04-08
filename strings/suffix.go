package strings

import (
	"slices"
	"strings"
)

// HasAnySuffix checks if the given string has any of the provided suffixes and returns true if so.
func HasAnySuffix(s string, suffixes ...string) bool {
	if IsEmpty(s) || len(suffixes) == 0 {
		return false
	}

	if slices.Contains(suffixes, s) {
		return true
	}

	for _, suffix := range suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}

	return false
}
