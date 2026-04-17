package types

import (
	"context"
)

type Closer interface {
	Close(ctx context.Context) error
}

type CloseFunc func()

func (f CloseFunc) Close(ctx context.Context) error {
	f()
	return nil
}

type CloseFuncErr func() error

func (f CloseFuncErr) Close(ctx context.Context) error {
	return f()
}

type CloseFuncCtx func(context.Context)

func (f CloseFuncCtx) Close(ctx context.Context) error {
	f(ctx)
	return nil
}

type CloseFuncCtxErr func(context.Context) error

func (f CloseFuncCtxErr) Close(ctx context.Context) error {
	return f(ctx)
}

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
