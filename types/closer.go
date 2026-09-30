package types

import (
	"context"
)

// Closer is implemented by types that release resources on Close.
type Closer interface {
	Close(ctx context.Context) error
}

// CloseFunc adapts a func() to a Closer, ignoring ctx and always returning nil.
type CloseFunc func()

// Close calls f and returns nil.
func (f CloseFunc) Close(ctx context.Context) error {
	f()
	return nil
}

// CloseFuncErr adapts a func() error to a Closer, ignoring ctx.
type CloseFuncErr func() error

// Close calls f and returns its result.
func (f CloseFuncErr) Close(ctx context.Context) error {
	return f()
}

// CloseFuncCtx adapts a func(context.Context) to a Closer, always returning nil.
type CloseFuncCtx func(context.Context)

// Close calls f with ctx and returns nil.
func (f CloseFuncCtx) Close(ctx context.Context) error {
	f(ctx)
	return nil
}

// CloseFuncCtxErr adapts a func(context.Context) error to a Closer.
type CloseFuncCtxErr func(context.Context) error

// Close calls f with ctx and returns its result.
func (f CloseFuncCtxErr) Close(ctx context.Context) error {
	return f(ctx)
}

// AsCloser adapts v to a Closer if v is a Closer or one of func(),
// func() error, func(context.Context), or func(context.Context) error.
// It reports false if v matches none of these shapes.
func AsCloser(v any) (Closer, bool) {
	switch f := v.(type) {
	case func():
		return CloseFunc(f), true
	case func() error:
		return CloseFuncErr(f), true
	case func(context.Context):
		return CloseFuncCtx(f), true
	case func(context.Context) error:
		return CloseFuncCtxErr(f), true
	case Closer:
		return f, true
	default:
		return nil, false
	}
}
