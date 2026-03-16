package option

type (
	// Builder is a generic option builder that collects functional options.
	Builder[T any] struct {
		Opts []Option[T]
	}
	// BuilderE is a generic option builder that collects functional error-returning options.
	BuilderE[T any] struct {
		Opts []OptionE[T]
	}
)

// List returns the collected options.
func (b *Builder[T]) List() []Option[T] {
	return b.Opts
}

// List returns the collected options.
func (b *BuilderE[T]) List() []OptionE[T] {
	return b.Opts
}
