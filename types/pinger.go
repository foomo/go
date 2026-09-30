package types

import (
	"context"
)

// Pinger is implemented by types that can check liveness of a connection
// or dependency on demand.
type Pinger interface {
	Ping(ctx context.Context) error
}

// PingFunc adapts a func() to a Pinger, ignoring ctx and always returning nil.
type PingFunc func()

// Ping calls f and returns nil.
func (f PingFunc) Ping(ctx context.Context) error {
	f()
	return nil
}

// PingFuncErr adapts a func() error to a Pinger, ignoring ctx.
type PingFuncErr func() error

// Ping calls f and returns its result.
func (f PingFuncErr) Ping(ctx context.Context) error {
	return f()
}

// PingFuncCtx adapts a func(context.Context) to a Pinger, always returning nil.
type PingFuncCtx func(context.Context)

// Ping calls f with ctx and returns nil.
func (f PingFuncCtx) Ping(ctx context.Context) error {
	f(ctx)
	return nil
}

// PingFuncCtxErr adapts a func(context.Context) error to a Pinger.
type PingFuncCtxErr func(context.Context) error

// Ping calls f with ctx and returns its result.
func (f PingFuncCtxErr) Ping(ctx context.Context) error {
	return f(ctx)
}

// AsPinger adapts v to a Pinger if v is a Pinger or one of func(),
// func() error, func(context.Context), or func(context.Context) error.
// It reports false if v matches none of these shapes.
func AsPinger(v any) (Pinger, bool) {
	switch f := v.(type) {
	case func():
		return PingFunc(f), true
	case func() error:
		return PingFuncErr(f), true
	case func(context.Context):
		return PingFuncCtx(f), true
	case func(context.Context) error:
		return PingFuncCtxErr(f), true
	case Pinger:
		return f, true
	default:
		return nil, false
	}
}
