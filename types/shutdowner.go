package types

import (
	"context"
)

type Shutdowner interface {
	Shutdown(ctx context.Context) error
}

type ShutdownFunc func()

func (f ShutdownFunc) Shutdown(ctx context.Context) error {
	f()
	return nil
}

type ShutdownFuncErr func() error

func (f ShutdownFuncErr) Shutdown(ctx context.Context) error {
	return f()
}

type ShutdownFuncCtx func(context.Context)

func (f ShutdownFuncCtx) Shutdown(ctx context.Context) error {
	f(ctx)
	return nil
}

type ShutdownFuncCtxErr func(context.Context) error

func (f ShutdownFuncCtxErr) Shutdown(ctx context.Context) error {
	return f(ctx)
}

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
