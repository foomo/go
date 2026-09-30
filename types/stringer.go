package types

// Stringer is implemented by types that can format themselves as a string.
// It mirrors [fmt.Stringer] and exists so packages that depend on the types
// package need not import fmt.
type Stringer interface {
	String() string
}
