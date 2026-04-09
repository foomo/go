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

	if strings.ContainsRune(root, '\x00') {
		return "", fmt.Errorf("null byte in root path")
	}

	for _, e := range elem {
		if strings.ContainsRune(e, '\x00') {
			return "", fmt.Errorf("null byte in path element")
		}
	}

	fullPath := filepath.Join(root, filepath.Join(elem...))
	cleanPath := filepath.Clean(fullPath)

	// Must stay within root after cleaning
	if !strings.HasPrefix(cleanPath, root) {
		return "", fmt.Errorf("path traversal attempt")
	}

	return cleanPath, nil
}
