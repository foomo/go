package slices_test

import (
	"errors"
	"testing"

	"github.com/foomo/go/slices"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilterE_Error(t *testing.T) {
	t.Parallel()

	errBoom := errors.New("boom")
	out, err := slices.FilterE([]int{1, 2, 3}, func(n int) (bool, error) {
		if n == 2 {
			return false, errBoom
		}

		return true, nil
	})
	require.ErrorIs(t, err, errBoom)
	assert.Nil(t, out)
}

func TestMapE_Error(t *testing.T) {
	t.Parallel()

	errBoom := errors.New("boom")
	out, err := slices.MapE([]int{1, 2, 3}, func(n int) (int, error) {
		if n == 2 {
			return 0, errBoom
		}

		return n, nil
	})
	require.ErrorIs(t, err, errBoom)
	assert.Nil(t, out)
}
