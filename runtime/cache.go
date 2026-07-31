package runtime

import (
	"maps"
	"sync"
	"sync/atomic"
)

// cache is a copy-on-write map keyed by program counter.
//
// sync.Map is the obvious choice here and the wrong one: its Load takes an
// any, so every lookup boxes the uintptr key and allocates. Because the key
// space is bounded by the number of call sites and saturates within the first
// moments of a process, copy-on-write is a better fit — reads are a lock-free
// pointer load and a map index, and the O(n) clone on insert happens at most
// once per call site.
type cache[T any] struct {
	mu sync.Mutex // serialises writers only
	m  atomic.Pointer[map[uintptr]T]
}

func (c *cache[T]) load(pc uintptr) (T, bool) {
	if m := c.m.Load(); m != nil {
		v, ok := (*m)[pc]
		return v, ok
	}

	var zero T

	return zero, false
}

// store inserts pc if absent. Concurrent first calls from the same site are
// harmless: the second simply overwrites an equal value.
func (c *cache[T]) store(pc uintptr, v T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	old := c.m.Load()

	next := make(map[uintptr]T, mapLen(old)+1)
	if old != nil {
		maps.Copy(next, *old)
	}

	next[pc] = v
	c.m.Store(&next)
}

func mapLen[T any](m *map[uintptr]T) int {
	if m == nil {
		return 0
	}

	return len(*m)
}
