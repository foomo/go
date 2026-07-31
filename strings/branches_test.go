package strings_test

import (
	"testing"

	"github.com/foomo/go/strings"
	"github.com/stretchr/testify/assert"
)

func TestIsBranches(t *testing.T) {
	t.Parallel()

	assert.True(t, strings.IsBlank(""))

	assert.True(t, strings.IsAnyEmpty())
	assert.False(t, strings.IsAnyEmpty("a", "b"))

	assert.True(t, strings.IsAnyBlank())
	assert.False(t, strings.IsAnyBlank("a", "b"))

	assert.False(t, strings.IsAlpha(""))
	assert.False(t, strings.IsAlpha("ab1"))

	assert.False(t, strings.IsAlphanumeric(""))
	assert.False(t, strings.IsAlphanumeric("ab!"))

	assert.False(t, strings.IsNumeric(""))
	assert.False(t, strings.IsNumeric("12a"))

	assert.False(t, strings.IsNumerical("abc"))
}

func TestHasAnyPrefixBranches(t *testing.T) {
	t.Parallel()

	assert.False(t, strings.HasAnyPrefix(""))
	assert.False(t, strings.HasAnyPrefix("foo"))
	assert.True(t, strings.HasAnyPrefix("foo", "foo"))    // exact match in list
	assert.True(t, strings.HasAnyPrefix("foobar", "foo")) // prefix match
	assert.False(t, strings.HasAnyPrefix("foobar", "baz"))
}

func TestHasAnySuffixBranches(t *testing.T) {
	t.Parallel()

	assert.False(t, strings.HasAnySuffix(""))
	assert.False(t, strings.HasAnySuffix("foo"))
	assert.True(t, strings.HasAnySuffix("bar", "bar"))    // exact match in list
	assert.True(t, strings.HasAnySuffix("foobar", "bar")) // suffix match
	assert.False(t, strings.HasAnySuffix("foobar", "baz"))
}

func TestToSnakeWithIgnore(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "hello_world.example", strings.ToSnakeWithIgnore("HelloWorld.Example", "."))
	assert.Empty(t, strings.ToSnakeWithIgnore("", "."))
}

func TestDelimitedAndCamelEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty(t, strings.ToScreamingDelimited("", '_', "", true))
	assert.Empty(t, strings.ToCamel(""))
	assert.Empty(t, strings.ToLowerCamel(""))
}

func TestRemoveAllBranches(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "unchanged", strings.RemoveAll("unchanged"))
	assert.Equal(t, "ac", strings.RemoveAll("abc", "b"))
}
