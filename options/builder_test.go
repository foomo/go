package options_test

import (
	"fmt"

	"github.com/foomo/go/options"
)

type MyOptions struct {
	Name string
}

type MyOptionsBuilder struct {
	options.Builder[*MyOptions]
}

func NewMyOptionsBuilder() *MyOptionsBuilder {
	return &MyOptionsBuilder{
		options.Builder[*MyOptions]{},
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
	options.Build(&o, b)

	fmt.Println(o)

	// Output: {example}
}

type MyOptionsBuilderE struct {
	options.BuilderE[*MyOptions]
}

func MyBuilderE() *MyOptionsBuilderE {
	return &MyOptionsBuilderE{
		options.BuilderE[*MyOptions]{},
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
	if err := options.BuildE(&o, b); err != nil {
		panic(err)
	}

	fmt.Println(o)

	// Output: {example}
}
