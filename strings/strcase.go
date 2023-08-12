package strings

import (
	"github.com/iancoleman/strcase"
)

var (
	ToSnake              = strcase.ToSnake
	ToSnakeWithIgnore    = strcase.ToSnakeWithIgnore
	ToScreamingSnake     = strcase.ToScreamingSnake
	ToKebab              = strcase.ToKebab
	ToScreamingKebab     = strcase.ToScreamingKebab
	ToDelimited          = strcase.ToDelimited
	ToScreamingDelimited = strcase.ToScreamingDelimited
	ToCamel              = strcase.ToCamel
	ToLowerCamel         = strcase.ToLowerCamel
)
