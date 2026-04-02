package option_test

import (
	"fmt"

	"github.com/foomo/go/option"
)

type MyOptions struct {
	Name string
}

type MyOptionsBuilder struct {
	option.Builder[*MyOptions]
}

func NewMyOptionsBuilder() *MyOptionsBuilder {
	return &MyOptionsBuilder{
		option.Builder[*MyOptions]{},
	}
}

func (b *MyOptionsBuilder) WithName(name string) *MyOptionsBuilder {
	b.Opts = append(b.Opts, func(o *MyOptions) {
		o.Name = name
	})

	return b
}

func ExampleBuilder() {
	b := NewMyOptionsBuilder()
	b.WithName("example")

	o := MyOptions{}
	option.Build(&o, b)

	fmt.Println(o)

	// Output: {example}
}

type MyOptionsBuilderE struct {
	option.BuilderE[*MyOptions]
}

func MyBuilderE() *MyOptionsBuilderE {
	return &MyOptionsBuilderE{
		option.BuilderE[*MyOptions]{},
	}
}

func (b *MyOptionsBuilderE) WithName(name string) *MyOptionsBuilderE {
	b.Opts = append(b.Opts, func(o *MyOptions) error {
		o.Name = name
		return nil
	})

	return b
}

func ExampleBuilderE() {
	b := MyBuilderE()
	b.WithName("example")

	o := MyOptions{}
	if err := option.BuildE(&o, b); err != nil {
		panic(err)
	}

	fmt.Println(o)

	// Output: {example}
}
