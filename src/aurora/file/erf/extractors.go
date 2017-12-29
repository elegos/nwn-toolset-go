package erf

import (
	auroraFile "aurora/file"
	"aurora/tools/fileReader"
	"os"
	"strings"
)

func extractHeader(file *os.File) (Header, error) {
	var result = Header{}

	bytes, err := fileReader.ReadAndCheck(file, 4)
	if err != nil {
		return result, err
	}
	result.FileType = strings.Trim(string(bytes), "\x00")

	bytes, err = fileReader.ReadAndCheck(file, 4)
	if err != nil {
		return result, err
	}
	result.Version = strings.Trim(string(bytes), "\x00")

	bytes, err = fileReader.ReadAndCheck(file, 4)
	if err != nil {
		return result, err
	}
	result.LanguageCount = fileReader.BytesToUint32LE(bytes)
	bytes, err = fileReader.ReadAndCheck(file, 4)
	if err != nil {
		return result, err
	}
	result.LocalizedStringSize = fileReader.BytesToUint32LE(bytes)
	bytes, err = fileReader.ReadAndCheck(file, 4)
	if err != nil {
		return result, err
	}
	result.EntryCount = fileReader.BytesToUint32LE(bytes)
	bytes, err = fileReader.ReadAndCheck(file, 4)
	if err != nil {
		return result, err
	}
	result.OffsetToLocalizedString = fileReader.BytesToUint32LE(bytes)
	bytes, err = fileReader.ReadAndCheck(file, 4)
	if err != nil {
		return result, err
	}
	result.OffsetToKeyList = fileReader.BytesToUint32LE(bytes)
	bytes, err = fileReader.ReadAndCheck(file, 4)
	if err != nil {
		return result, err
	}
	result.OffsetToResourceList = fileReader.BytesToUint32LE(bytes)
	bytes, err = fileReader.ReadAndCheck(file, 4)
	if err != nil {
		return result, err
	}
	result.BuildYear = fileReader.BytesToUint32LE(bytes)
	bytes, err = fileReader.ReadAndCheck(file, 4)
	if err != nil {
		return result, err
	}
	result.BuildDay = fileReader.BytesToUint32LE(bytes)
	bytes, err = fileReader.ReadAndCheck(file, 4)
	if err != nil {
		return result, err
	}
	result.DescriptionStrRef = fileReader.BytesToUint32LE(bytes)

	bytes, err = fileReader.ReadAndCheck(file, 116)
	if err != nil {
		return result, err
	}
	copy(result.Reserved[:], bytes)

	return result, nil
}

func extractLocalizedStringList(file *os.File, localizedStringSize uint32) ([]LocalizedStringElement, error) {
	var result = []LocalizedStringElement{}

	var i = uint32(0)
	for ; i < localizedStringSize; i++ {
		langID, err := fileReader.ReadAndCheck(file, 4)
		if err != nil {
			return result, err
		}
		size, err := fileReader.ReadAndCheck(file, 4)
		if err != nil {
			return result, err
		}

		var element = LocalizedStringElement{
			LanguageID: auroraFile.Language(fileReader.BytesToUint32LE(langID)),
			StringSize: fileReader.BytesToUint32LE(size),
		}

		str, err := fileReader.ReadAndCheck(file, uint32(element.StringSize))
		if err != nil {
			return result, err
		}
		element.String = strings.Trim(string(str), "\x00")

		result = append(result, element)
	}

	return result, nil
}

func extractKeyList(file *os.File, offsetToKeyList int64, entryCount uint32) ([]KeyElement, error) {
	var result = []KeyElement{}

	_, err := file.Seek(offsetToKeyList, 0)
	if err != nil {
		return result, err
	}

	var i = uint32(0)
	for ; i < entryCount; i++ {
		resRef, err := fileReader.ReadAndCheck(file, 16)
		if err != nil {
			return result, err
		}

		resID, err := fileReader.ReadAndCheck(file, 4)
		if err != nil {
			return result, err
		}

		resType, err := fileReader.ReadAndCheck(file, 2)
		if err != nil {
			return result, err
		}

		var element = KeyElement{
			ResRef:  strings.Trim(string(resRef), "\x00"),
			ResID:   fileReader.BytesToUint32LE(resID),
			ResType: auroraFile.ResourceType(fileReader.BytesToUint16LE(resType)),
		}

		unused, err := fileReader.ReadAndCheck(file, 2)
		if err != nil {
			return result, err
		}
		copy(element.Unused[:], unused)

		result = append(result, element)
	}

	return result, nil
}

func extractResourceList(file *os.File, offsetToResourceList int64, entryCount uint32) ([]ResourceElement, error) {
	var result = []ResourceElement{}

	_, err := file.Seek(offsetToResourceList, 0)
	if err != nil {
		return result, err
	}

	var i = uint32(0)
	for ; i < entryCount; i++ {
		offset, err := fileReader.ReadAndCheck(file, 4)
		if err != nil {
			return result, err
		}

		size, err := fileReader.ReadAndCheck(file, 4)
		if err != nil {
			return result, err
		}

		result = append(result, ResourceElement{
			OffsetToResource: fileReader.BytesToUint32LE(offset),
			ResourceSize:     fileReader.BytesToUint32LE(size),
		})
	}

	return result, nil
}

func extractResourceData(file *os.File, offsetToResourceList uint32, entryCount uint32) ([]byte, error) {
	var toSkip = offsetToResourceList + entryCount*8

	// Seek to the end of the resources list
	_, err := file.Seek(int64(toSkip), 0)
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
