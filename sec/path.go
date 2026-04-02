package sec

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Filename joins a root directory with one or more path elements and ensures the resulting path is safely within the root.
func Filename(root string, elem ...string) (string, error) {
	if root == "" {
		return "", fmt.Errorf("root required")
	}

	fullPath := filepath.Join(root, filepath.Join(elem...))
	cleanPath := filepath.Clean(fullPath)

	// Must stay within root after cleaning
	if !strings.HasPrefix(cleanPath, root) {
		return "", fmt.Errorf("path traversal attempt")
	}

	return cleanPath, nil
}
