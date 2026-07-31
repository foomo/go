package options_test

import (
	"errors"
	"testing"

	"github.com/foomo/go/options"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApplyE_Error(t *testing.T) {
	t.Parallel()

	s := &Server{}
	err := options.ApplyE(s, WithNameE("")) // empty name returns an error
	require.Error(t, err)
	assert.Contains(t, err.Error(), "apply option")
}

func TestBuildE_Error(t *testing.T) {
	t.Parallel()

	b := MyBuilderE()
	b.Opts = append(b.Opts, func(*MyOptions) error {
		return errors.New("boom")
	})

	err := options.BuildE(&MyOptions{}, b)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "build option")
}
