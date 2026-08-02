package slog_test

import (
	"log/slog"
	"testing"

	goslog "github.com/foomo/go/slog"
)

// TestWithGroupEmpty covers the empty-name short-circuit in WithGroup.
// slog.Logger.WithGroup("") short-circuits in stdlib and never reaches the
// handler, so call the handler directly.
func TestWithGroupEmpty(t *testing.T) {
	h := goslog.NewTestHandler(t)
	if got := h.WithGroup(""); got != h {
		t.Error("WithGroup(\"\") should return the same handler")
	}

	_ = slog.New(h)
}
