# testing

Testing utilities including tag-based test filtering, network helpers, and cryptographic key generation.

## Import

```go
import testingx "github.com/foomo/go/testing"
import tagx "github.com/foomo/go/testing/tag"
```

## Tag-Based Test Filtering

### Why not build tags?

Go's only built-in mechanism for selectively running tests is [build tags](https://pkg.go.dev/cmd/go#hdr-Build_constraints) (`//go:build integration`). Build tags work at the **file level** — an entire file is either compiled or excluded. This has several drawbacks:

- You must split tests across separate files just to control which ones run.
- You can only **include** files that match; there is no way to **exclude** a specific tag while running everything else.
- Tags are resolved at compile time, so switching between tag sets requires a full rebuild.

This package takes a different approach: tag filtering happens at **runtime** via the `GO_TEST_TAGS` environment variable. Tests declare their tags with a single function call and the filter supports both **includes** and **excludes**, giving you fine-grained control without restructuring your test files.

### How it works

Set the `GO_TEST_TAGS` environment variable to a comma-separated list of tags. Prefix a tag with `-` to exclude it.

| `GO_TEST_TAGS` | Test tagged `integration` | Test tagged `short` |
|---|---|---|
| *(unset)* | runs | runs |
| `integration` | runs | skipped |
| `short` | skipped | runs |
| `integration,short` | runs | runs |
| `-integration` | skipped | runs |
| `short,-integration` | skipped | runs |

**Rules:**

- When `GO_TEST_TAGS` is **unset**, all tagged tests run (no filtering).
- When only **include** tags are listed, a test must match at least one to run.
- When only **exclude** tags are listed (all prefixed with `-`), any matching test is skipped; non-matching tests still run.
- **Excludes always win** — if a test matches both an include and an exclude, it is skipped.

### Tags

```go
func Tags(t *testing.T, tags ...tagx.Tag)
```

Marks a test with the given tags and skips it if the current `GO_TEST_TAGS` rules say so. Call it as the first line in your test function.

### SkipTags

```go
func SkipTags(tags ...tagx.Tag) bool
```

Returns `true` if the given tags should be skipped under the current `GO_TEST_TAGS` value. Useful when you need the boolean result without calling `t.Skip` yourself.

### Tag Constants

```go
const (
	Always      Tag = "always"
	Benchmark   Tag = "benchmark"
	Blocking    Tag = "blocking"
	CI          Tag = "ci"
	Docker      Tag = "docker"
	E2          Tag = "e2e"
	Integration Tag = "integration"
	Load        Tag = "load"
	Parallel    Tag = "parallel"
	Performance Tag = "performance"
	Race        Tag = "race"
	Regression  Tag = "regression"
	Security    Tag = "security"
	Sequence    Tag = "sequence"
	Short       Tag = "short"
	Skip        Tag = "skip"
	Suite       Tag = "suite"
	Update      Tag = "update"
	// ... and more
)
```

### Examples

```go
func TestDatabaseIntegration(t *testing.T) {
	testingx.Tags(t, tagx.Integration)
	// test body ...
}

func TestQuickCheck(t *testing.T) {
	testingx.Tags(t, tagx.Short)
	// test body ...
}
```

```bash
# No filter — all tagged tests run
go test ./...

# Run only integration tests
GO_TEST_TAGS=integration go test ./...

# Run everything except integration tests
GO_TEST_TAGS=-integration go test ./...

# Run short and integration, but not performance
GO_TEST_TAGS=short,integration,-performance go test ./...
```

## Network Helpers

### FreePort

```go
func FreePort(t *testing.T) string
```

Returns a free port on localhost. Calls `t.Fatal` on error.

```go
func TestServer(t *testing.T) {
	addr := testingx.FreePort(t)
	// addr = "127.0.0.1:54321"
}
```

## Cryptographic Key Generation

Helpers for generating key pairs in tests. Keys are written to temporary files and cleaned up automatically.

### RSA

```go
func GenerateRSAPrivateKey(t *testing.T) *rsa.PrivateKey
func EncodeRSAPrivateKey(t *testing.T, privateKey *rsa.PrivateKey) string
func EncodeRSAPublicKey(t *testing.T, publicKey *rsa.PublicKey) string
func GenerateRSAKeyPair(t *testing.T) (public, private string)
func GenerateRSAPublicKey(t *testing.T, privateKey *rsa.PrivateKey, filePath string)
```

### ED25519

```go
func GenerateED25519PrivateKey(t *testing.T) ed25519.PrivateKey
func EncodeED25519PrivateKey(t *testing.T, privateKey ed25519.PrivateKey) string
func EncodeED25519PublicKey(t *testing.T, publicKey ed25519.PublicKey) string
func GenerateED25519KeyPair(t *testing.T) (public, private string)
func GenerateED25519PublicKey(t *testing.T, privateKey ed25519.PrivateKey, filePath string)
```

### Example

```go
func TestJWT(t *testing.T) {
	// Generate RSA key pair (files cleaned up after test)
	publicKeyPath, privateKeyPath := testingx.GenerateRSAKeyPair(t)

	// Or generate ED25519 keys
	edPub, edPriv := testingx.GenerateED25519KeyPair(t)
}
```
