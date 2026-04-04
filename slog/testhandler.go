package slog

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"testing"
)

type (
	testHandler struct {
		w      io.Writer
		level  slog.Leveler
		groups []string
		attrs  []slog.Attr
	}
	// TestHandlerOption is a functional option for [NewTestHandler].
	TestHandlerOption func(*testHandler)
)

// TestHandlerWithLevel sets the minimum log level for the test handler.
func TestHandlerWithLevel(level slog.Leveler) TestHandlerOption {
	return func(h *testHandler) {
		h.level = level
	}
}

// NewTestHandler returns an [slog.Handler] that writes log records
// to tb.Output in a compact format: [LEVEL] msg key=value ...
// Defaults to [slog.LevelDebug].
func NewTestHandler(tb testing.TB, opts ...TestHandlerOption) slog.Handler {
	tb.Helper()

	h := &testHandler{w: tb.Output(), level: slog.LevelDebug}
	for _, opt := range opts {
		opt(h)
	}

	return h
}

func (h *testHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level.Level()
}

func (h *testHandler) Handle(_ context.Context, r slog.Record) error {
	var b strings.Builder

	if r.PC != 0 {
		f, _ := runtime.CallersFrames([]uintptr{r.PC}).Next()
		_, _ = fmt.Fprintf(&b, "%s:%d: ", filepath.Base(f.File), f.Line)
	}

	_, _ = fmt.Fprintf(&b, "[%s] %s", r.Level.String(), r.Message)
	writeAttrs(&b, h.groups, h.attrs)
	r.Attrs(func(a slog.Attr) bool {
		writeAttrs(&b, h.groups, []slog.Attr{a})
		return true
	})
	b.WriteByte('\n')
	_, err := io.WriteString(h.w, b.String())

	return err
}

func (h *testHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &testHandler{
		w:      h.w,
		level:  h.level,
		groups: h.groups,
		attrs:  append(slices.Clone(h.attrs), attrs...),
	}
}

func (h *testHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}

	return &testHandler{
		w:      h.w,
		level:  h.level,
		groups: append(slices.Clone(h.groups), name),
	}
}

func writeAttrs(b *strings.Builder, groups []string, attrs []slog.Attr) {
	prefix := ""
	if len(groups) > 0 {
		prefix = strings.Join(groups, ".") + "."
	}

	for _, a := range attrs {
		a.Value = a.Value.Resolve()
		if a.Equal(slog.Attr{}) {
			continue
		}

		if a.Value.Kind() == slog.KindGroup {
			writeAttrs(b, append(groups, a.Key), a.Value.Group())
			continue
		}

		fmt.Fprintf(b, " %s%s=%s", prefix, a.Key, a.Value.String())
	}
}
