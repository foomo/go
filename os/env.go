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

// GetenvInt32 wraps os.Getenv and returns an int32 or the given default if empty or not defined.
func GetenvInt32(key string, def int32) (int32, error) {
	str := os.Getenv(key)
	if str == "" {
		return def, nil
	}

	value, err := strconv.ParseInt(str, 0, 32)
	if err != nil {
		return 0, err
	}

	return int32(value), nil
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

// GetenvFloat32 wraps os.Getenv and returns a float32 or the given default if empty or not defined.
func GetenvFloat32(key string, def float32) (float32, error) {
	str := os.Getenv(key)
	if str == "" {
		return def, nil
	}

	value, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}

	return float32(value), nil
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

// GetenvStringMapString wraps os.Getenv and returns a map[string]string or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "a:1,b:2").
func GetenvStringMapString(key string, def map[string]string) (map[string]string, error) {
	str := os.Getenv(key)
	if str == "" {
		return def, nil
	}

	parts := strings.Split(str, ",")
	result := make(map[string]string, len(parts))

	for _, part := range parts {
		part = strings.TrimSpace(part)

		kv := strings.SplitN(part, ":", 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid key:value pair: %q", part)
		}

		result[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
	}

	return result, nil
}
