package testing

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/stretchr/testify/require"
)

// GenerateRSAPrivateKey generates a new RSA private key.
func GenerateRSAPrivateKey(t *testing.T) *rsa.PrivateKey {
	t.Helper()

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	return privateKey
}

// EncodeRSAPrivateKey encodes an RSA private key to PEM format.
func EncodeRSAPrivateKey(t *testing.T, privateKey *rsa.PrivateKey) string {
	t.Helper()

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	return string(pem.EncodeToMemory(block))
}

// EncodeRSAPublicKey encodes an RSA public key to PEM format.
func EncodeRSAPublicKey(t *testing.T, publicKey *rsa.PublicKey) string {
	t.Helper()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	require.NoError(t, err)

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	return string(pem.EncodeToMemory(block))
}

// GenerateRSAKeyPair generates a new RSA key pair and returns the paths to the public and private key files.
//
// Example usage:
//
//	func TestExample(t *testing.T) {
//		publicKeyPath, privateKeyPath := GenerateRSAKeyPair(t)
//		// Use the key paths for testing crypto operations
//		// Files are automatically cleaned up when test completes
//	}
func GenerateRSAKeyPair(t *testing.T) (public, private string) { //nolint:nonamedreturns // clearifies return values
	t.Helper()

	privateKey := GenerateRSAPrivateKey(t)
	tempDir := t.TempDir()

	privateKeyPath := writePEMToTempFile(t, tempDir, "*_rsa", &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	require.NoError(t, err)

	publicKeyPath := writePEMToTempFile(t, tempDir, "*_rsa.pub", &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	return publicKeyPath, privateKeyPath
}

// GenerateRSAPublicKey generates a new public key from an RSA private key and writes it to a file.
func GenerateRSAPublicKey(t *testing.T, privateKey *rsa.PrivateKey, filePath string) {
	t.Helper()

	publicKeyPem := EncodeRSAPublicKey(t, &privateKey.PublicKey)
	writeFile(t, publicKeyPem, filePath)
}
