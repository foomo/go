package runtime

// Memo caches a value of type T per call site.
//
// It exists so that per-site derived data — a span name, a preformatted
// attribute slice, a logger — is computed once for the life of the process
// rather than on every call. Each Memo is its own namespace, so several may
// derive different values from the same call site. The zero value is ready to
// use and a Memo must not be copied after first use.
type Memo[T any] struct {
	c cache[T]
}

// Get returns the memoised value for the call site skip levels above Get's
// caller: skip=0 is the function that called Get.
//
// derive is invoked at most once per call site under normal conditions, but
// concurrent first calls from the same site may each run it and race to store;
// derive must therefore be pure and its result safe to share. Values handed
// out by Get are shared across all calls from that site and must be treated as
// read-only.
func (mo *Memo[T]) Get(skip int, derive func(Frame) T) T {
	pc, f := at(skip + 1)
	if v, ok := mo.c.load(pc); ok {
		return v
	}

	v := derive(f)
	if pc != 0 { // unresolvable stack: derive, but never poison the cache
		mo.c.store(pc, v)
	}

	return v
}
