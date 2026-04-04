# sec

Security utilities.

## Import

```go
import secx "github.com/foomo/go/sec"
```

## API

### Filename

```go
func Filename(root string, elem ...string) (string, error)
```

Safely joins a root directory with one or more path elements and ensures the resulting path stays within the root. Returns an error if the root is empty or if the resolved path escapes the root directory (e.g. via `../` traversal).

This is the recommended way to construct file paths from user-supplied input, addressing [gosec G304](https://cwe.mitre.org/data/definitions/22.html) (CWE-22: Improper Limitation of a Pathname to a Restricted Directory).

## Example

```go
// Safe path — stays within root
path, err := secx.Filename("/srv/data", "users", "profile.json")
// path = "/srv/data/users/profile.json", err = nil

// Traversal attempt — blocked
path, err = secx.Filename("/srv/data", "../etc/passwd")
// path = "", err = "path traversal attempt"

// Empty root — rejected
path, err = secx.Filename("", "file.txt")
// path = "", err = "root required"
```
