package strings

import (
	"strings"
)

// ToDelimited converts a string to delimited.snake.case
func ToDelimited(s string, delimiter uint8) string {
	return ToScreamingDelimited(s, delimiter, "", false)
}

// ToScreamingDelimited converts a string to SCREAMING_DELIMITED_CASE
func ToScreamingDelimited(s string, delimiter uint8, ignore string, screaming bool) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}

	n := strings.Builder{}
	n.Grow(len(s) + 2)
	for i := 0; i < len(s); i++ {
		v := s[i]
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if vIsLow && screaming {
			v -= 'a' - 'A'
		} else if vIsCap && !screaming {
			v += 'a' - 'A'
		}

		if i+1 < len(s) {
			next := s[i+1]
			vIsNum := v >= '0' && v <= '9'
			nextIsCap := next >= 'A' && next <= 'Z'
			nextIsLow := next >= 'a' && next <= 'z'
			nextIsNum := next >= '0' && next <= '9'
			if (vIsCap && (nextIsLow || nextIsNum)) || (vIsLow && (nextIsCap || nextIsNum)) || (vIsNum && (nextIsCap || nextIsLow)) {
				prevIgnore := ignore != "" && i > 0 && strings.IndexByte(ignore, s[i-1]) >= 0
				if !prevIgnore {
					if vIsCap && nextIsLow {
						if prevIsCap := i > 0 && s[i-1] >= 'A' && s[i-1] <= 'Z'; prevIsCap {
							n.WriteByte(delimiter)
						}
					}
					n.WriteByte(v)
					if vIsLow || vIsNum || nextIsNum {
						n.WriteByte(delimiter)
					}
					continue
				}
			}
		}

		if v == ' ' || v == '_' || v == '-' || v == '.' {
			if ignore == "" || strings.IndexByte(ignore, v) < 0 {
				n.WriteByte(delimiter)
			} else {
				n.WriteByte(v)
			}
		} else {
			n.WriteByte(v)
		}
	}

	return n.String()
}
