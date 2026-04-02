package runtime

import (
	"path"
	"runtime"
	"strconv"
	"strings"
)

// StackTrace captures and formats a stack trace up to the specified number of frames, skipping the given number of initial frames.
func StackTrace(size, skip int) string {
	skip = max(0, min(20, skip))
	size = max(1, min(20, size)+skip)

	pcs := make([]uintptr, size)

	n := runtime.Callers(skip+2, pcs)
	if n == 0 {
		return "stack trace out of bounds"
	}

	pcs = pcs[:n]

	var ret strings.Builder

	frames := runtime.CallersFrames(pcs)
	for {
		frame, more := frames.Next()

		ret.WriteString(frame.Function)
		ret.WriteString("\n  ")

		dir, file := path.Split(frame.File)
		ret.WriteString(path.Join(path.Base(dir), file))
		ret.WriteString(":")
		ret.WriteString(strconv.Itoa(frame.Line))

		if !more || len(ret.String()) == size {
			break
		}

		ret.WriteString("\n")
	}

	return ret.String()
}
