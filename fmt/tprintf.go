package fmt

import "strings"

// Tprintf formats a string by replacing %{.key} placeholders with values from the given key-value pairs.
func Tprintf(format string, pairs ...string) string {
	m := argsToAttr(pairs)

	var b strings.Builder
	b.Grow(len(format))

	s := format
	for {
		i := strings.Index(s, "%{.")
		if i < 0 {
			b.WriteString(s)

			break
		}

		b.WriteString(s[:i])

		rest := s[i+3:] // after "%{."

		key, after, ok := strings.Cut(rest, "}")
		if !ok {
			// No closing brace — write the rest literally
			b.WriteString(s[i:])

			break
		}

		b.WriteString(m[key])

		s = after
	}

	return b.String()
}

// argsToAttr converts a list of key-value pairs to a map.
func argsToAttr(args []string) map[string]string {
	data := make(map[string]string, len(args)/2)
	for i := 0; i+1 < len(args); i = i + 2 {
		data[args[i]] = args[i+1]
	}

	return data
}
