package os

import (
	"os"
	"path/filepath"
	"strings"
)

// Expand replaces a leading "~/" in s with the current user's home directory,
// then expands any ${var} or $var references using [os.ExpandEnv].
func Expand(s string) (string, error) {
	if strings.HasPrefix(s, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		s = filepath.Join(home, s[2:])
	}

	return os.ExpandEnv(s), nil
}
