package testing

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/stretchr/testify/require"
)

// GenerateED25519PrivateKey generates a new ED25519 private key.
func GenerateED25519PrivateKey(tb testing.TB) ed25519.PrivateKey {
	tb.Helper()

	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(tb, err)

	return privateKey
}

// EncodeED25519PrivateKey encodes an ED25519 private key to PEM format.
func EncodeED25519PrivateKey(tb testing.TB, privateKey ed25519.PrivateKey) string {
	tb.Helper()

	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	require.NoError(tb, err)

	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	return string(pem.EncodeToMemory(block))
}

// EncodeED25519PublicKey encodes an ED25519 public key to PEM format.
func EncodeED25519PublicKey(tb testing.TB, publicKey ed25519.PublicKey) string {
	tb.Helper()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	require.NoError(tb, err)

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	return string(pem.EncodeToMemory(block))
}

// GenerateED25519KeyPair generates a new ED25519 key pair and returns the paths to the public and private key files.
//
// Example usage:
//
//	func TestExample(t *testing.T) {
//		publicKeyPath, privateKeyPath := GenerateED25519KeyPair(t)
//		// Use the key paths for testing crypto operations
//		// Files are automatically cleaned up when test completes
//	}
func GenerateED25519KeyPair(tb testing.TB) (public, private string) { //nolint:nonamedreturns // clearifies return values
	tb.Helper()

	privateKey := GenerateED25519PrivateKey(tb)
	publicKey := privateKey.Public().(ed25519.PublicKey) //nolint:forcetypeassert // cast is safe
	tempDir := tb.TempDir()

	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	require.NoError(tb, err)

	privateKeyPath := writePEMToTempFile(tb, tempDir, "*_ed25519", &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	require.NoError(tb, err)

	publicKeyPath := writePEMToTempFile(tb, tempDir, "*_ed25519.pub", &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	return publicKeyPath, privateKeyPath
}

// GenerateED25519PublicKey generates a new public key from an ED25519 private key and writes it to a file.
func GenerateED25519PublicKey(tb testing.TB, privateKey ed25519.PrivateKey, filePath string) {
	tb.Helper()

	publicKey := privateKey.Public().(ed25519.PublicKey) //nolint:forcetypeassert // cast is safe
	publicKeyPem := EncodeED25519PublicKey(tb, publicKey)
	writeFile(tb, publicKeyPem, filePath)
}
