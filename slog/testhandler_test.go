package slog_test

import (
	"log/slog"
	"testing"

	slogx "github.com/foomo/go/slog"
)

func TestNewTestHandler(t *testing.T) {
	t.Run("default options", func(t *testing.T) {
		l := slog.New(slogx.NewTestHandler(t))
		l.Debug("debug message", "key", "value")
		l.Info("info message", "key", "value")
		l.Warn("warn message", "key", "value")
		l.Error("error message", "key", "value")
	})

	t.Run("custom level", func(t *testing.T) {
		l := slog.New(slogx.NewTestHandler(t, slogx.TestHandlerWithLevel(slog.LevelWarn)))
		l.Debug("should be filtered")
		l.Info("should be filtered")
		l.Warn("should appear", "key", "value")
		l.Error("should appear", "key", "value")
	})

	t.Run("with group and attrs", func(t *testing.T) {
		l := slog.New(slogx.NewTestHandler(t))
		l.WithGroup("request").With("method", "GET").Info("handled")
	})

	t.Run("nested groups", func(t *testing.T) {
		l := slog.New(slogx.NewTestHandler(t))
		l.WithGroup("http").WithGroup("request").Info("handled", "path", "/api")
	})

	t.Run("inline group attr", func(t *testing.T) {
		l := slog.New(slogx.NewTestHandler(t))
		l.Info("handled", slog.Group("request", "method", "GET", "path", "/api"))
	})
}
