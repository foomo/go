package types

// Namer is implemented by types that expose an identifying name.
type Namer interface {
	Name() string
}
