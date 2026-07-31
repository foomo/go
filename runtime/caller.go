package runtime

import (
	"path"
	stdruntime "runtime"
	"strings"
)

// Caller returns caller information for the function skip frames up the call stack.
func Caller(skip int) (shortName, fullName, file string, line int, ok bool) { //nolint:nonamedreturns
	var pcs [1]uintptr
	if stdruntime.Callers(skip+2, pcs[:]) == 0 {
		return "unknown", "Unknown", "unknown", 0, false
	}

	fr, _ := stdruntime.CallersFrames(pcs[:]).Next()
	if fr.Function == "" {
		return "unknown", "Unknown", "unknown", 0, false
	}

	fullName, file, line = fr.Function, fr.File, fr.Line

	dirname, filename := path.Split(file)
	file = path.Join(path.Base(dirname), filename)
	// Split fullName by last slash to separate package path and the rest
	lastSlash := strings.LastIndex(fullName, "/")
	if lastSlash != -1 {
		return fullName[lastSlash+1:], fullName, file, line, true
	}

	return fullName, fullName, file, line, true
}
