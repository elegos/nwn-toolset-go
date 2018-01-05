package gff

import (
	auroraFile "aurora/file"
	"aurora/tools/fileReader"
	"math"
)

func fieldExtractorCExoString(fieldDataBlock []byte, offset uint32) string {
	var size = fileReader.BytesToUint32LE(fieldDataBlock[offset : offset+4])
	var start = offset + 4
	var end = start + size

	return string(fieldDataBlock[start:end])
}

func fieldExtractorCResRef(fieldDataBlock []byte, offset uint32) string {
	var size = uint32(fieldDataBlock[offset])
	var start = offset + 1
	var end = start + size

	return string(fieldDataBlock[start:end])
}

func fieldExtractorCExoLocString(fieldDataBlock []byte, offset uint32) CExoLocString {
	var cExoLocString = CExoLocString{}

	var totalSize = fileReader.BytesToUint32LE(fieldDataBlock[offset : offset+4])
	if totalSize == 0 {
		return cExoLocString
	}
	var index = offset + 4

	cExoLocString.StringRef = fileReader.BytesToUint32LE(fieldDataBlock[index : index+4])
	index += 4

	var stringCount = fileReader.BytesToUint32LE(fieldDataBlock[index : index+4])
	index += 4

	var i = uint32(0)
	for ; i < stringCount; i++ {
		var subString = CExoLocStringSubString{}

		var rawStringID = fileReader.BytesToUint32LE(fieldDataBlock[index : index+4])
		index += 4
		if (rawStringID % 2) == 0 {
			subString.Masculine = false
			subString.LanguageID = auroraFile.Language(rawStringID / 2)
		} else {
			subString.Masculine = true
			subString.LanguageID = auroraFile.Language((rawStringID - 1) / 2)
		}

		var stringSize = fileReader.BytesToUint32LE(fieldDataBlock[index : index+4])
		index += 4
		var start = index
		var end = start + stringSize
		index = index + stringSize

		subString.String = string(fieldDataBlock[start:end])

		cExoLocString.SubStrings = append(cExoLocString.SubStrings, subString)
	}

	return cExoLocString
}

func extractFieldData(fieldType FieldType, dataOrDataOffset uint32, file GFF) FieldDataStruct {
	var result = FieldDataStruct{}

	switch fieldType {
	// Simple data
	case FieldTypeByte:
		result.ByteValue = int8(dataOrDataOffset)
		break
	case FieldTypeChar:
		result.CharValue = string(dataOrDataOffset)
		break
	case FieldTypeWord:
		result.WordValue = uint16(dataOrDataOffset)
		break
	case FieldTypeShort:
		result.ShortValue = int16(dataOrDataOffset)
		break
	case FieldTypeDword:
		result.DwordValue = dataOrDataOffset
		break
	case FieldTypeInt:
		result.IntValue = int32(dataOrDataOffset)
		break
	case FieldTypeFloat:
		result.FloatValue = math.Float32frombits(dataOrDataOffset)
		break
		// Complex data (TODO)
	case FieldTypeDword64:
	case FieldTypeInt64:
	case FieldTypeDouble:
		break
	case FieldTypeCExoString:
		result.CExoStringValue = fieldExtractorCExoString(file.FieldDataBlock, dataOrDataOffset)
		break
	case FieldTypeCResRef:
		result.CResRefValue = fieldExtractorCResRef(file.FieldDataBlock, dataOrDataOffset)
		break
	case FieldTypeCExoLocString:
		result.CExoLocStringValue = fieldExtractorCExoLocString(file.FieldDataBlock, dataOrDataOffset)
		break
	case FieldTypeVoid:
	// Complex special data (TODO)
	case FieldTypeStruct:
		break
	case FieldTypeList:
		result.ListOffsetValue = dataOrDataOffset
		break
	}

	return result
}
