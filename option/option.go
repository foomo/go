package option

type (
	// Option is a functional option.
	Option[T any] func(T)
	// OptionE is a functional error-returning option.
	OptionE[T any] func(T) error
)
