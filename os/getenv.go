package os

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	slicesx "github.com/foomo/go/slices"
)

var (
	// SliceSeparator is the delimiter used to separate elements in a slice when parsing comma-separated values.
	SliceSeparator = ","
	// MapSeparator is the delimiter used to separate key-value pairs within a string representing a map.
	MapSeparator = ","
	// MapKVSeparator is the delimiter used to separate key and value in a key-value pair.
	MapKVSeparator = ":"
)

// ---------------------------------- env -------------------------------------

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

// -------------------------------- scalars -----------------------------------

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
	return getenvParse(key, def, parseBool)
}

// GetenvInt wraps os.Getenv and returns an int or the given default if empty or not defined.
func GetenvInt(key string, def int) (int, error) {
	return getenvParse(key, def, parseInt)
}

// GetenvInt8 wraps os.Getenv and returns an int8 or the given default if empty or not defined.
func GetenvInt8(key string, def int8) (int8, error) {
	return getenvParse(key, def, parseInt8)
}

// GetenvInt16 wraps os.Getenv and returns an int16 or the given default if empty or not defined.
func GetenvInt16(key string, def int16) (int16, error) {
	return getenvParse(key, def, parseInt16)
}

// GetenvInt32 wraps os.Getenv and returns an int32 or the given default if empty or not defined.
func GetenvInt32(key string, def int32) (int32, error) {
	return getenvParse(key, def, parseInt32)
}

// GetenvInt64 wraps os.Getenv and returns an int64 or the given default if empty or not defined.
func GetenvInt64(key string, def int64) (int64, error) {
	return getenvParse(key, def, parseInt64)
}

// GetenvUint wraps os.Getenv and returns a uint or the given default if empty or not defined.
func GetenvUint(key string, def uint) (uint, error) {
	return getenvParse(key, def, parseUint)
}

// GetenvUint8 wraps os.Getenv and returns a uint8 or the given default if empty or not defined.
func GetenvUint8(key string, def uint8) (uint8, error) {
	return getenvParse(key, def, parseUint8)
}

// GetenvUint16 wraps os.Getenv and returns a uint16 or the given default if empty or not defined.
func GetenvUint16(key string, def uint16) (uint16, error) {
	return getenvParse(key, def, parseUint16)
}

// GetenvUint32 wraps os.Getenv and returns a uint32 or the given default if empty or not defined.
func GetenvUint32(key string, def uint32) (uint32, error) {
	return getenvParse(key, def, parseUint32)
}

// GetenvUint64 wraps os.Getenv and returns a uint64 or the given default if empty or not defined.
func GetenvUint64(key string, def uint64) (uint64, error) {
	return getenvParse(key, def, parseUint64)
}

// GetenvFloat32 wraps os.Getenv and returns a float32 or the given default if empty or not defined.
func GetenvFloat32(key string, def float32) (float32, error) {
	return getenvParse(key, def, parseFloat32)
}

// GetenvFloat64 wraps os.Getenv and returns a float64 or the given default if empty or not defined.
func GetenvFloat64(key string, def float64) (float64, error) {
	return getenvParse(key, def, parseFloat64)
}

// GetenvDuration wraps os.Getenv and returns a time.Duration or the given default if empty or not defined.
func GetenvDuration(key string, def time.Duration) (time.Duration, error) {
	return getenvParse(key, def, parseDuration)
}

// ------------------------------ must scalars --------------------------------

// MustGetenv returns the value of the environment variable or panics if it is not defined.
func MustGetenv(key string) string {
	str, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("required environment variable '%s' is not defined", key))
	}

	return str
}

// MustGetenvBool returns the bool value of the environment variable or panics if it is not defined or cannot be parsed.
func MustGetenvBool(key string) bool {
	return mustGetenvParse(key, parseBool)
}

// MustGetenvInt returns the int value of the environment variable or panics if it is not defined or cannot be parsed.
func MustGetenvInt(key string) int {
	return mustGetenvParse(key, parseInt)
}

// MustGetenvInt8 returns the int8 value of the environment variable or panics if it is not defined or cannot be parsed.
func MustGetenvInt8(key string) int8 {
	return mustGetenvParse(key, parseInt8)
}

// MustGetenvInt16 returns the int16 value of the environment variable or panics if it is not defined or cannot be parsed.
func MustGetenvInt16(key string) int16 {
	return mustGetenvParse(key, parseInt16)
}

