package runtime_test

import (
	"sync"
	"testing"

	goruntime "github.com/foomo/go/runtime"
)

func TestMemoPerCallSite(t *testing.T) {
	var (
		m     goruntime.Memo[string]
		calls int
	)

	derive := func(f goruntime.Frame) string { calls++; return f.Short() }

	// Same call site (the loop body), repeated: derive must run once.
	for range 3 {
		if v := m.Get(0, derive); v != "TestMemoPerCallSite" {
			t.Fatalf("got %q", v)
		}
	}

	if calls != 1 {
		t.Fatalf("derive ran %d times, want 1", calls)
	}

	// A second, textually distinct call site in the same function is a
	// distinct program counter and therefore a distinct entry.
	if v := m.Get(0, derive); v != "TestMemoPerCallSite" || calls != 2 {
		t.Fatalf("got %q after %d derives", v, calls)
	}
}

func TestMemoConcurrent(t *testing.T) {
	var (
		m  goruntime.Memo[int]
		wg sync.WaitGroup
	)
	for range 64 {
		wg.Go(func() {
			if v := m.Get(0, func(goruntime.Frame) int { return 7 }); v != 7 {
				t.Errorf("got %d", v)
			}
		})
	}

	wg.Wait()
}

func BenchmarkMemoGet(b *testing.B) {
	var m goruntime.Memo[string]

	derive := func(f goruntime.Frame) string { return f.Short() }
	m.Get(0, derive)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = m.Get(0, derive)
	}
}
