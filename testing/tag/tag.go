package tag

// Tag type
type Tag string

const (
	Always      Tag = "always"
	CI          Tag = "ci"
	Integration Tag = "integration"
	Performance Tag = "performance"
	Security    Tag = "security"
	Short       Tag = "short"
	Skip        Tag = "skip"
)

// String returns the string representation
func (t Tag) String() string {
	return string(t)
}