// MustGetenvInt32 returns the int32 value of the environment variable or panics if it is not defined or cannot be parsed.
func MustGetenvInt32(key string) int32 {
	return mustGetenvParse(key, parseInt32)
}

// MustGetenvInt64 returns the int64 value of the environment variable or panics if it is not defined or cannot be parsed.
func MustGetenvInt64(key string) int64 {
	return mustGetenvParse(key, parseInt64)
}

// MustGetenvUint returns the uint value of the environment variable or panics if it is not defined or cannot be parsed.
func MustGetenvUint(key string) uint {
	return mustGetenvParse(key, parseUint)
}

// MustGetenvUint8 returns the uint8 value of the environment variable or panics if it is not defined or cannot be parsed.
func MustGetenvUint8(key string) uint8 {
	return mustGetenvParse(key, parseUint8)
}

// MustGetenvUint16 returns the uint16 value of the environment variable or panics if it is not defined or cannot be parsed.
func MustGetenvUint16(key string) uint16 {
	return mustGetenvParse(key, parseUint16)
}

// MustGetenvUint32 returns the uint32 value of the environment variable or panics if it is not defined or cannot be parsed.
func MustGetenvUint32(key string) uint32 {
	return mustGetenvParse(key, parseUint32)
}

// MustGetenvUint64 returns the uint64 value of the environment variable or panics if it is not defined or cannot be parsed.
func MustGetenvUint64(key string) uint64 {
	return mustGetenvParse(key, parseUint64)
}

// MustGetenvFloat32 returns the float32 value of the environment variable or panics if it is not defined or cannot be parsed.
func MustGetenvFloat32(key string) float32 {
	return mustGetenvParse(key, parseFloat32)
}

// MustGetenvFloat64 returns the float64 value of the environment variable or panics if it is not defined or cannot be parsed.
func MustGetenvFloat64(key string) float64 {
	return mustGetenvParse(key, parseFloat64)
}

// MustGetenvDuration returns the time.Duration value of the environment variable or panics if it is not defined or cannot be parsed.
func MustGetenvDuration(key string) time.Duration {
	return mustGetenvParse(key, parseDuration)
}

// --------------------------------- slices -----------------------------------

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

// GetenvBoolSlice wraps os.Getenv and returns a bool slice or the given default if empty or not defined.
func GetenvBoolSlice(key string, def []bool) ([]bool, error) {
	return getenvSlice(key, def, parseBool)
}

// GetenvIntSlice wraps os.Getenv and returns an int slice or the given default if empty or not defined.
func GetenvIntSlice(key string, def []int) ([]int, error) {
	return getenvSlice(key, def, parseInt)
}

// GetenvInt8Slice wraps os.Getenv and returns an int8 slice or the given default if empty or not defined.
func GetenvInt8Slice(key string, def []int8) ([]int8, error) {
	return getenvSlice(key, def, parseInt8)
}

// GetenvInt16Slice wraps os.Getenv and returns an int16 slice or the given default if empty or not defined.
func GetenvInt16Slice(key string, def []int16) ([]int16, error) {
	return getenvSlice(key, def, parseInt16)
}

// GetenvInt32Slice wraps os.Getenv and returns an int32 slice or the given default if empty or not defined.
func GetenvInt32Slice(key string, def []int32) ([]int32, error) {
	return getenvSlice(key, def, parseInt32)
}

// GetenvInt64Slice wraps os.Getenv and returns an int64 slice or the given default if empty or not defined.
func GetenvInt64Slice(key string, def []int64) ([]int64, error) {
	return getenvSlice(key, def, parseInt64)
}

// GetenvUintSlice wraps os.Getenv and returns a uint slice or the given default if empty or not defined.
func GetenvUintSlice(key string, def []uint) ([]uint, error) {
	return getenvSlice(key, def, parseUint)
}

// GetenvUint8Slice wraps os.Getenv and returns a uint8 slice or the given default if empty or not defined.
func GetenvUint8Slice(key string, def []uint8) ([]uint8, error) {
	return getenvSlice(key, def, parseUint8)
}

// GetenvUint16Slice wraps os.Getenv and returns a uint16 slice or the given default if empty or not defined.
func GetenvUint16Slice(key string, def []uint16) ([]uint16, error) {
	return getenvSlice(key, def, parseUint16)
}

