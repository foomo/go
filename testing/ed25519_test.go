package testing_test

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"testing"

	testingx "github.com/foomo/go/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateED25519KeyPair(t *testing.T) {
	t.Parallel()

	publicPath, privatePath := testingx.GenerateED25519KeyPair(t)

	require.FileExists(t, publicPath)
	require.FileExists(t, privatePath)

	privateData, err := os.ReadFile(privatePath)
	require.NoError(t, err)

	privateBlock, rest := pem.Decode(privateData)
	require.NotNil(t, privateBlock)
	assert.Empty(t, rest)
	assert.Equal(t, "PRIVATE KEY", privateBlock.Type)

	key, err := x509.ParsePKCS8PrivateKey(privateBlock.Bytes)
	require.NoError(t, err)
	assert.IsType(t, ed25519.PrivateKey{}, key)

	publicData, err := os.ReadFile(publicPath)
	require.NoError(t, err)

	publicBlock, rest := pem.Decode(publicData)
	require.NotNil(t, publicBlock)
	assert.Empty(t, rest)
	assert.Equal(t, "PUBLIC KEY", publicBlock.Type)
}

func TestEncodeED25519Keys(t *testing.T) {
	t.Parallel()

	privateKey := testingx.GenerateED25519PrivateKey(t)

	privatePEM := testingx.EncodeED25519PrivateKey(t, privateKey)
	block, _ := pem.Decode([]byte(privatePEM))
	require.NotNil(t, block)
	assert.Equal(t, "PRIVATE KEY", block.Type)

	publicKey, ok := privateKey.Public().(ed25519.PublicKey)
	require.True(t, ok)

	publicPEM := testingx.EncodeED25519PublicKey(t, publicKey)
	block, _ = pem.Decode([]byte(publicPEM))
	require.NotNil(t, block)
	assert.Equal(t, "PUBLIC KEY", block.Type)
}

func TestGenerateED25519PublicKey(t *testing.T) {
	t.Parallel()

	privateKey := testingx.GenerateED25519PrivateKey(t)
	path := filepath.Join(t.TempDir(), "id_ed25519.pub")

	testingx.GenerateED25519PublicKey(t, privateKey, path)

	require.FileExists(t, path)

	data, err := os.ReadFile(path)
	require.NoError(t, err)

	block, _ := pem.Decode(data)
	require.NotNil(t, block)
	assert.Equal(t, "PUBLIC KEY", block.Type)
}
