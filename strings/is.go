package strings

import (
	"regexp"
	"slices"
	"sync"
	"unicode"
)

var isNumericalRegex = sync.OnceValue(func() *regexp.Regexp {
	return regexp.MustCompile(`^\d+.?\d*$`)
})

// IsEmpty checks if the given string is empty and returns true if it is, otherwise returns false.
func IsEmpty(s string) bool {
	return s == ""
}

// IsBlank checks if the given string is blank (contains only whitespace characters) and returns true if it is, otherwise returns false.
func IsBlank(s string) bool {
	if s == "" {
		return true
	}

	for _, c := range s {
		if !unicode.IsSpace(c) {
			return false
		}
	}

	return true
}

// IsAnyEmpty checks if any of the provided strings in the variadic argument is empty and returns true if so.
func IsAnyEmpty(s ...string) bool {
	if len(s) == 0 {
		return true
	}

	return slices.Contains(s, "")
}

// IsAnyBlank checks if any of the provided strings in the variadic argument is blank and returns true if so.
func IsAnyBlank(strings ...string) bool {
	if len(strings) == 0 {
		return true
	}

	return slices.ContainsFunc(strings, IsBlank)
}

// IsAlpha checks if the given string contains only alphabetic characters and returns true if it is, otherwise returns false.
func IsAlpha(s string) bool {
	if s == "" {
		return false
	}

	for _, v := range s {
		if !unicode.IsLetter(v) {
			return false
		}
	}

	return true
}

// IsAlphanumeric checks if the given string is alphanumeric and returns true if it is, otherwise returns false.
func IsAlphanumeric(s string) bool {
	if s == "" {
		return false
	}

	for _, v := range s {
		if !isAlphanumeric(v) {
			return false
		}
	}

	return true
}

// IsNumeric checks if the given string is numeric and returns true if it is, otherwise returns false.
func IsNumeric(s string) bool {
	if s == "" {
		return false
	}

	for _, v := range s {
		if !unicode.IsDigit(v) {
			return false
		}
	}

	return true
}

// IsNumerical checks if the given string is numerical and returns true if it is, otherwise returns false.
func IsNumerical(s string) bool {
	return isNumericalRegex().MatchString(s)
}

func isAlphanumeric(v rune) bool {
	return unicode.IsDigit(v) || unicode.IsLetter(v)
}