// GetenvUint32Slice wraps os.Getenv and returns a uint32 slice or the given default if empty or not defined.
func GetenvUint32Slice(key string, def []uint32) ([]uint32, error) {
	return getenvSlice(key, def, parseUint32)
}

// GetenvUint64Slice wraps os.Getenv and returns a uint64 slice or the given default if empty or not defined.
func GetenvUint64Slice(key string, def []uint64) ([]uint64, error) {
	return getenvSlice(key, def, parseUint64)
}

// GetenvFloat32Slice wraps os.Getenv and returns a float32 slice or the given default if empty or not defined.
func GetenvFloat32Slice(key string, def []float32) ([]float32, error) {
	return getenvSlice(key, def, parseFloat32)
}

// GetenvFloat64Slice wraps os.Getenv and returns a float64 slice or the given default if empty or not defined.
func GetenvFloat64Slice(key string, def []float64) ([]float64, error) {
	return getenvSlice(key, def, parseFloat64)
}

// GetenvDurationSlice wraps os.Getenv and returns a time.Duration slice or the given default if empty or not defined.
func GetenvDurationSlice(key string, def []time.Duration) ([]time.Duration, error) {
	return getenvSlice(key, def, parseDuration)
}

// ---------------------------------- maps ------------------------------------

// GetenvStringMap wraps os.Getenv and returns a map[string]string or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "a:1,b:2").
func GetenvStringMap(key string, def map[string]string) (map[string]string, error) {
	return getenvMap(key, def, parseString)
}

// Deprecated: Use GetenvStringMap instead.
func GetenvStringMapString(key string, def map[string]string) (map[string]string, error) {
	return GetenvStringMap(key, def)
}

// GetenvBoolMap wraps os.Getenv and returns a map[string]bool or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "debug:true,verbose:false").
func GetenvBoolMap(key string, def map[string]bool) (map[string]bool, error) {
	return getenvMap(key, def, parseBool)
}

// GetenvIntMap wraps os.Getenv and returns a map[string]int or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "a:1,b:2").
func GetenvIntMap(key string, def map[string]int) (map[string]int, error) {
	return getenvMap(key, def, parseInt)
}

// GetenvInt8Map wraps os.Getenv and returns a map[string]int8 or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "a:1,b:2").
func GetenvInt8Map(key string, def map[string]int8) (map[string]int8, error) {
	return getenvMap(key, def, parseInt8)
}

// GetenvInt16Map wraps os.Getenv and returns a map[string]int16 or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "a:1,b:2").
func GetenvInt16Map(key string, def map[string]int16) (map[string]int16, error) {
	return getenvMap(key, def, parseInt16)
}

// GetenvInt32Map wraps os.Getenv and returns a map[string]int32 or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "a:1,b:2").
func GetenvInt32Map(key string, def map[string]int32) (map[string]int32, error) {
	return getenvMap(key, def, parseInt32)
}

// GetenvInt64Map wraps os.Getenv and returns a map[string]int64 or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "a:1,b:2").
func GetenvInt64Map(key string, def map[string]int64) (map[string]int64, error) {
	return getenvMap(key, def, parseInt64)
}

// GetenvUintMap wraps os.Getenv and returns a map[string]uint or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "a:1,b:2").
func GetenvUintMap(key string, def map[string]uint) (map[string]uint, error) {
	return getenvMap(key, def, parseUint)
}

// GetenvUint8Map wraps os.Getenv and returns a map[string]uint8 or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "a:1,b:2").
func GetenvUint8Map(key string, def map[string]uint8) (map[string]uint8, error) {
	return getenvMap(key, def, parseUint8)
}

// GetenvUint16Map wraps os.Getenv and returns a map[string]uint16 or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "a:1,b:2").
func GetenvUint16Map(key string, def map[string]uint16) (map[string]uint16, error) {
	return getenvMap(key, def, parseUint16)
}

// GetenvUint32Map wraps os.Getenv and returns a map[string]uint32 or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "a:1,b:2").
func GetenvUint32Map(key string, def map[string]uint32) (map[string]uint32, error) {
	return getenvMap(key, def, parseUint32)
}

// GetenvUint64Map wraps os.Getenv and returns a map[string]uint64 or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "a:1,b:2").
func GetenvUint64Map(key string, def map[string]uint64) (map[string]uint64, error) {
	return getenvMap(key, def, parseUint64)
}

