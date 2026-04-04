package sec_test

import (
	"testing"

	"github.com/foomo/go/sec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestFilename verifies that Filename safely joins paths within a root directory,
// preventing path traversal attacks (gosec G304 / CWE-22).
func TestFilename(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		root    string
		elem    []string
		want    string
		wantErr string
	}{
		{
			name: "valid simple path",
			root: "/tmp",
			elem: []string{"file.txt"},
			want: "/tmp/file.txt",
		},
		{
			name: "valid nested path",
			root: "/tmp",
			elem: []string{"a", "b", "file.txt"},
			want: "/tmp/a/b/file.txt",
		},
		{
			name:    "traversal blocked",
			root:    "/tmp",
			elem:    []string{"../etc/passwd"},
			wantErr: "path traversal attempt",
		},
		{
			name:    "traversal via nested elements",
			root:    "/tmp",
			elem:    []string{"a", "../../etc/passwd"},
			wantErr: "path traversal attempt",
		},
		{
			name:    "empty root",
			root:    "",
			elem:    []string{"file.txt"},
			wantErr: "root required",
		},
		{
			name: "empty elem",
			root: "/tmp",
			want: "/tmp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := sec.Filename(tt.root, tt.elem...)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
