package types

import (
	"context"
)

// Stopper is implemented by types that perform a stop action.
type Stopper interface {
	Stop(ctx context.Context) error
}

// StopFunc adapts a func() to a Stopper, ignoring ctx and always returning nil.
type StopFunc func()

// Stop calls f and returns nil.
func (f StopFunc) Stop(ctx context.Context) error {
	f()
	return nil
}

// StopFuncErr adapts a func() error to a Stopper, ignoring ctx.
type StopFuncErr func() error

// Stop calls f and returns its result.
func (f StopFuncErr) Stop(ctx context.Context) error {
	return f()
}

// StopFuncCtx adapts a func(context.Context) to a Stopper, always returning nil.
type StopFuncCtx func(context.Context)

// Stop calls f with ctx and returns nil.
func (f StopFuncCtx) Stop(ctx context.Context) error {
	f(ctx)
	return nil
}

// StopFuncCtxErr adapts a func(context.Context) error to a Stopper.
type StopFuncCtxErr func(context.Context) error

// Stop calls f with ctx and returns its result.
func (f StopFuncCtxErr) Stop(ctx context.Context) error {
	return f(ctx)
}

// AsStopper adapts v to a Stopper if v is a Stopper or one of func(),
// func() error, func(context.Context), or func(context.Context) error.
// It reports false if v matches none of these shapes.
func AsStopper(v any) (Stopper, bool) {
	switch f := v.(type) {
	case func():
		return StopFunc(f), true
	case func() error:
		return StopFuncErr(f), true
	case func(context.Context):
		return StopFuncCtx(f), true
	case func(context.Context) error:
		return StopFuncCtxErr(f), true
	case Stopper:
		return f, true
	default:
		return nil, false
	}
}
