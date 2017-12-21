package tools

// EasyPanic panic if the error is not nil
func EasyPanic(err error) {
	if err != nil {
		panic(err)
	}
}
