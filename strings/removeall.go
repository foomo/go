package strings

import (
	"strings"
)

// RemoveAll removes all occurrences of each substring from s in a single pass.
func RemoveAll(s string, substrings ...string) string {
	if len(substrings) == 0 {
		return s
	}

	oldnew := make([]string, 0, len(substrings)*2)
	for _, sub := range substrings {
		oldnew = append(oldnew, sub, "")
	}

	return strings.NewReplacer(oldnew...).Replace(s)
}
