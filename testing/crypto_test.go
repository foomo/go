package testing_test

import (
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"

	testingx "github.com/foomo/go/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRSAKeys(t *testing.T) {
	t.Parallel()

	publicPath, privatePath := testingx.NewRSAKeys(t)

	// verify files exist
	require.FileExists(t, publicPath)
	require.FileExists(t, privatePath)

	// read and decode private key
	privateData, err := os.ReadFile(privatePath)
	require.NoError(t, err)

	privateBlock, rest := pem.Decode(privateData)
	require.NotNil(t, privateBlock)
	assert.Empty(t, rest)
	assert.Equal(t, "RSA PRIVATE KEY", privateBlock.Type)

	privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	require.NoError(t, err)
	assert.Equal(t, 2048, privateKey.N.BitLen())

	// read and decode public key
	publicData, err := os.ReadFile(publicPath)
	require.NoError(t, err)

	publicBlock, rest := pem.Decode(publicData)
	require.NotNil(t, publicBlock)
	assert.Empty(t, rest)
	assert.Equal(t, "PUBLIC KEY", publicBlock.Type)

	pubKey, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	require.NoError(t, err)

	// verify the public key matches the private key
	assert.True(t, privateKey.PublicKey.Equal(pubKey))
}
