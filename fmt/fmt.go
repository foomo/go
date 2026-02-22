package fmt

import (
	"bytes"
	"text/template"
)

// Tprintf formats a string using a template
func Tprintf(format string, pairs ...string) string {
	t, err := template.New("").Delims("%{", "}").Parse(format)
	if err != nil {
		return "parse error: " + err.Error()
	}

	buf := &bytes.Buffer{}
	d := argsToAttr(pairs)
	if err := t.Execute(buf, d); err != nil {
		return "render error: " + err.Error()
	}

	return buf.String()
}

// argsToAttr converts a list of key-value pairs to a map
func argsToAttr(args []string) map[string]string {
	data := map[string]string{}
	for i := 0; i < len(args); i = i + 2 {
		data[args[i]] = args[i+1]
	}

	return data
}
