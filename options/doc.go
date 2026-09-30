// Package options implements the generic functional options pattern:
// [Option] and [OptionE] are functions that configure a value of type T,
// [Apply] and [ApplyE] run a slice of them against a value, and [Builder]
// and [BuilderE] collect options for reuse across multiple Apply calls.
package options
