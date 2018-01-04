package command

import (
	"aurora/file/are"
	"cli/tools"
	"fmt"
	"io/ioutil"
)

// ReadAreFromFile read an ARE file's information
func ReadAreFromFile(file string) {
	bytes, err := ioutil.ReadFile(file)
	tools.EasyPanic(err)

	areFile, err := are.FromBytes(bytes)
	tools.EasyPanic(err)

	fmt.Println(areFile.Header)
}
