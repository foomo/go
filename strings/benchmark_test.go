package strings_test

import (
	"testing"

	"github.com/foomo/go/strings"
)

func BenchmarkToCamel(b *testing.B) {
	for b.Loop() {
		strings.ToCamel("some_snake_case_string")
	}
}

func BenchmarkToLowerCamel(b *testing.B) {
	for b.Loop() {
		strings.ToLowerCamel("some_snake_case_string")
	}
}

func BenchmarkToSnake(b *testing.B) {
	for b.Loop() {
		strings.ToSnake("SomeCamelCaseString")
	}
}

func BenchmarkToKebab(b *testing.B) {
	for b.Loop() {
		strings.ToKebab("SomeCamelCaseString")
	}
}

func BenchmarkFirstToUpper(b *testing.B) {
	for b.Loop() {
		strings.FirstToUpper("hello")
	}
}

func BenchmarkFirstToLower(b *testing.B) {
	for b.Loop() {
		strings.FirstToLower("Hello")
	}
}
