package strings

import (
	"github.com/iancoleman/strcase"
)

var (
	// ToSnake converts a string to snake_case format.
	ToSnake = strcase.ToSnake
	// ToSnakeWithIgnore converts a string to snake_case format, ignoring specific characters.
	ToSnakeWithIgnore = strcase.ToSnakeWithIgnore
	// ToScreamingSnake converts a string to SCREAMING_SNAKE_CASE format.
	ToScreamingSnake = strcase.ToScreamingSnake
	// ToKebab converts a string to kebab-case format.
	ToKebab = strcase.ToKebab
	// ToScreamingKebab converts a string to SCREAMING-KEBAB-CASE format.
	ToScreamingKebab = strcase.ToScreamingKebab
	// ToDelimited converts a string to a custom delimiter-separated format.
	ToDelimited = strcase.ToDelimited
	// ToScreamingDelimited converts a string to SCREAMING custom delimiter-separated format.
	ToScreamingDelimited = strcase.ToScreamingDelimited
	// ToCamel converts a string to CamelCase format.
	ToCamel = strcase.ToCamel
	// ToLowerCamel converts a string to lowerCamelCase format.
	ToLowerCamel = strcase.ToLowerCamel
)
