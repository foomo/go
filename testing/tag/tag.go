package tag

// Tag represents a type for categorizing tasks, processes, or features using predefined string identifiers.
type Tag string

const (
	Always      Tag = "always"
	Benchmark   Tag = "benchmark"
	Blocking    Tag = "blocking"
	CI          Tag = "ci"
	CPU         Tag = "profile"
	Docker      Tag = "docker"
	E2          Tag = "e2e"
	Generate    Tag = "generate"
	Integration Tag = "integration"
	Load        Tag = "load"
	Memory      Tag = "memory"
	MemoryLeak  Tag = "unsafe"
	Mutex       Tag = "mutex"
	MutexLeak   Tag = "mutexleak"
	Parallel    Tag = "parallel"
	Performance Tag = "performance"
	Profile     Tag = "profile"
	Race        Tag = "race"
	Regression  Tag = "regression"
	Safe        Tag = "safe"
	Security    Tag = "security"
	Sequence    Tag = "sequence"
	Short       Tag = "short"
	Skip        Tag = "skip"
	Suite       Tag = "suite"
	Table       Tag = "table"
	Unsafe      Tag = "unsafe"
	Update      Tag = "update"
)

// String returns the string representation
func (t Tag) String() string {
	return string(t)
}
