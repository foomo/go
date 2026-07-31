package runtime

// Frame identifies a single call site.
//
// Pkg is the full import path, Inst the receiver type name for methods (empty
// for package-level functions), and Func the function or method name including
// any compiler-generated suffix such as ".func1" or "-fm".
type Frame struct {
	Pkg  string // "github.com/foomo/go/runtime"
	Inst string // "Memo"; empty for package-level functions
	Func string // "Get", "Meemo.func1", "Map[go.shape.int]"
	File string // absolute path as recorded at build time
	Line int
}

// Name reassembles the qualified name, suitable for the code.function.name
// attribute. The receiver's pointer marker is not preserved: a method on
// *Payment and one on Payment both render as "pkg.Payment.Method".
func (f Frame) Name() string {
	switch {
	case f.Pkg == "":
		return f.Func
	case f.Inst == "":
		return f.Pkg + "." + f.Func
	default:
		return f.Pkg + "." + f.Inst + "." + f.Func
	}
}

// Short renders the receiver-qualified name without the package path, e.g.
// "Payment.Process" or "Process" for a package-level function.
func (f Frame) Short() string {
	if f.Inst == "" {
		return f.Func
	}

	return f.Inst + "." + f.Func
}

// Zero reports whether the frame could not be resolved. A zero Frame is
// returned rather than a panic when the stack is unavailable.
func (f Frame) Zero() bool { return f.Func == "" && f.Pkg == "" }
