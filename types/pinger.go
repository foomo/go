package types

import (
	"context"
)

type Pinger interface {
	Ping(ctx context.Context) error
}

type PingFunc func()

func (f PingFunc) Ping(ctx context.Context) error {
	f()
	return nil
}

type PingFuncErr func() error

func (f PingFuncErr) Ping(ctx context.Context) error {
	return f()
}

type PingFuncCtx func(context.Context)

func (f PingFuncCtx) Ping(ctx context.Context) error {
	f(ctx)
	return nil
}

type PingFuncCtxErr func(context.Context) error

func (f PingFuncCtxErr) Ping(ctx context.Context) error {
	return f(ctx)
}

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
