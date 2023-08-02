package os

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// HasEnv return true if given env var is defined.
func HasEnv(key string) bool {
	_, ok := os.LookupEnv(key)
	return ok
}

// MustHasEnv panics if the given env var does not exists.
func MustHasEnv(key string) {
	if _, ok := os.LookupEnv(key); !ok {
		panic(fmt.Sprintf("required environment variable '%s' is not defined", key))
	}
}

// Getenv wraps os.Getenv and returns the given default if empty or not defined.
func Getenv(key string, def string) string {
	str := os.Getenv(key)
	if str == "" {
		return def
	}
	return str
}

// GetenvBool wraps os.Getenv and returns a bool or the given default if empty or not defined.
func GetenvBool(key string, def bool) (bool, error) {
	str := os.Getenv(key)
	if str == "" {
		return def, nil
	}
	value, err := strconv.ParseBool(str)
	if err != nil {
		return false, err
	}
	return value, nil
}

// GetenvInt64 wraps os.Getenv and returns an int or the given default if empty or not defined.
func GetenvInt64(key string, def int64) (int64, error) {
	str := os.Getenv(key)
	if str == "" {
		return def, nil
	}
	value, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// GetenvFloat64 wraps os.Getenv and returns an int or the given default if empty or not defined.
func GetenvFloat64(key string, def float64) (float64, error) {
	str := os.Getenv(key)
	if str == "" {
		return def, nil
	}
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// GetenvStringSlice wraps os.Getenv and returns a string slice or the given default if empty or not defined.
func GetenvStringSlice(key string, def []string) []string {
	str := os.Getenv(key)
	if str == "" {
		return def
	}
	value := strings.Split(str, ",")
	for i, s := range value {
		value[i] = strings.TrimSpace(s)
	}
	return value
}
