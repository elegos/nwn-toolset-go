package erf

import (
	auroraFile "aurora/file"
	"aurora/tools"
	"aurora/tools/fileReader"
	"os"
	"strings"
)

func extractHeader(file *os.File, errorBag *tools.ErrorBag) Header {
	if errorBag.Error != nil {
		return Header{}
	}

	var readerBag = fileReader.ByteReaderBag{File: file}
	var result = Header{
		FileType:                strings.Trim(fileReader.ReadStringFromBytes(&readerBag, 4), "\x00"),
		Version:                 strings.Trim(fileReader.ReadStringFromBytes(&readerBag, 4), "\x00"),
		LanguageCount:           fileReader.ReadUint32FromBytes(&readerBag),
		LocalizedStringSize:     fileReader.ReadUint32FromBytes(&readerBag),
		EntryCount:              fileReader.ReadUint32FromBytes(&readerBag),
		OffsetToLocalizedString: fileReader.ReadUint32FromBytes(&readerBag),
		OffsetToKeyList:         fileReader.ReadUint32FromBytes(&readerBag),
		OffsetToResourceList:    fileReader.ReadUint32FromBytes(&readerBag),
		BuildYear:               fileReader.ReadUint32FromBytes(&readerBag),
		BuildDay:                fileReader.ReadUint32FromBytes(&readerBag),
		DescriptionStrRef:       fileReader.ReadUint32FromBytes(&readerBag),
	}

	var reservedBytes = fileReader.ReadBytes(&readerBag, 116)

	copy(result.Reserved[:], reservedBytes)

	errorBag.Error = readerBag.Err

	return result
}

func extractLocalizedStringList(file *os.File, localizedStringSize uint32, errorBag *tools.ErrorBag) []LocalizedStringElement {
	if errorBag.Error != nil {
		return []LocalizedStringElement{}
	}

	var result = []LocalizedStringElement{}
	// TODO bug? From where do I read? Missing file.Seek

	var i = uint32(0)
	var readerBag = fileReader.ByteReaderBag{File: file}
	for ; i < localizedStringSize; i++ {
		var element = LocalizedStringElement{
			LanguageID: auroraFile.Language(fileReader.ReadUint32FromBytes(&readerBag)),
			StringSize: fileReader.ReadUint32FromBytes(&readerBag),
		}
		element.String = strings.Trim(fileReader.ReadStringFromBytes(&readerBag, element.StringSize), "\x00")

		result = append(result, element)
	}

	errorBag.Error = readerBag.Err

	return result
}

func extractKeyList(file *os.File, offsetToKeyList int64, entryCount uint32, errorBag *tools.ErrorBag) []KeyElement {
	var result = []KeyElement{}

	if errorBag.Error != nil {
		return result
	}

	_, err := file.Seek(offsetToKeyList, os.SEEK_SET)
	if err != nil {
		errorBag.Error = err

		return result
	}

	var i = uint32(0)
	var readerBag = fileReader.ByteReaderBag{File: file}
	for ; i < entryCount; i++ {
		var element = KeyElement{
			ResRef:  strings.Trim(fileReader.ReadStringFromBytes(&readerBag, 16), "\x00"),
			ResID:   fileReader.ReadUint32FromBytes(&readerBag),
			ResType: auroraFile.ResourceType(fileReader.ReadUint16FromBytes(&readerBag)),
		}

		unused := fileReader.ReadBytes(&readerBag, 2)
		copy(element.Unused[:], unused)

		result = append(result, element)
	}

	errorBag.Error = readerBag.Err

	return result
}

func extractResourceList(file *os.File, offsetToResourceList int64, entryCount uint32, errorBag *tools.ErrorBag) []ResourceElement {
	var result = []ResourceElement{}

	if errorBag.Error != nil {
		return []ResourceElement{}
	}

	_, err := file.Seek(offsetToResourceList, os.SEEK_SET)
	if err != nil {
		errorBag.Error = err

		return result
	}

	var i = uint32(0)
	var readerBag = fileReader.ByteReaderBag{File: file}
	for ; i < entryCount; i++ {
		result = append(result, ResourceElement{
			OffsetToResource: fileReader.ReadUint32FromBytes(&readerBag),
			ResourceSize:     fileReader.ReadUint32FromBytes(&readerBag),
		})
	}

	errorBag.Error = readerBag.Err

	return result
}

func extractResourceData(file *os.File, offsetToResourceList uint32, entryCount uint32, errorBag *tools.ErrorBag) []byte {
	if errorBag.Error != nil {
		return []byte{}
	}

	var toSkip = offsetToResourceList + entryCount*8

	// Seek to the end of the resources list
	_, err := file.Seek(int64(toSkip), os.SEEK_SET)
	if err != nil {
		errorBag.Error = err

		return []byte{}
	}

	stat, err := file.Stat()
	if err != nil {
		errorBag.Error = err

		return []byte{}
	}

	toRead := stat.Size() - int64(toSkip)

	result, err := fileReader.ReadAndCheck(file, uint32(toRead))
	errorBag.Error = err

	return result
}
