package strings

import (
	"strings"
)

// toCamelInitCase converts a string to CamelCase or lowerCamelCase based on the initCase parameter.
func toCamelInitCase(s string, initCase bool) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}

	n := strings.Builder{}
	n.Grow(len(s))

	capNext := initCase

	for i := 0; i < len(s); i++ {
		v := s[i]
		vIsCap := v >= 'A' && v <= 'Z'

		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v -= 'a' - 'A'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a' - 'A'
			}
		}

		if vIsCap || vIsLow {
			n.WriteByte(v)

			capNext = false
		} else if v >= '0' && v <= '9' {
			n.WriteByte(v)

			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}

	return n.String()
}

// ToCamel converts a string to CamelCase format by removing delimiters and capitalizing appropriate letters.
func ToCamel(s string) string {
	return toCamelInitCase(s, true)
}

// ToLowerCamel converts a string to lowerCamelCase format by removing delimiters and capitalizing appropriate letters.
func ToLowerCamel(s string) string {
	return toCamelInitCase(s, false)
}
