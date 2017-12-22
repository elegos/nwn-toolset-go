package erf

import (
	"aurora/tools"
	"aurora/tools/fileReader"
	"os"
)

func extractHeader(file *os.File) Header {
	var result = Header{}

	result.FileType = string(fileReader.ReadAndCheck(file, 4))
	result.Version = string(fileReader.ReadAndCheck(file, 4))
	result.LanguageCount = fileReader.BytesToUint32LE(fileReader.ReadAndCheck(file, 4))
	result.LocalizedStringSize = fileReader.BytesToUint32LE(fileReader.ReadAndCheck(file, 4))
	result.EntryCount = fileReader.BytesToUint32LE(fileReader.ReadAndCheck(file, 4))
	result.OffsetToLocalizedString = fileReader.BytesToUint32LE(fileReader.ReadAndCheck(file, 4))
	result.OffsetToKeyList = fileReader.BytesToUint32LE(fileReader.ReadAndCheck(file, 4))
	result.OffsetToResourceList = fileReader.BytesToUint32LE(fileReader.ReadAndCheck(file, 4))
	result.BuildYear = fileReader.BytesToUint32LE(fileReader.ReadAndCheck(file, 4))
	result.BuildDay = fileReader.BytesToUint32LE(fileReader.ReadAndCheck(file, 4))
	result.DescriptionStrRef = fileReader.BytesToUint32LE(fileReader.ReadAndCheck(file, 4))

	copy(result.Reserved[:], fileReader.ReadAndCheck(file, 116))

	return result
}

func extractLocalizedStringList(file *os.File, localizedStringSize uint32) []LocalizedStringElement {
	var result = []LocalizedStringElement{}

	var i = uint32(0)
	for ; i < localizedStringSize; i++ {
		var element = LocalizedStringElement{
			LanguageID: Language(fileReader.BytesToUint32LE(fileReader.ReadAndCheck(file, 4))),
			StringSize: fileReader.BytesToUint32LE(fileReader.ReadAndCheck(file, 4)),
		}

		element.String = string(fileReader.ReadAndCheck(file, element.StringSize))

		result = append(result, element)
	}

	return result
}

func extractKeyList(file *os.File, offsetToKeyList int64, entryCount uint32) []KeyElement {
	var result = []KeyElement{}

	_, err := file.Seek(offsetToKeyList, 0)
	tools.EasyPanic(err)

	var i = uint32(0)
	for ; i < entryCount; i++ {
		var element = KeyElement{
			ResRef:  string(fileReader.ReadAndCheck(file, 16)),
			ResID:   fileReader.BytesToUint32LE(fileReader.ReadAndCheck(file, 4)),
			ResType: fileReader.BytesToUint16LE(fileReader.ReadAndCheck(file, 2)),
		}

		var unused = fileReader.ReadAndCheck(file, 2)
		copy(element.Unused[:], unused)

		result = append(result, element)
	}

	return result
}

func extractResourceList(file *os.File, offsetToResourceList int64, entryCount uint32) []ResourceElement {
	var result = []ResourceElement{}

	_, err := file.Seek(offsetToResourceList, 0)
	tools.EasyPanic(err)

	var i = uint32(0)
	for ; i < entryCount; i++ {
		result = append(result, ResourceElement{
			OffsetToResource: fileReader.BytesToUint32LE(fileReader.ReadAndCheck(file, 4)),
			ResourceSize:     fileReader.BytesToUint32LE(fileReader.ReadAndCheck(file, 4)),
		})
	}

	return result
}

func extractResourceData(file *os.File, offsetToResourceList uint32, entryCount uint32) []byte {
	var toSkip = offsetToResourceList + entryCount*8

	// Seek to the end of the resources list
	_, err := file.Seek(int64(toSkip), 0)
	tools.EasyPanic(err)

	stat, err := file.Stat()
	tools.EasyPanic(err)

	toRead := stat.Size() - int64(toSkip)

	return fileReader.ReadAndCheck(file, uint32(toRead))
}
