package os

import (
	"os"
	"path/filepath"
	"strings"
)

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
