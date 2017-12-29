package erf

import (
	auroraFile "aurora/file"
	"aurora/tools/fileReader"
	"os"
	"strings"
)

func extractHeader(file *os.File) (Header, error) {
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

	return result, readerBag.Err
}

func extractLocalizedStringList(file *os.File, localizedStringSize uint32) ([]LocalizedStringElement, error) {
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

	return result, readerBag.Err
}

func extractKeyList(file *os.File, offsetToKeyList int64, entryCount uint32) ([]KeyElement, error) {
	var result = []KeyElement{}

	_, err := file.Seek(offsetToKeyList, os.SEEK_SET)
	if err != nil {
		return result, err
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

	return result, readerBag.Err
}

func extractResourceList(file *os.File, offsetToResourceList int64, entryCount uint32) ([]ResourceElement, error) {
	var result = []ResourceElement{}

	_, err := file.Seek(offsetToResourceList, os.SEEK_SET)
	if err != nil {
		return result, err
	}

	var i = uint32(0)
	var readerBag = fileReader.ByteReaderBag{File: file}
	for ; i < entryCount; i++ {
		result = append(result, ResourceElement{
			OffsetToResource: fileReader.ReadUint32FromBytes(&readerBag),
			ResourceSize:     fileReader.ReadUint32FromBytes(&readerBag),
		})
	}

	return result, readerBag.Err
}

func extractResourceData(file *os.File, offsetToResourceList uint32, entryCount uint32) ([]byte, error) {
	var toSkip = offsetToResourceList + entryCount*8

	// Seek to the end of the resources list
	_, err := file.Seek(int64(toSkip), os.SEEK_SET)
	if err != nil {
		return []byte{}, err
	}

	stat, err := file.Stat()
	if err != nil {
		return []byte{}, err
	}

	toRead := stat.Size() - int64(toSkip)

	return fileReader.ReadAndCheck(file, uint32(toRead))
}
