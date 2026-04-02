package testing

import (
	"encoding/pem"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func writeFile(tb testing.TB, data string, name string) {
	tb.Helper()

	err := os.WriteFile(name, []byte(data), 0o600)
	require.NoError(tb, err)
}

func writePEMToTempFile(tb testing.TB, tempDir string, pattern string, block *pem.Block) string {
	tb.Helper()

	file, err := os.CreateTemp(tempDir, pattern)
	require.NoError(tb, err)

	defer func() {
		_ = file.Close()
	}()

	require.NoError(tb, pem.Encode(file, block))

	return file.Name()
}
