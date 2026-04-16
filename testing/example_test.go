package testing_test

import (
	"testing"

	testingx "github.com/foomo/go/testing"
)

func helper(tb testing.TB) {
	tb.Helper()
	tb.Log("log called")
	tb.Logf("logf called with tb: %T", tb)
	tb.Error("error called")
	tb.Errorf("error called with tb: %T", tb)
}

func ExampleNewExampleTB() {
	tb := testingx.NewExampleTB()
	helper(tb)

	// Output:
	// log called
	// logf called with tb: *testing.ExampleTB
	// error: error called
	// error: error: error called with tb: *testing.ExampleTB
}
