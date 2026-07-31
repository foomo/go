package runtime

import (
	stdruntime "runtime"
	"strings"
)

// frames memoises symbolisation itself. Program counters are stable for the
// lifetime of the process, so this map is bounded by the number of distinct
// call sites that reach this package.
var frames cache[Frame]

// CallFrame returns the call site skip levels above its own caller: skip=0 is the
// function that called Caller, skip=1 that function's caller, and so on.
//
// Symbolisation is memoised per program counter, so the steady-state cost is
// one runtime.Callers into a stack array plus a map load, with no allocation.
func CallFrame(skip int) Frame {
	_, f := at(skip + 1)
	return f
}

// at resolves the call site skip levels above at's caller (skip=0 being at's
// immediate caller) and returns its program counter alongside the frame. The
// pc doubles as the cache key for Memo.
func at(skip int) (uintptr, Frame) {
	var pcs [1]uintptr
	// +2 accounts for stdruntime.Callers itself and for at, so that skip=0
	// resolves to at's caller. Inlined frames are counted logically, so this
	// arithmetic holds regardless of inlining decisions.
	if stdruntime.Callers(skip+2, pcs[:]) == 0 {
		return 0, Frame{}
	}

	pc := pcs[0]
	if f, ok := frames.load(pc); ok {
		return pc, f
	}

	return pc, resolve(pc)
}

// resolve symbolises pc and caches the result. It is deliberately a separate,
// non-inlined function: stdruntime.CallersFrames retains the slice it is given,
// so calling it in at would make at's pc buffer escape to the heap on every
// call — escape analysis is not path-sensitive and cannot see that this branch
// is taken at most once per call site.
//
//go:noinline
func resolve(pc uintptr) Frame {
	// CallersFrames expands inlined frames and adjusts return PCs to call PCs.
	// stdruntime.FuncForPC on the same pc would report the inlining parent,
	// silently attributing the call site to whatever function absorbed it.
	fr, _ := stdruntime.CallersFrames([]uintptr{pc}).Next()
	f := Frame{File: fr.File, Line: fr.Line}
	f.Pkg, f.Inst, f.Func = parseName(fr.Function)
	frames.store(pc, f)

	return f
}

// parseName decomposes a qualified function name as reported by the runtime.
// The shapes it handles:
//
//	pkg/path.Func
//	pkg/path.Recv.Func                  value receiver
//	pkg/path.(*Recv).Func               pointer receiver
//	pkg/path.(Recv[go.shape.int]).Func  generic receiver
//	pkg/path.Func.func1                 closure
//	pkg/path.(*Recv).Func-fm            method value
//	pkg/path.glob..func1                package-level var initializer
//
// One ambiguity is not resolvable from the string alone: a package path whose
// final element contains a dot, such as gopkg.in/yaml.v3, is split at that dot
// and its version element is mistaken for a receiver. The Go runtime has the
// same limitation. Module paths using the standard /v2 suffix are unaffected.
func parseName(fq string) (pkg, inst, fn string) { //nolint:nonamedreturns
	if fq == "" {
		return "", "", ""
	}
	// The package path ends at the first separator dot after the final slash.
	start := strings.LastIndexByte(fq, '/') + 1

	d := indexSep(fq[start:])
	if d < 0 {
		return fq, "", ""
	}

	pkg, rest := fq[:start+d], fq[start+d+1:]
	if rest == "" {
		return pkg, "", ""
	}

	// Parenthesised receiver. Parentheses do not nest in these names, so the
	// first ')' followed by '.' closes the receiver.
	if rest[0] == '(' {
		if end := strings.IndexByte(rest, ')'); end > 1 && end+1 < len(rest) && rest[end+1] == '.' {
			return pkg, strings.TrimPrefix(rest[1:end], "*"), rest[end+2:]
		}
	}

	// Bare receiver. "Recv.Func" and "Func.func1" are indistinguishable by
	// shape, so the second segment decides: a compiler-generated segment means
	// the first one was the function, not a receiver type.
	if segs := splitSep(rest); len(segs) >= 2 && !isGenerated(segs[1]) {
		return pkg, segs[0], strings.Join(segs[1:], ".")
	}

	return pkg, "", rest
}

// indexSep returns the index of the first separator dot in s, ignoring dots
// nested inside generic instantiation brackets such as [go.shape.int].
func indexSep(s string) int {
	depth := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[':
			depth++
		case ']':
			depth--
		case '.':
			if depth == 0 {
				return i
			}
		}
	}

	return -1
}

// splitSep splits s on separator dots, ignoring those inside brackets.
func splitSep(s string) []string {
	var (
		segs  []string
		depth int
		prev  int
	)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[':
			depth++
		case ']':
			depth--
		case '.':
			if depth == 0 {
				segs = append(segs, s[prev:i])
				prev = i + 1
			}
		}
	}

	return append(segs, s[prev:])
}

// isGenerated reports whether seg is a compiler-generated name segment rather
// than something written by a human: closure and wrapper markers, ordinal
// suffixes, and the empty segment produced by "glob..func1".
func isGenerated(seg string) bool {
	switch seg {
	case "", "glob", "stub":
		return true
	}

	for _, prefix := range [...]string{"func", "deferwrap", "gowrap"} {
		if suffix, ok := strings.CutPrefix(seg, prefix); ok && allDigits(suffix) {
			return true
		}
	}

	return allDigits(seg)
}

// allDigits reports whether s consists solely of ASCII digits. The empty
// string counts, which is what lets "func" match alongside "func1".
func allDigits(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}

	return true
}
