package runtime_test

import (
	"testing"

	"github.com/foomo/go/runtime"
)

func TestParseName(t *testing.T) {
	cases := []struct{ in, pkg, inst, fn string }{
		{"main.main", "main", "", "main"},
		{"github.com/b/billing/payment.Process", "github.com/b/billing/payment", "", "Process"},
		{"github.com/b/billing/payment.(*Payment).Process", "github.com/b/billing/payment", "Payment", "Process"},
		{"github.com/b/billing/payment.Payment.Process", "github.com/b/billing/payment", "Payment", "Process"},
		{"github.com/b/billing/payment.(*Payment).Process.func1", "github.com/b/billing/payment", "Payment", "Process.func1"},
		{"github.com/b/billing/payment.Process.func1", "github.com/b/billing/payment", "", "Process.func1"},
		{"github.com/b/billing/payment.Process.func1.2", "github.com/b/billing/payment", "", "Process.func1.2"},
		{"github.com/b/billing/payment.Map[go.shape.int]", "github.com/b/billing/payment", "", "Map[go.shape.int]"},
		{"github.com/b/billing/payment.(*Cache[go.shape.string]).Get", "github.com/b/billing/payment", "Cache[go.shape.string]", "Get"},
		{"github.com/b/billing/payment.(*Payment).Process-fm", "github.com/b/billing/payment", "Payment", "Process-fm"},
		{"github.com/b/billing/payment.glob..func1", "github.com/b/billing/payment", "", "glob..func1"},
		{"github.com/b/billing/payment.init.0", "github.com/b/billing/payment", "", "init.0"},
		{"github.com/b/svc/v2.Run", "github.com/b/svc/v2", "", "Run"},
		{"", "", "", ""},
	}
	for _, c := range cases {
		pkg, inst, fn := runtime.ParseName(c.in)
		if pkg != c.pkg || inst != c.inst || fn != c.fn {
			t.Errorf("parseName(%q)\n got  pkg=%q inst=%q fn=%q\n want pkg=%q inst=%q fn=%q",
				c.in, pkg, inst, fn, c.pkg, c.inst, c.fn)
		}
	}
}

//nolint:recvcheck // intentionally
type widget struct{}

func (w *widget) ptrMethod() runtime.Frame { return runtime.CallFrame(0) }
func (w widget) valMethod() runtime.Frame  { return runtime.CallFrame(0) }
func pkgFunc() runtime.Frame               { return runtime.CallFrame(0) }
func nested() runtime.Frame                { return func() runtime.Frame { return runtime.CallFrame(0) }() }
func indirect() runtime.Frame              { return level2() }
func level2() runtime.Frame                { return runtime.CallFrame(1) }
func generic[T any](v T) runtime.Frame     { return runtime.CallFrame(0) }

func TestCallFrame(t *testing.T) {
	cases := []struct {
		got      runtime.Frame
		inst, fn string
	}{
		{(&widget{}).ptrMethod(), "widget", "ptrMethod"},
		{widget{}.valMethod(), "widget", "valMethod"},
		{pkgFunc(), "", "pkgFunc"},
		{nested(), "", "nested.func1"},
		{indirect(), "", "indirect"},
		{generic(42), "", "generic[...]"},
	}
	for _, c := range cases {
		if c.got.Inst != c.inst || c.got.Func != c.fn {
			t.Errorf("got inst=%q fn=%q, want inst=%q fn=%q (name=%q)",
				c.got.Inst, c.got.Func, c.inst, c.fn, c.got.Name())
		}

		if c.got.Line == 0 || c.got.File == "" {
			t.Errorf("%s: missing file/line", c.got.Name())
		}

		if c.got.Pkg != "github.com/foomo/go/runtime_test" {
			t.Errorf("%s: unexpected pkg %q", c.got.Name(), c.got.Pkg)
		}
	}
}

func BenchmarkCallFrame(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = runtime.CallFrame(0)
	}
}
