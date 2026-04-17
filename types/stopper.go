package types

import (
	"context"
)

type Stopper interface {
	Stop(ctx context.Context) error
}

type StopFunc func()

func (f StopFunc) Stop(ctx context.Context) error {
	f()
	return nil
}

type StopFuncErr func() error

func (f StopFuncErr) Stop(ctx context.Context) error {
	return f()
}

type StopFuncCtx func(context.Context)

func (f StopFuncCtx) Stop(ctx context.Context) error {
	f(ctx)
	return nil
}

type StopFuncCtxErr func(context.Context) error

func (f StopFuncCtxErr) Stop(ctx context.Context) error {
	return f(ctx)
}

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
