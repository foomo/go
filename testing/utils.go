package testing

import (
	"encoding/pem"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// writeKeyToFile writes a PEM-encoded key to a file.
func writeKeyToFile(tb testing.TB, pemData string, filePath string) {
	tb.Helper()

	err := os.WriteFile(filePath, []byte(pemData), 0o600)
	require.NoError(tb, err)
}

// writePEMToTempFile writes a PEM block to a temporary file and returns the file path.
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
