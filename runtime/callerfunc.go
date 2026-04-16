package runtime

import (
	"runtime"
	"strings"
)

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
