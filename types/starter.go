package types

import (
	"context"
)

type Starter interface {
	Start(ctx context.Context) error
}

type StartFunc func()

func (f StartFunc) Start(ctx context.Context) error {
	f()
	return nil
}

type StartFuncErr func() error

func (f StartFuncErr) Start(ctx context.Context) error {
	return f()
}

type StartFuncCtx func(context.Context)

func (f StartFuncCtx) Start(ctx context.Context) error {
	f(ctx)
	return nil
}

type StartFuncCtxErr func(context.Context) error

func (f StartFuncCtxErr) Start(ctx context.Context) error {
	return f(ctx)
}

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
