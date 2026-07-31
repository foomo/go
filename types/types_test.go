package types_test

import (
	"context"
	"errors"
	"testing"

	"github.com/foomo/go/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The six adapter families (Closer/Pinger/Starter/Stopper/Shutdowner/
// Unsubscriber) are structurally identical: four func adapters and one As<X>
// helper each. These tests exercise every adapter method and every branch of
// the As<X> type switch.

var errBoom = errors.New("boom")

func TestCloser(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("adapters", func(t *testing.T) {
		t.Parallel()

		require.NoError(t, types.CloseFunc(func() {}).Close(ctx))
		require.ErrorIs(t, types.CloseFuncErr(func() error { return errBoom }).Close(ctx), errBoom)
		require.NoError(t, types.CloseFuncCtx(func(context.Context) {}).Close(ctx))
		require.ErrorIs(t, types.CloseFuncCtxErr(func(context.Context) error { return errBoom }).Close(ctx), errBoom)
	})

	t.Run("AsCloser", func(t *testing.T) {
		t.Parallel()

		for name, v := range map[string]any{
			"func":       func() {},
			"funcErr":    func() error { return nil },
			"funcCtx":    func(context.Context) {},
			"funcCtxErr": func(context.Context) error { return nil },
			"closer":     types.CloseFunc(func() {}),
		} {
			c, ok := types.AsCloser(v)
			assert.Truef(t, ok, "case %s", name)
			require.NotNilf(t, c, "case %s", name)
			require.NoErrorf(t, c.Close(ctx), "case %s", name)
		}

		c, ok := types.AsCloser("nope")
		assert.False(t, ok)
		assert.Nil(t, c)
	})
}

func TestPinger(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("adapters", func(t *testing.T) {
		t.Parallel()

		require.NoError(t, types.PingFunc(func() {}).Ping(ctx))
		require.ErrorIs(t, types.PingFuncErr(func() error { return errBoom }).Ping(ctx), errBoom)
		require.NoError(t, types.PingFuncCtx(func(context.Context) {}).Ping(ctx))
		require.ErrorIs(t, types.PingFuncCtxErr(func(context.Context) error { return errBoom }).Ping(ctx), errBoom)
	})

	t.Run("AsPinger", func(t *testing.T) {
		t.Parallel()

		for name, v := range map[string]any{
			"func":       func() {},
			"funcErr":    func() error { return nil },
			"funcCtx":    func(context.Context) {},
			"funcCtxErr": func(context.Context) error { return nil },
			"pinger":     types.PingFunc(func() {}),
		} {
			p, ok := types.AsPinger(v)
			assert.Truef(t, ok, "case %s", name)
			require.NotNilf(t, p, "case %s", name)
			require.NoErrorf(t, p.Ping(ctx), "case %s", name)
		}

		p, ok := types.AsPinger("nope")
		assert.False(t, ok)
		assert.Nil(t, p)
	})
}

func TestStarter(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("adapters", func(t *testing.T) {
		t.Parallel()

		require.NoError(t, types.StartFunc(func() {}).Start(ctx))
		require.ErrorIs(t, types.StartFuncErr(func() error { return errBoom }).Start(ctx), errBoom)
		require.NoError(t, types.StartFuncCtx(func(context.Context) {}).Start(ctx))
		require.ErrorIs(t, types.StartFuncCtxErr(func(context.Context) error { return errBoom }).Start(ctx), errBoom)
	})

	t.Run("AsStarter", func(t *testing.T) {
		t.Parallel()

		for name, v := range map[string]any{
			"func":       func() {},
			"funcErr":    func() error { return nil },
			"funcCtx":    func(context.Context) {},
			"funcCtxErr": func(context.Context) error { return nil },
			"starter":    types.StartFunc(func() {}),
		} {
			s, ok := types.AsStarter(v)
			assert.Truef(t, ok, "case %s", name)
			require.NotNilf(t, s, "case %s", name)
			require.NoErrorf(t, s.Start(ctx), "case %s", name)
		}

		s, ok := types.AsStarter("nope")
		assert.False(t, ok)
		assert.Nil(t, s)
	})
}

func TestStopper(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("adapters", func(t *testing.T) {
		t.Parallel()

		require.NoError(t, types.StopFunc(func() {}).Stop(ctx))
		require.ErrorIs(t, types.StopFuncErr(func() error { return errBoom }).Stop(ctx), errBoom)
		require.NoError(t, types.StopFuncCtx(func(context.Context) {}).Stop(ctx))
		require.ErrorIs(t, types.StopFuncCtxErr(func(context.Context) error { return errBoom }).Stop(ctx), errBoom)
	})

	t.Run("AsStopper", func(t *testing.T) {
		t.Parallel()

		for name, v := range map[string]any{
			"func":       func() {},
			"funcErr":    func() error { return nil },
			"funcCtx":    func(context.Context) {},
			"funcCtxErr": func(context.Context) error { return nil },
			"stopper":    types.StopFunc(func() {}),
		} {
			s, ok := types.AsStopper(v)
			assert.Truef(t, ok, "case %s", name)
			require.NotNilf(t, s, "case %s", name)
			require.NoErrorf(t, s.Stop(ctx), "case %s", name)
		}

		s, ok := types.AsStopper("nope")
		assert.False(t, ok)
		assert.Nil(t, s)
	})
}

func TestShutdowner(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("adapters", func(t *testing.T) {
		t.Parallel()

		require.NoError(t, types.ShutdownFunc(func() {}).Shutdown(ctx))
		require.ErrorIs(t, types.ShutdownFuncErr(func() error { return errBoom }).Shutdown(ctx), errBoom)
		require.NoError(t, types.ShutdownFuncCtx(func(context.Context) {}).Shutdown(ctx))
		require.ErrorIs(t, types.ShutdownFuncCtxErr(func(context.Context) error { return errBoom }).Shutdown(ctx), errBoom)
	})

	t.Run("AsShutdowner", func(t *testing.T) {
		t.Parallel()

		for name, v := range map[string]any{
			"func":       func() {},
			"funcErr":    func() error { return nil },
			"funcCtx":    func(context.Context) {},
			"funcCtxErr": func(context.Context) error { return nil },
			"shutdowner": types.ShutdownFunc(func() {}),
		} {
			s, ok := types.AsShutdowner(v)
			assert.Truef(t, ok, "case %s", name)
			require.NotNilf(t, s, "case %s", name)
			require.NoErrorf(t, s.Shutdown(ctx), "case %s", name)
		}

		s, ok := types.AsShutdowner("nope")
		assert.False(t, ok)
		assert.Nil(t, s)
	})
}

func TestUnsubscriber(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("adapters", func(t *testing.T) {
		t.Parallel()

		require.NoError(t, types.UnsubscribeFunc(func() {}).Unsubscribe(ctx))
		require.ErrorIs(t, types.UnsubscribeFuncErr(func() error { return errBoom }).Unsubscribe(ctx), errBoom)
		require.NoError(t, types.UnsubscribeFuncCtx(func(context.Context) {}).Unsubscribe(ctx))
		require.ErrorIs(t, types.UnsubscribeFuncCtxErr(func(context.Context) error { return errBoom }).Unsubscribe(ctx), errBoom)
	})

	t.Run("AsUnsubscriber", func(t *testing.T) {
		t.Parallel()

		for name, v := range map[string]any{
			"func":         func() {},
			"funcErr":      func() error { return nil },
			"funcCtx":      func(context.Context) {},
			"funcCtxErr":   func(context.Context) error { return nil },
			"unsubscriber": types.UnsubscribeFunc(func() {}),
		} {
			u, ok := types.AsUnsubscriber(v)
			assert.Truef(t, ok, "case %s", name)
			require.NotNilf(t, u, "case %s", name)
			require.NoErrorf(t, u.Unsubscribe(ctx), "case %s", name)
		}

		u, ok := types.AsUnsubscriber("nope")
		assert.False(t, ok)
		assert.Nil(t, u)
	})
}
