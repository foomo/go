package types

import (
	"context"
)

// Shutdowner is implemented by types that perform a graceful shutdown.
type Shutdowner interface {
	Shutdown(ctx context.Context) error
}

// ShutdownFunc adapts a func() to a Shutdowner, ignoring ctx and always returning nil.
type ShutdownFunc func()

// Shutdown calls f and returns nil.
func (f ShutdownFunc) Shutdown(ctx context.Context) error {
	f()
	return nil
}

// ShutdownFuncErr adapts a func() error to a Shutdowner, ignoring ctx.
type ShutdownFuncErr func() error

// Shutdown calls f and returns its result.
func (f ShutdownFuncErr) Shutdown(ctx context.Context) error {
	return f()
}

// ShutdownFuncCtx adapts a func(context.Context) to a Shutdowner, always returning nil.
type ShutdownFuncCtx func(context.Context)

// Shutdown calls f with ctx and returns nil.
func (f ShutdownFuncCtx) Shutdown(ctx context.Context) error {
	f(ctx)
	return nil
}

// ShutdownFuncCtxErr adapts a func(context.Context) error to a Shutdowner.
type ShutdownFuncCtxErr func(context.Context) error

// Shutdown calls f with ctx and returns its result.
func (f ShutdownFuncCtxErr) Shutdown(ctx context.Context) error {
	return f(ctx)
}

// AsShutdowner adapts v to a Shutdowner if v is a Shutdowner or one of func(),
// func() error, func(context.Context), or func(context.Context) error.
// It reports false if v matches none of these shapes.
func AsShutdowner(v any) (Shutdowner, bool) {
	switch f := v.(type) {
	case func():
		return ShutdownFunc(f), true
	case func() error:
		return ShutdownFuncErr(f), true
	case func(context.Context):
		return ShutdownFuncCtx(f), true
	case func(context.Context) error:
		return ShutdownFuncCtxErr(f), true
	case Shutdowner:
		return f, true
	default:
		return nil, false
	}
}
