package os

import (
	"fmt"
	"os"
)

// Exists return true if given env var is defined.
func Exists(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}

// MustExists panics if the given env var does not exists.
func MustExists(key string) {
	if _, ok := os.LookupEnv(key); !ok {
		panic(fmt.Sprintf("required environment variable '%s' is not defined", key))
	}
}
