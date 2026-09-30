package runtime

import (
	"runtime"
	"strings"
)

// CallerFunc returns the unqualified name of the function skip frames up the
// call stack, stripped of its package path and any receiver type. It returns
// "unknown", false if the stack does not have that many frames.
func CallerFunc(skip int) (string, bool) {
	pc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		return "unknown", false
	}

	fullName := runtime.FuncForPC(pc).Name()

	// Split fullName by last slash to separate package path and the rest
	lastSlash := strings.LastIndex(fullName, "/")
	if lastSlash != -1 {
		fullName = fullName[lastSlash+1:]
	}

	lastDot := strings.LastIndex(fullName, ".")
	if lastDot != -1 {
		fullName = fullName[lastDot+1:]
	}

	return fullName, true
}
