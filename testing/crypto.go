package testing

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func NewRSAKeys(t *testing.T) (public, private string) { //nolint:nonamedreturns // clearifies return values
	t.Helper()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	tempDir := t.TempDir()
	privatePem, err := os.CreateTemp(tempDir, "private.*.pem")
	require.NoError(t, err)

	defer privatePem.Close()

	require.NoError(t, pem.Encode(privatePem, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}))

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	require.NoError(t, err)

	publicPem, err := os.CreateTemp(tempDir, "public.*.pem")
	require.NoError(t, err)

	defer publicPem.Close()

	require.NoError(t, pem.Encode(publicPem, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}))

	return publicPem.Name(), privatePem.Name()
}
