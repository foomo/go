package testing

// M is an interface representing the testing.M type, which is used for customizing the behavior of TestMain.
type M interface {
	Run() int
}
