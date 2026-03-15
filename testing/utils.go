package testing

import (
	"encoding/pem"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

// writeKeyToFile writes a PEM-encoded key to a file.
func writeKeyToFile(t *testing.T, pemData string, filePath string) {
	t.Helper()

	err := os.WriteFile(filePath, []byte(pemData), 0o600)
	require.NoError(t, err)
}

// writePEMToTempFile writes a PEM block to a temporary file and returns the file path.
func writePEMToTempFile(t *testing.T, tempDir string, pattern string, block *pem.Block) string {
	t.Helper()

	file, err := os.CreateTemp(tempDir, pattern)
	require.NoError(t, err)

	defer func() {
		_ = file.Close()
	}()

	require.NoError(t, pem.Encode(file, block))

	return file.Name()
}
