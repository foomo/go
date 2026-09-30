package types

import (
	"context"
)

// Unsubscriber is implemented by types that cancel a subscription.
type Unsubscriber interface {
	Unsubscribe(ctx context.Context) error
}

// UnsubscribeFunc adapts a func() to an Unsubscriber, ignoring ctx and always returning nil.
type UnsubscribeFunc func()

// Unsubscribe calls f and returns nil.
func (f UnsubscribeFunc) Unsubscribe(ctx context.Context) error {
	f()
	return nil
}

// UnsubscribeFuncErr adapts a func() error to an Unsubscriber, ignoring ctx.
type UnsubscribeFuncErr func() error

// Unsubscribe calls f and returns its result.
func (f UnsubscribeFuncErr) Unsubscribe(ctx context.Context) error {
	return f()
}

// UnsubscribeFuncCtx adapts a func(context.Context) to an Unsubscriber, always returning nil.
type UnsubscribeFuncCtx func(context.Context)

// Unsubscribe calls f with ctx and returns nil.
func (f UnsubscribeFuncCtx) Unsubscribe(ctx context.Context) error {
	f(ctx)
	return nil
}

// UnsubscribeFuncCtxErr adapts a func(context.Context) error to an Unsubscriber.
type UnsubscribeFuncCtxErr func(context.Context) error

// Unsubscribe calls f with ctx and returns its result.
func (f UnsubscribeFuncCtxErr) Unsubscribe(ctx context.Context) error {
	return f(ctx)
}

// AsUnsubscriber adapts v to an Unsubscriber if v is an Unsubscriber or one of
// func(), func() error, func(context.Context), or func(context.Context) error.
// It reports false if v matches none of these shapes.
func AsUnsubscriber(v any) (Unsubscriber, bool) {
	switch f := v.(type) {
	case func():
		return UnsubscribeFunc(f), true
	case func() error:
		return UnsubscribeFuncErr(f), true
	case func(context.Context):
		return UnsubscribeFuncCtx(f), true
	case func(context.Context) error:
		return UnsubscribeFuncCtxErr(f), true
	case Unsubscriber:
		return f, true
	default:
		return nil, false
	}
}
