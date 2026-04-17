package types

import (
	"context"
)

type Unsubscriber interface {
	Unsubscribe(ctx context.Context) error
}

type UnsubscribeFunc func()

func (f UnsubscribeFunc) Unsubscribe(ctx context.Context) error {
	f()
	return nil
}

type UnsubscribeFuncErr func() error

func (f UnsubscribeFuncErr) Unsubscribe(ctx context.Context) error {
	return f()
}

type UnsubscribeFuncCtx func(context.Context)

func (f UnsubscribeFuncCtx) Unsubscribe(ctx context.Context) error {
	f(ctx)
	return nil
}

type UnsubscribeFuncCtxErr func(context.Context) error

func (f UnsubscribeFuncCtxErr) Unsubscribe(ctx context.Context) error {
	return f(ctx)
}

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
