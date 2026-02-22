package testing_test

import (
	"testing"

	testingx "github.com/foomo/go/testing"
	tagx "github.com/foomo/go/testing/tag"
	"github.com/stretchr/testify/assert"
)

func TestSkipTags(t *testing.T) {
	tests := []struct {
		name string
		env  string
		tags []tagx.Tag
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
			tags: []tagx.Tag{},
			want: true,
		},
		{
			name: "skip for empty and short",
			env:  "short",
			tags: []tagx.Tag{},
			want: true,
		},
		{
			name: "skip for integration,security and integration",
			env:  "short",
			tags: []tagx.Tag{tagx.Integration, tagx.Security},
			want: true,
		},
		{
			name: "skip for integration,security and short",
			env:  "short",
			tags: []tagx.Tag{tagx.Integration, tagx.Security},
			want: true,
		},
		{
			name: "don't skip for integration and integration",
			env:  "integration",
			tags: []tagx.Tag{tagx.Integration},
			want: false,
		},
		{
			name: "don't skip for integration,security and integration",
			env:  "integration",
			tags: []tagx.Tag{tagx.Integration, tagx.Security},
			want: false,
		},
		{
			name: "don't skip for integration,security and security",
			env:  "security",
			tags: []tagx.Tag{tagx.Integration, tagx.Security},
			want: false,
		},
		{
			name: "skip for integration,security and -integration,security",
			env:  "-integration,security",
			tags: []tagx.Tag{tagx.Integration, tagx.Security},
			want: true,
		},
		{
			name: "skip for integration,security and integration,-security",
			env:  "integration,-security",
			tags: []tagx.Tag{tagx.Integration, tagx.Security},
			want: true,
		},
		{
			name: "don't skip for integration,security and -short",
			env:  "-short",
			tags: []tagx.Tag{tagx.Integration, tagx.Security},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.env != "" {
				t.Setenv(testingx.EnvTestTags, tt.env)
			}

			assert.Equalf(t, tt.want, testingx.SkipTags(tt.tags...), "skipTags(%v)", tt.tags)
		})
	}
}
