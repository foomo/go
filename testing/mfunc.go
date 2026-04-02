package testing

// MFunc is a wrapper type for testing.M that allows for custom behavior
//
//	func TestMain(m *testing.M) {
//	  goleak.VerifyTestMain(testingx.MFunc(func() int {
//			return m.Run()
//		}))
//	}
type MFunc func() int

// Run executes the testing.M function and returns its result
func (r MFunc) Run() int {
	return r()
}
