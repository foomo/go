package types

import (
	"context"
)

// Starter is implemented by types that perform a startup action.
type Starter interface {
	Start(ctx context.Context) error
}

// StartFunc adapts a func() to a Starter, ignoring ctx and always returning nil.
type StartFunc func()

// Start calls f and returns nil.
func (f StartFunc) Start(ctx context.Context) error {
	f()
	return nil
}

// StartFuncErr adapts a func() error to a Starter, ignoring ctx.
type StartFuncErr func() error

// Start calls f and returns its result.
func (f StartFuncErr) Start(ctx context.Context) error {
	return f()
}

// StartFuncCtx adapts a func(context.Context) to a Starter, always returning nil.
type StartFuncCtx func(context.Context)

// Start calls f with ctx and returns nil.
func (f StartFuncCtx) Start(ctx context.Context) error {
	f(ctx)
	return nil
}

// StartFuncCtxErr adapts a func(context.Context) error to a Starter.
type StartFuncCtxErr func(context.Context) error

// Start calls f with ctx and returns its result.
func (f StartFuncCtxErr) Start(ctx context.Context) error {
	return f(ctx)
}

// AsStarter adapts v to a Starter if v is a Starter or one of func(),
// func() error, func(context.Context), or func(context.Context) error.
// It reports false if v matches none of these shapes.
func AsStarter(v any) (Starter, bool) {
	switch f := v.(type) {
	case func():
		return StartFunc(f), true
	case func() error:
		return StartFuncErr(f), true
	case func(context.Context):
		return StartFuncCtx(f), true
	case func(context.Context) error:
		return StartFuncCtxErr(f), true
	case Starter:
		return f, true
	default:
		return nil, false
	}
}
