package command

import (
	"aurora/file"
	"aurora/file/erf"
	"aurora/tools"
	cliTools "cli/tools"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func createDirectory(destinationPath string) error {
	var stdPermissions = os.FileMode(0744)
	stat, err := os.Stat(destinationPath)

	if os.IsNotExist(err) {
		cliTools.EasyPanic(os.Mkdir(destinationPath, stdPermissions))
	} else if !stat.IsDir() {
		panic(fmt.Sprintf(
			"The given path '%s' is a file. Can't proceed extracting the files.",
			destinationPath,
		))
	} else {
		fmt.Println("The directory already exists.")
		proceed := tools.ReadCommandLineResponse("Do you want to procceed deleting it? [N/y]:")

		if strings.ToLower(proceed)[0] != 'y' {
			return errors.New("Aborting due to user choice")
		}

		cliTools.EasyPanic(tools.RemoveDirectory(destinationPath))
		cliTools.EasyPanic(os.Mkdir(destinationPath, stdPermissions))

	}

	return nil
}

// ExtractErf extract the contents of an erf file
func ExtractErf(module *erf.ERF, destinationPath string) {
	err := createDirectory(destinationPath)

	if err != nil {
		fmt.Println(err.Error())

		return
	}

	var resourceDataOffset = module.Header.OffsetToResourceList + module.Header.EntryCount*8

	for index, key := range module.KeyList {
		var fileName = destinationPath + string(os.PathSeparator) + key.ResRef + "." + file.ResourceTypeLookup[key.ResType]
		var listElement = module.ResourceList[index]
		var offset = listElement.OffsetToResource - resourceDataOffset

		fmt.Println(fmt.Sprintf("Offset: %010d; File size: %010d bytes; File name: %s", offset, listElement.ResourceSize, fileName))
		var fileContent = module.ResourceData[offset : offset+listElement.ResourceSize-1]
		cliTools.EasyPanic(ioutil.WriteFile(fileName, fileContent, os.FileMode(0644)))
	}
}