// GetenvFloat32Map wraps os.Getenv and returns a map[string]float32 or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "a:1.5,b:2.5").
func GetenvFloat32Map(key string, def map[string]float32) (map[string]float32, error) {
	return getenvMap(key, def, parseFloat32)
}

// GetenvFloat64Map wraps os.Getenv and returns a map[string]float64 or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "a:1.5,b:2.5").
func GetenvFloat64Map(key string, def map[string]float64) (map[string]float64, error) {
	return getenvMap(key, def, parseFloat64)
}

// GetenvDurationMap wraps os.Getenv and returns a map[string]time.Duration or the given default if empty or not defined.
// The expected format is comma-separated key:value pairs (e.g. "timeout:5s,interval:100ms").
func GetenvDurationMap(key string, def map[string]time.Duration) (map[string]time.Duration, error) {
	return getenvMap(key, def, parseDuration)
}

// ---------------------------------- helpers ----------------------------------

// getenv returns the raw env value and whether it's empty/unset.
func getenv(key string) (string, bool) {
	str := os.Getenv(key)
	return str, str != ""
}

// mustGetenvParse is the generic must-have scalar helper that panics on missing or invalid values.
func mustGetenvParse[T any](key string, parse func(string) (T, error)) T {
	str, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("required environment variable '%s' is not defined", key))
	}

	v, err := parse(str)
	if err != nil {
		panic(fmt.Sprintf("environment variable '%s': %v", key, err))
	}

	return v
}

// getenvParse is the generic scalar helper.
func getenvParse[T any](key string, def T, parse func(string) (T, error)) (T, error) {
	str, ok := getenv(key)
	if !ok {
		return def, nil
	}

	return parse(str)
}

// getenvSlice is the generic comma-separated slice helper.
func getenvSlice[T any](key string, def []T, parse func(string) (T, error)) ([]T, error) {
	str, ok := getenv(key)
	if !ok {
		return def, nil
	}

	parts := strings.Split(str, SliceSeparator)

	return slicesx.MapE(parts, func(p string) (T, error) {
		return parse(strings.TrimSpace(p))
	})
}

// getenvMap is the generic comma-separated key:value map helper.
func getenvMap[T any](key string, def map[string]T, parse func(string) (T, error)) (map[string]T, error) {
	str, ok := getenv(key)
	if !ok {
		return def, nil
	}

	parts := strings.Split(str, MapSeparator)

	result := make(map[string]T, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)

		kv := strings.SplitN(part, MapKVSeparator, 2)
		if len(kv) != 2 {
			return nil, fmt.Errorf("invalid key:value pair: %q", part)
		}

		v, err := parse(strings.TrimSpace(kv[1]))
		if err != nil {
			return nil, err
		}

		result[strings.TrimSpace(kv[0])] = v
	}

	return result, nil
}

// --------------------------------- parsers ----------------------------------

func parseString(s string) (string, error) { return s, nil }
func parseBool(s string) (bool, error)     { return strconv.ParseBool(s) }
func parseInt(s string) (int, error)       { return strconv.Atoi(s) }
func parseInt8(s string) (int8, error)     { v, err := strconv.ParseInt(s, 0, 8); return int8(v), err }
func parseInt16(s string) (int16, error)   { v, err := strconv.ParseInt(s, 0, 16); return int16(v), err }
func parseInt32(s string) (int32, error)   { v, err := strconv.ParseInt(s, 0, 32); return int32(v), err }
func parseInt64(s string) (int64, error)   { v, err := strconv.ParseInt(s, 0, 64); return v, err }
func parseUint(s string) (uint, error) {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	return uint(v), err
}
func parseUint8(s string) (uint8, error) { v, err := strconv.ParseUint(s, 0, 8); return uint8(v), err }
func parseUint16(s string) (uint16, error) {
	v, err := strconv.ParseUint(s, 0, 16)
	return uint16(v), err
}
func parseUint32(s string) (uint32, error) {
	v, err := strconv.ParseUint(s, 0, 32)
	return uint32(v), err
}
func parseUint64(s string) (uint64, error) {
	v, err := strconv.ParseUint(s, 0, 64)
	return v, err
}
func parseFloat32(s string) (float32, error) {
	v, err := strconv.ParseFloat(s, 32)
	return float32(v), err
}
func parseFloat64(s string) (float64, error)        { return strconv.ParseFloat(s, 64) }
func parseDuration(s string) (time.Duration, error) { return time.ParseDuration(s) }
