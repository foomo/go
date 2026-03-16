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

func MyBuilder() *MyOptionsBuilder {
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
	b := MyBuilder()
	b.WithName("example")

	o := MyOptions{}
	option.Apply(&o, b.List()...)

	fmt.Println(o)

	// Output: {example}
}
