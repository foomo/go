package testing //nolint:testpackage

import (
	"os"
	"testing"

	"github.com/foomo/go/testing/tag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_skipTags(t *testing.T) {
	tests := []struct {
		name string
		env  string
		tags []tag.Tag
		want bool
	}{
		{
			name: "skip for nil",
			env:  "",
			tags: nil,
			want: true,
		},
		{
			name: "skip for nil and short",
			env:  "short",
			tags: nil,
			want: true,
		},
		{
			name: "skip for empty and short",
			env:  "short",
			tags: []tag.Tag{},
			want: true,
		},
		{
			name: "skip for empty and short",
			env:  "short",
			tags: []tag.Tag{},
			want: true,
		},
		{
			name: "skip for integration,security and integration",
			env:  "short",
			tags: []tag.Tag{tag.Integration, tag.Security},
			want: true,
		},
		{
			name: "skip for integration,security and short",
			env:  "short",
			tags: []tag.Tag{tag.Integration, tag.Security},
			want: true,
		},
		{
			name: "don't skip for integration and integration",
			env:  "integration",
			tags: []tag.Tag{tag.Integration},
			want: false,
		},
		{
			name: "don't skip for integration,security and integration",
			env:  "integration",
			tags: []tag.Tag{tag.Integration, tag.Security},
			want: false,
		},
		{
			name: "don't skip for integration,security and security",
			env:  "security",
			tags: []tag.Tag{tag.Integration, tag.Security},
			want: false,
		},
		{
			name: "skip for integration,security and -integration,security",
			env:  "-integration,security",
			tags: []tag.Tag{tag.Integration, tag.Security},
			want: true,
		},
		{
			name: "skip for integration,security and integration,-security",
			env:  "integration,-security",
			tags: []tag.Tag{tag.Integration, tag.Security},
			want: true,
		},
		{
			name: "don't skip for integration,security and -short",
			env:  "-short",
			tags: []tag.Tag{tag.Integration, tag.Security},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.env == "" {
				require.NoError(t, os.Unsetenv(envTestTags))
			} else {
				require.NoError(t, os.Setenv(envTestTags, tt.env))
			}
			assert.Equalf(t, tt.want, skipTags(tt.tags), "skipTags(%v)", tt.tags)
		})
	}
}
