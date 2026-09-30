package types

type Causer interface {
	Cause() error
}

type CauserFunc func() error

func (f CauserFunc) Cause() error {
	return f()
}
