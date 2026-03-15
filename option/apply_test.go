package option_test

import (
	"errors"
	"fmt"

	"github.com/foomo/go/option"
)

type Server struct {
	Name string
}

func WithName(p string) option.Option[*Server] {
	return func(s *Server) {
		s.Name = p
	}
}

func ExampleApply() {
	s := &Server{}

	option.Apply(
		s,
		WithName("localhost"),
	)

	fmt.Printf("%s\n", s.Name)

	// Output:
	// localhost
}

func WithNameE(p string) option.OptionE[*Server] {
	return func(s *Server) error {
		if len(p) == 0 {
			return errors.New("invalid name")
		}

		s.Name = p

		return nil
	}
}

func ExampleApplyE() {
	s := &Server{}

	err := option.ApplyE(
		s,
		WithNameE("localhost"),
	)

	fmt.Println(err)
	fmt.Println(s.Name)

	// Output:
	// <nil>
	// localhost
}
